console.log("JS file loaded!");

function Sign_In() {
    let xhr = new XMLHttpRequest();
    let log = document.getElementById("login");
    let pass = document.getElementById("password");

    let json = JSON.stringify({
        login: log.value,
        password: pass.value
    });

    xhr.open("POST", '/sign-in');
    xhr.setRequestHeader('Content-type', 'application/json; charset=utf-8');
    xhr.send(json);

    xhr.onload = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var response = JSON.parse(xhr.responseText);

            if (response.token) {
                localStorage.setItem("token", response.token);

                loadProfile();

                window.location.href = "/feed";
            } else {
                document.getElementById("message").innerText = "Ошибка входа";
            }
        } else if (xhr.readyState === 4) {
            document.getElementById("message").innerText = "Логин или пароль неверны";
        }
    };

    console.log(json);
}

function loadProfile() {
    const token = localStorage.getItem("token");

    if (!token) {
        console.log("Нет токена, загрузка профиля невозможна.");
        return;
    }

    fetch("/users", {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${token}`
        }
    })
        .then(response => {
            if (!response.ok) {
                throw new Error("Ошибка при загрузке профиля");
            }
            return response.json();
        })
        .then(user => {
            console.log("Профиль пользователя:", user);

            document.getElementById("profile").innerHTML = `
            <h2> ${user.full_name}</h2>
            <img src="${user.photo}" alt="Фото профиля" width="100">
        `;
        })
        .catch(error => {
            console.error("Ошибка загрузки профиля:", error);
        });
}

function renderComments(comments) {
    const commentsSection = document.getElementById("comments-section");
    commentsSection.innerHTML = "";

    comments.forEach(comment => {
        const commentDiv = document.createElement("div");
        commentDiv.classList.add("comment");
        commentDiv.style.color = "red";
        commentDiv.innerHTML = `
            <p><strong>Пользователь ${comment.user_id}:</strong></p>
            <p>${comment.content}</p>
            <p><em>Создано: ${formatDate(comment.date_created)}</em></p>
        `;

        commentsSection.appendChild(commentDiv);
    });
}

function loadPosts() {
    fetch("/post", { method: "GET" })
        .then(response => {
            if (!response.ok) {
                throw new Error("Ошибка сети: " + response.status);
            }
            return response.json();
        })
        .then(data => {
            console.log("Данные получены:", data);
            renderComments(data);
        })
        .catch(error => {
            console.error("Ошибка:", error);
        });
}

function signup() {
    const SIMPLE_EMAIL_REGEXP = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;

    const login = document.getElementById("login").value;
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const full_name = document.getElementById("full_name_user").value;
    const photo = document.getElementById("photo").value;
    const confirmPassword = document.getElementById("confirmPassword").value;
    const errorMessage = document.getElementById("errorMessage");

    errorMessage.textContent = "";
    if (!login || !email || !full_name || !password)
    {
        errorMessage.style.animation = "none"
        errorMessage.offsetHeight;
        errorMessage.style.animation = "shake 0.3s"
        errorMessage.textContent = "Заполните полную форму";
        return;
    }



    if (!SIMPLE_EMAIL_REGEXP.test(email)) {
        errorMessage.textContent = "Некорректный формат email";
        return;
    }

    if (password !== confirmPassword) {
        errorMessage.textContent = "Пароли не совпадают!";
        return;
    }

    const user = {
        login: login,
        email: email,
        password: password,
        full_name: full_name,
        photo: photo
    };

    fetch("/sign-up", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(user)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error("Ошибка регистрации");
            }
            return response.json();
        })
        .then(data => {
            alert("Регистрация успешна!");
        })
        .catch(error => {
            errorMessage.textContent = error.message;
        });
}

window.onload = function() {
    const token = localStorage.getItem("token");
    const profile = document.getElementById("profile");


    if (token) {
        console.log("Токен найден, загружаем профиль...");
        loadProfile();
    } else {
        profile.style.display = "none";
        console.log("No token found. Please log in.");
    }
};

function toggleDropdown() {
    const dropdownMenu = document.getElementById('dropdown-menu');
    dropdownMenu.classList.toggle('show');
}

window.addEventListener('click', function(event) {
    const dropdownMenu = document.getElementById('dropdown-menu');
    const profile = document.getElementById('profile');
    if (!profile.contains(event.target)) {
        dropdownMenu.classList.remove('show');
    }
});