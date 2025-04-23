document.addEventListener("DOMContentLoaded", () => {
    const searchInput = document.getElementById("searchInput");

    async function searchUsers(login) {
        const userList = document.getElementById("userList");

        if (!login.trim()) {
            userList.innerHTML = "";
            return;
        }

        try {
            const response = await fetch(`/users/${login}`);
            const data = await response.json();

            userList.innerHTML = "";

            for (const user of data) {
                const userItem = document.createElement("div");
                userItem.classList.add("user-item");

                const photoSrc = user.photo && user.photo !== "null" ? user.photo : "static/PhotoBase/default.jpg";

                userItem.innerHTML = `
                <img src="${photoSrc}" alt="Фото профиля" style="width: 50px; height: 50px; border-radius: 50%; margin-right: 10px;">
                <h3>${user.full_name}</h3>
                <p><strong>Логин:</strong> ${user.login}</p>
                <p><strong>Email:</strong> ${user.email}</p>
            `;

                const subscribeButton = document.createElement("button");
                subscribeButton.classList.add("subscribeButton");

                // 🔍 Проверка подписки
                const isSubscribed = await checkSubscription(user.id);
                subscribeButton.textContent = isSubscribed ? "Отписаться" : "Подписаться";

                subscribeButton.onclick = () => subscribeUser(user.id, subscribeButton);

                userItem.appendChild(subscribeButton);
                userList.appendChild(userItem);
            }

        } catch (error) {
            console.error("Ошибка поиска:", error);
            const userItem = document.createElement("div");
            userItem.innerHTML = `<p>Пользователь не найден</p>`;
            userList.appendChild(userItem);
        }
    }


    async function subscribeUser(userId, button) {
        try {
            const response = await fetch(`http://localhost:8080/subscribe/${userId}`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ userId })
            });

            if (response.ok) {
                button.textContent = button.textContent === "Подписаться" ? "Отписаться" : "Подписаться";
                alert("Подписка изменена!");
            } else {
                alert("Произошла ошибка при подписке.");
            }
        } catch (error) {
            console.error("Ошибка подписки:", error);
        }
    }

    // 👇 Только по Enter
    searchInput.addEventListener("keydown", (e) => {
        if (e.key === "Enter") {
            e.preventDefault();
            searchUsers(e.target.value);
        }
    });
});


async function checkSubscription(userId) {
    try {
        const response = await fetch(`http://localhost:8080/check-sub/${userId}`);
        if (!response.ok) return false;

        const result = await response.json();
        return result.subscribed; // true / false
    } catch (error) {
        console.error("Ошибка проверки подписки:", error);
        return false;
    }
}
