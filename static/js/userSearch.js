document.addEventListener("DOMContentLoaded", () => {
    const searchInput = document.getElementById("search");
    const userList = document.getElementById("userList");

    function renderUsers() {
        userList.innerHTML = "";
        const searchValue = searchInput.value.toLowerCase();
        const filteredUsers = mockUsers.filter(user => user.login.toLowerCase().includes(searchValue));

        if (filteredUsers.length === 0) {
            userList.innerHTML = "<p class='no-users'>Пользователи не найдены</p>";
            return;
        }

        filteredUsers.forEach(user => {
            const userItem = document.createElement("div");
            userItem.classList.add("user-item");

            const userName = document.createElement("span");
            userName.textContent = user.login;

            const button = document.createElement("button");
            button.textContent = user.subscribed ? "Отписаться" : "Подписаться";
            button.classList.add(user.subscribed ? "unsubscribe-btn" : "subscribe-btn");
            button.onclick = () => {
                user.subscribed = !user.subscribed;
                renderUsers();
            };

            userItem.appendChild(userName);
            userItem.appendChild(button);
            userList.appendChild(userItem);
        });
    }

    window.onload = function() {
        document.getElementById('subscribeButton').addEventListener('click', renderUsers);

        document.addEventListener('keydown', function(event) {
            if (event.key === 'Enter') {
                renderUsers();
            }
        });
    };


    searchInput.addEventListener("input", renderUsers);
    renderUsers();
});