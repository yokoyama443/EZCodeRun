window.onload = function() {
    document.getElementById('register-form').addEventListener('submit', function(e) {
        e.preventDefault();
        const name = document.getElementById('name').value;
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        
        fetch('/api/auth/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name, email, password }),
        })
        .then(response => {
            if (response.status === 201) {
                alert('アカウントが作成されました。ログインしてください。');
                window.location.href = 'login.html';
            } else {
                alert('アカウント作成に失敗しました。');
            }
        })
        .catch((error) => {
            console.error('Error:', error);
            alert('エラーが発生しました。');
        });
    });
}