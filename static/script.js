function calculate() {
    const num1 = document.getElementById("number1").value;
    const num2 = document.getElementById("number2").value;

    fetch("/add?num1=" + num1 + "&num2=" + num2)
        .then(response => response.json())
        .then(data => {
            document.getElementById("result").innerText = data.result;
        })
        .catch(error => {
            console.error("Erreur:", error);
        });
}
