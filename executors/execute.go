package executors

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"ez-code-run/models"
)

func ExecuteCode(submission *models.Submission, problem *models.Problem) {
	go func() {
		// 一時ディレクトリ作成
		tempDir, err := ioutil.TempDir("", "submission")
		if err != nil {
			updateSubmissionStatus(submission, "Failed", "Failed to create temporary directory")
			return
		}
		defer os.RemoveAll(tempDir)

		// ソースコード書き込み
		sourceFile := filepath.Join(tempDir, "main.cpp")
		err = ioutil.WriteFile(sourceFile, []byte(submission.SourceCode), 0644)
		if err != nil {
			updateSubmissionStatus(submission, "Failed", "Failed to write source code")
			return
		}

		// コンパイル
		cmd := exec.Command("g++", sourceFile, "-o", filepath.Join(tempDir, "a.out"))
		var compileErr bytes.Buffer
		cmd.Stderr = &compileErr
		err = cmd.Run()
		if err != nil {
			updateSubmissionStatus(submission, "Failed", fmt.Sprintf("Compilation error: %s", compileErr.String()))
			return
		}

		// 入力ファイル作成
		inputFile := filepath.Join(tempDir, "input.txt")
		err = ioutil.WriteFile(inputFile, []byte(problem.TestCaseInput), 0644)
		if err != nil {
			updateSubmissionStatus(submission, "Failed", "Failed to write input file")
			return
		}

		// 出力ファイル作成
		outputFile := filepath.Join(tempDir, "output.txt")

		// 実行 (ジョブのシェルを作って時間制限とメモリ制限を設定)
		cmd = exec.Command("bash", "-c", fmt.Sprintf(`
			ulimit -v %d
			timeout %ds ./a.out < input.txt > output.txt
		`, problem.MemoryLimit*1024, problem.TimeLimit))
		cmd.Dir = tempDir
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

		start := time.Now()
		err = cmd.Run()
		executionTime := time.Since(start)

		// タイムアウトやメモリ制限超過の場合
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				if exitError.ExitCode() == 124 {
					updateSubmissionStatus(submission, "Failed", "Time Limit Exceeded")
				} else if exitError.ExitCode() == 137 {
					updateSubmissionStatus(submission, "Failed", "Memory Limit Exceeded")
				} else {
					updateSubmissionStatus(submission, "Failed", fmt.Sprintf("Runtime Error: %s", err))
				}
			} else {
				updateSubmissionStatus(submission, "Failed", fmt.Sprintf("Execution error: %s", err))
			}
			return
		}

		// 出力の比較
		output, err := ioutil.ReadFile(outputFile)
		if err != nil {
			updateSubmissionStatus(submission, "Failed", "Failed to read output file")
			return
		}

		if string(output) == problem.TestCaseOutput {
			updateSubmissionStatus(submission, "Success", fmt.Sprintf("Execution time: %v", executionTime))
		} else {
			updateSubmissionStatus(submission, "Failed", "Wrong Answer")
		}
	}()
}

func updateSubmissionStatus(submission *models.Submission, status string, message string) {
	fmt.Println(message)
	submission.ResultStatus = status
	models.DB.Save(submission)
}
