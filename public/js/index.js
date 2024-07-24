function checkLoginStatus() {
    fetch('/api/auth/user', {
        method: 'GET',
        credentials: 'include'
    })
    .then(response => response.json())
    .then(data => {
        const buttonContainer = document.getElementById('button-container');
        if (data.id) {
            buttonContainer.innerHTML = '<a href="problems.html" class="button">問題一覧へ</a>';
        } else {
            buttonContainer.innerHTML = `
                <a href="login.html" class="button">ログイン</a>
                <a href="register.html" class="button">アカウント登録</a>
            `;
        }
    })
    .catch(error => {
        console.error('Error:', error);
        // エラー時はログアウト状態として扱う
        const buttonContainer = document.getElementById('button-container');
        buttonContainer.innerHTML = `
            <a href="login.html" class="button">ログイン</a>
            <a href="register.html" class="button">アカウント登録</a>
        `;
    });
}

// ページ読み込み時にログイン状態をチェック
window.onload = checkLoginStatus;