const urlParams = new URLSearchParams(window.location.search);
const problemId = urlParams.get('id');

function loadProblemById() {
    fetch(`/api/v1/problem/${problemId}`)
    .then(response => response.json())
    .then(problem => {
        document.getElementById('problem-title').textContent = problem.title;
        document.getElementById('problem-body').textContent = problem.body;
    })
    .catch(error => {
        console.error('Error:', error);
        alert('問題の読み込みに失敗しました。');
    });
}

window.onload = function() {
    loadProblemById();
    document.getElementById('submit-form').addEventListener('submit', function(e) {
        e.preventDefault();
        const sourceCode = document.getElementById('source-code').value;
        
        fetch(`/api/v1/problem/${problemId}/submission`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ sourceCode }),
        })
        .then(response => response.json())
        .then(data => {
            alert('提出が完了しました。');
            window.location.href = `submission.html?problemId=${problemId}`;
        })
        .catch((error) => {
            console.error('Error:', error);
            alert('提出に失敗しました。');
        });
    });

    document.getElementById('submission-link').href = `submission.html?problemId=${problemId}`;
}