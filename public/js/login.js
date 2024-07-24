window.onload = function() {
    document.getElementById('login-form').addEventListener('submit', function(e) {
        e.preventDefault();
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        
        fetch('/api/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
        })
        .then(response => response.json())
        .then(data => {
            if (data.message === "Logged in successfully") {
                window.location.href = 'problems.html';
            } else {
                alert('ログインに失敗しました。');
            }
        })
        .catch((error) => {
            console.error('Error:', error);
            alert('エラーが発生しました。');
        });
    });
}