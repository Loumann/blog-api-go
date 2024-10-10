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
                document.getElementById("message").innerText =  "Неверный логин или пароль";
                window.location.href = "/users";

            }
        } else if (xhr.readyState === 4) {
            document.getElementById("message").innerText = "Ошибка при входе";
        }
    }
    console.log(json)
}


const createPost = () => {
    const post = document.createElement("div")
    post.className = "post-on-title-box"
    post.append(createPost("Коммент") )

}


 function loadPosts() {
    fetch('/post', {
        method: 'GET',
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка сети: ' + response.status); // Обрабатываем ошибки
            }
            return response.json();
        })
        .then(data => {
            console.log('Данные получены:', data); //

        })
        .catch(error => {
            console.error('Ошибка:', error);
        });
}

loadPosts()