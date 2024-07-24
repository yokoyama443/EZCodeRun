const urlParams = new URLSearchParams(window.location.search);
const problemId = urlParams.get('problemId');

function loadSubmissions() {
    fetch(`/api/v1/problem/${problemId}/submission`)
    .then(response => response.json())
    .then(submissions => {
        const submissionList = document.getElementById('submission-list');
        submissions.forEach(submission => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${new Date(submission.created_at).toLocaleString()}</td>
                <td class="status-${submission.resultStatus.toLowerCase()}">${submission.resultStatus}</td>
                <td>${submission.executionTime || '-'} ms</td>
                <td>${submission.memoryUsed || '-'} KB</td>
            `;
            submissionList.appendChild(row);
        });
    })
    .catch(error => {
        console.error('Error:', error);
        alert('提出結果の読み込みに失敗しました。');
    });
}

window.onload = loadSubmissions();