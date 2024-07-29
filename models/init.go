package models

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(os.Getenv("DB_PATH")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&User{}, &Problem{}, &Submission{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	addSampleProblems()
}

// サンプル問題追加用
func addSampleProblems() {
	var count int64
	DB.Model(&Problem{}).Count(&count)
	if count == 0 {
		problems := []Problem{
			{
				Title:             "Hello, World!",
				TimeLimit:         1,
				MemoryLimit:       128,
				FreeFeedBackLimit: 3,
				Body:              "標準出力に「Hello, World!」と出力するプログラムを作成してください。",
				TestCaseInput:     "",
				TestCaseOutput:    "Hello, World!\n",
				Priority:          1,
			},
			{
				Title:             "足し算",
				TimeLimit:         1,
				MemoryLimit:       128,
				FreeFeedBackLimit: 3,
				Body:              "2つの整数 A と B が与えられます。A + B の結果を出力するプログラムを作成してください。\n\n入力:\n2つの整数 A, B (0 ≤ A, B ≤ 1000)\n\n出力:\nA + B の結果",
				TestCaseInput:     "5 3\n",
				TestCaseOutput:    "8\n",
				Priority:          2,
			},
		}

		for _, problem := range problems {
			result := DB.Create(&problem)
			if result.Error != nil {
				log.Printf("Failed to create sample problem: %v", result.Error)
			} else {
				log.Printf("Sample problem created: %s", problem.Title)
			}
		}
	}
}
