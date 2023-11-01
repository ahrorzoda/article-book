const getBooks = async () => {
    try {
        const response = await fetch("http://localhost:8080/book/cover", {
            method: 'GET'
        })
        const responseJson = await response.json();
        let tableBody = document.getElementById("tableBody");
        const data = responseJson.files;
        console.log(data)
        for (const book of data) {
            tableBody.innerHTML += `<li>
        <a href="/home/shohrukh/Рабочий стол/article/back_end/uploads/${book}" download="/home/shohrukh/Рабочий стол/article/back_end/uploads/${book}">
            <img src="/home/shohrukh/Рабочий стол/article/back_end/uploads/${book}" alt="${book}" >
        </a>
    </li>`
        }
    } catch (e) {
        console.log(e)
    }
}
getBooks();