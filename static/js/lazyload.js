function formatDate(isoString) {
    const date = new Date(isoString);
    return date.toLocaleString("ru-RU", {
        timeZone: "UTC",
        day: "numeric",
        month: "long",
        year: "numeric",
        hour: "2-digit",
        minute: "2-digit"
    });
}

let page = 1;
const postsContainer = document.getElementById('posts-container');
const loadingIndicator = document.getElementById('loading');

function loadPosts() {
    fetch(`/post/?page=${page}`)
        .then(response => response.json())
        .then(data => {
            const posts = data.posts;

            posts.forEach(post => {
                const postElement = document.createElement('div');
                postElement.classList.add('post');
                postElement.setAttribute('data-user-id', post.id_user_create_post);

                let shortText = post.content_post.slice(0, 500);
                let isLong = post.is_long;
                let formattedDate = formatDate(post.date_create);

                let postContent = `
                    <div class="post-header">
                        <img src="${post.photo}">
                        <div class="post-user-info">
                            <span class="fullname">${post.fullname}</span>
                            <span class="login">@${post.login}</span>
                        </div>
                    </div>
                    <h3>${post.theme}</h3>
                    <p class="short-text">${shortText}${isLong ? '...' : ''}</p>
                `;

                if (isLong) {
                    postContent += `
                        <p class="full-text" style="display: none;">${post.content_post}</p>
                        <button class="toggle-btn">Показать все</button>
                    `;
                }

                postContent += `
                    <button class="subscribe-btn">Подписаться</button>
                    <div id="comments-${post.id_post}" class="comments-container"></div>
                    <p class="time">${formattedDate}</p>
                    <form class="comment-form" data-post-id="${post.id_post}">
                      <div class="comment-box">
                        <textarea class="comment-text" placeholder="Напиши комментарий..."></textarea>
                        <button type="submit" class="comment-submit">Отправить</button>
                      </div>
                    </form>
                `;

                postElement.innerHTML = postContent;
                postsContainer.appendChild(postElement);

                // Показать весь текст
                if (isLong) {
                    const toggleBtn = postElement.querySelector(".toggle-btn");
                    const fullText = postElement.querySelector(".full-text");
                    const shortTextEl = postElement.querySelector(".short-text");

                    toggleBtn.addEventListener("click", function () {
                        const showAll = fullText.style.display === "none";
                        fullText.style.display = showAll ? "block" : "none";
                        shortTextEl.style.display = showAll ? "none" : "block";
                        toggleBtn.textContent = showAll ? "Скрыть" : "Показать все";
                    });
                }

                // Отправка комментария
                const commentForm = postElement.querySelector('.comment-form');
                commentForm.addEventListener('submit', function (e) {
                    e.preventDefault();
                    const postId = this.dataset.postId;
                    const commentText = this.querySelector('.comment-text').value;
                    submitComment(postId, commentText, this);
                });

                // Загрузка комментариев
                fetch(`/comment?post_id=${post.id_post}`)
                    .then(res => res.json())
                    .then(comments => {
                        const commentsContainer = document.getElementById(`comments-${post.id_post}`);
                        commentsContainer.innerHTML = "";

                        if (comments.length === 0) {
                            commentsContainer.innerHTML = "<p>Комментариев пока нет.</p>";
                            return;
                        }

                        comments.forEach(comment => {
                            const commentDiv = document.createElement("div");
                            commentDiv.classList.add("comment");
                            commentDiv.innerHTML = `
                                <div class="comment-user-head">
                                    <img class="comment-photo-user" src=${comment.photo}>
                                    <p><strong>@ ${comment.login}</strong></p>
                                </div>
                                <p>${comment.content}</p>
                                <p class="comment-date-create"><em>${formatDate(comment.date_create)}</em></p>
                            `;
                            commentsContainer.appendChild(commentDiv);
                        });
                    })
                    .catch(err => {
                        const commentsContainer = document.getElementById(`comments-${post.id_post}`);
                        console.error(err);
                    });

                const subscribeBtn = postElement.querySelector('.subscribe-btn');
                subscribeBtn.addEventListener('click', function () {
                    subscribeToUser(this, post.id_user_create_post);
                });
            });

            page++;
            const lastPost = postsContainer.querySelector('.post:last-child');
            observer.observe(lastPost);
        })
        .finally(() => {
            loadingIndicator.style.display = 'none';
        });
}

const observer = new IntersectionObserver((entries) => {
    const lastPost = entries[0];
    if (lastPost.isIntersecting) {
        loadPosts();
    }
}, {
    rootMargin: '100px',
    threshold: 0.5
});

async function submitComment(postId, commentText, formElement) {
    if (!commentText) {
        alert("Напишите комментарий");
        return;
    }

    const response = await fetch(`http://localhost:8080/comment/${postId}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ content: commentText })
    });

    if (response.ok) {
        formElement.querySelector('.comment-text').value = '';
        location.reload();
    } else {
        alert('Ошибка при отправке комментария');
    }
}

function subscribeToUser(button, userId) {
    const token = localStorage.getItem("token");

    if (!token) {
        alert("Вы не авторизованы!");
        return;
    }

    fetch(`/subscribe/${userId}`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
        }
    })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                alert("Ошибка: " + data.error);
            } else {
                alert(data.message || "Вы успешно подписались!");
                button.disabled = true;
                button.innerText = "Вы подписаны";
            }
        })
        .catch(error => {
            console.error("Ошибка подписки:", error);
            alert("Ошибка подписки: " + error.message);
        });
}




loadPosts();
