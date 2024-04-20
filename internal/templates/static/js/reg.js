<script>
    document.getElementById("myForm").addEventListener("submit", function(event) {
        event.preventDefault(); // Предотвращаем отправку формы по умолчанию
    var formData = new FormData(this); // Создаем объект FormData для сбора данных формы
    var requestData = { }; // Создаем объект для хранения данных формы
    formData.forEach(function(value, key) {
        requestData[key] = value; // Заполняем объект данными из формы
        });
    // Отправляем данные на сервер
    fetch("/submit", {
        method: "POST",
    headers: {
        "Content-Type": "application/json"
            },
    body: JSON.stringify(requestData)
        }).then(response => {
            // Обрабатываем ответ от сервера
            if (!response.ok) {
                throw new Error("Ошибка при отправке данных");
            }
    return response.json();
        }).then(data => {
        // Обрабатываем данные от сервера (если необходимо)
        console.log(data);
        }).catch(error => {
        // Обрабатываем ошибки
        console.error(error);
        });
    });
</script>
