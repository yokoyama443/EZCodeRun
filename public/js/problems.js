function loadProblems() {
    fetch('/api/v1/problem')
    .then(response => response.json())
    .then(problems => {
        const problemList = document.getElementById('problem-list');
        problems.forEach(problem => {
            const problemCard = document.createElement('div');
            problemCard.className = 'problem-card';
            problemCard.innerHTML = `
                <h2>${problem.title}</h2>
                <p>${problem.description}</p>
                <a href="problem-detail.html?id=${problem.id}">問題を解く</a>
            `;
            problemList.appendChild(problemCard);
        });
    })
    .catch(error => {
        console.error('Error:', error);
        alert('問題の読み込みに失敗しました。');
    });
}

window.onload = loadProblems;
