function formatDate(isoString) {
    const date = new Date(isoString);

    return date.toLocaleString("ru-RU", {
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
    fetch(`http://localhost:8080/post/?page=${page}`)
        .then(response => response.json())
        .then(data => {
            const posts = data.posts;
            posts.forEach(post => {
                const postElement = document.createElement('div');
                postElement.classList.add('post');

                let shortText = post.content_post.slice(0, 500);
                let isLong = post.is_long;
                let formattedDate = formatDate(post.date_create);

                let postContent = `
                <div class="post-header">
                   <img src="${post.photo}" >
                   <div class="post-user-info">
                       <span class="fullname">${post.fullname}</span>
                       <span class="login">@${post.login}</span>
                   </div>
               </div>
               <h3>${post.theme}</h3>
               <p class="short-text">${shortText}${isLong ? '...' : ''}</p>`;

                if (isLong) {
                    postContent += `
                        <p class="full-text" style="display: none;">${post.content_post}</p>
                        <button class="toggle-btn">Показать все</button>`;
                }

                postContent += `<p class="time">${formattedDate}</p>`;

                postElement.innerHTML = postContent;
                postsContainer.appendChild(postElement);

                if (isLong) {
                    const toggleBtn = postElement.querySelector(".toggle-btn");
                    const fullText = postElement.querySelector(".full-text");
                    const shortTextEl = postElement.querySelector(".short-text");

                    toggleBtn.addEventListener("click", function () {
                        if (fullText.style.display === "none") {
                            fullText.style.display = "block";
                            shortTextEl.style.display = "none";
                            toggleBtn.textContent = "Скрыть";
                        } else {
                            fullText.style.display = "none";
                            shortTextEl.style.display = "block";
                            toggleBtn.textContent = "Показать все";
                        }
                    });
                }
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

loadPosts();
