console.log("js file download!")

function Sign_In() {
    let xhr = new XMLHttpRequest()
    log = document.getElementById("login")
    pass = document.getElementById("password")

    let json = JSON.stringify({
        login: log.value,
        password: pass.value
    });

    xhr.open("POST", '/sig-in')
    xhr.setRequestHeader('Content-type', 'application/json; charset=utf-8');
    xhr.send(json);


    xhr.onload = function () {
            if (xhr.readyState === 4 && xhr.status === 200) {
                var response = JSON.parse(xhr.responseText);

                if (response.success) {
                } else {
                    document.getElementById("message").innerText =  "Вы зашли";
                    window.location.href = "/feed";
                }
            } else if (xhr.readyState === 4) {
                document.getElementById("message").innerText = "Логин или пароль не существует";
            }
    }
    console.log(json)
}

function formatDate(dateString) {
    const options = { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' };
    return new Date(dateString).toLocaleDateString('ru-RU', options);
}

function renderComments(comments) {
    const commentsSection = document.getElementById('comments-section');
    commentsSection.innerHTML = '';

    comments.forEach(comment => {
        const commentDiv = document.createElement('div');
        commentDiv.classList.add('comment');
        commentDiv.style.color = 'red'; // Исправлено на 'red'
        commentDiv.innerHTML = `
            <p><strong>Пользователь ${comment.user_id}:</strong></p>
            <p>${comment.content}</p>
            <p><em>Создано: ${formatDate(comment.date_created)}</em></p>
        `;

        commentsSection.appendChild(commentDiv);
    });
}

function loadPosts() {
    fetch('/post', {
        method: 'GET',
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка сети: ' + response.status);
            }
            return response.json();
        })
        .then(data => {
            console.log('Данные получены:', data);
            renderComments(data); // Используем функцию для рендеринга комментариев
        })
        .catch(error => {
            console.error('Ошибка:', error);
        });
}


function signup() {

        const login = document.getElementById('login').value;
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        const full_name = document.getElementById('full_name_user').value;
        const birthday = document.getElementById('birthday_user').value;
        const photo = document.getElementById('photo').value;
        const confirmPassword = document.getElementById('confirmPassword').value;
        const errorMessage = document.getElementById('errorMessage');

        errorMessage.textContent = '';

        if (password !== confirmPassword) {
            errorMessage.textContent = 'Passwords do not match!';
            return;
        }

        const user = {
            login: login,
            email: email,
            password: password,
            full_name: full_name,
            birthday: birthday,
            photo: photo

        };

        fetch('/sig-up', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(user)
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Registration failed');
                }
                return response.json();
            })
            .then(data => {
                alert('Registration successful!');
            })
            .catch(error => {
                errorMessage.textContent = error.message;
            });

}

window.onload = function() {

    const token = document.cookie
        .split('; ')
        .find(row => row.startsWith('token='));

    if (token) {
        // Если токен найден, перенаправляем пользователя на страницу с профилем
        window.location.href = '/feed';
    } else {
        // Если токена нет, остаемся на текущей странице (например, странице входа)
        console.log("No token found. Please log in.");

    }
};

loadPosts();
