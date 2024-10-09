console.log("js file download!")


function SignIn(){

     let xrh = new XMLHttpRequest()
     xrh.open("GET", 'http://localhost:8080/users')
     xrh.responseType = "json"

     xrh.send()
    xrh.onload = function (){
         let responseJSON = xrh.response;
         alert(responseJSON.message);
         console.log(responseJSON.message, responseJSON)

    }
}

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