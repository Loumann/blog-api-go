
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
            document.getElementById("message").style.color = "red";
            document.getElementById("message").innerText = "Неверный логин или пароль";
        }
    };

    console.log(json);
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


function createPost() {
    let titleInput = document.getElementById("postTitle");
    let contentInput = document.getElementById("postContent");
    let title = titleInput.value.trim();
    let content = contentInput.value.trim();
    const token = localStorage.getItem("token");

    if (!title ) {
        titleInput.style.animation = "none";
        titleInput.offsetHeight;
        titleInput.style.animation = "shake 0.3s";
        titleInput.placeholder = "Заполните все поля!";
        return;

    } else if (!content) {
        contentInput.style.animation = "none";
        contentInput.offsetHeight;
        contentInput.style.animation = "shake 0.3s";
        contentInput.placeholder = "Заполните все поля!";
        return;
    }

    const post = {
        theme: title,
        content_post: content,
        date_create_post: new Date().toISOString()
    };

    fetch("/post", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
        },
        body: JSON.stringify(post)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error("Ошибка создания поста: " + response.status);
            }
            return response.json();
        })
        .then(data => {
            alert("Создание успешно!");
            window.location.reload();
            titleInput.value = "";
            contentInput.value = "";
            closeModal()
        })
        .catch(error => {
            alert(error.message);
        });
}

function editPost(button) {
    const postElement = button.closest(".post");

    if (!postElement) {
        console.error("Ошибка: Не найден родительский элемент .post");
        return;
    }

    const postId = postElement.getAttribute("data-id");

    const title = postElement.querySelector("h3")?.textContent || "Без заголовка";

    const contentElement = postElement.querySelector(".full-text") || postElement.querySelector(".short-text");
    const content = contentElement ? contentElement.textContent : "Нет контента";

    document.getElementById("edit-post-id").value = postId;
    document.getElementById("edit-title").value = title;
    document.getElementById("edit-content").value = content;

    document.getElementById("edit-modal").style.display = "flex";
}

async function savePost() {
    try {
        const postId = document.getElementById("edit-post-id").value;
        const title = document.getElementById("edit-title").value;
        const content = document.getElementById("edit-content").value;

        const post = {
            id_post: parseInt(postId),
            theme: title,
            content_post: content
        };

        const response = await fetch(`/post`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(post)
        });

        if (!response.ok) {
            throw new Error(`Ошибка запроса: ${response.status}`);
        }

        const result = response.headers.get("Content-Length") === "0" ? {} : await response.json();

        alert(result.message || result.error || "Пост обновлён!");

        if (response.ok) {
            location.reload();
        }
    } catch (error) {
        console.error("Ошибка сохранения поста:", error);
        alert("Ошибка сохранения поста: " + error.message);
    }
}


window.addEventListener('click', function(event) {
    const dropdownMenu = document.getElementById('dropdown-menu');
    const profile = document.getElementById('profile');
    if (!profile.contains(event.target)) {
        dropdownMenu.classList.remove('show');
    }
});


function closeModalEdit() {
    document.getElementById("edit-modal").style.display = "none";
}
function toggleDropdown() {
    const dropdownMenu = document.getElementById('dropdown-menu');
    dropdownMenu.classList.toggle('show');
}
function openModal() {
    document.getElementById("modalOverlay").style.display = "flex";
}
function closeModal() {

    document.getElementById("modalOverlay").style.display = "none";

    document.getElementById("postTitle").value = "";
    document.getElementById("contentInput").value= "";

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
window.addEventListener('click', function(event) {
    const dropdownMenu = document.getElementById('dropdown-menu');
    const profile = document.getElementById('profile');
    if (!profile.contains(event.target)) {
        dropdownMenu.classList.remove('show');
    }
});
