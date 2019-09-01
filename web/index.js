const btnCSV = document.getElementById("btn-csv");
const btnSend = document.getElementById("btn-send");

// Get the modal
var modal = document.getElementById("myModal");

// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];

// When the user clicks on the button, open the modal
btnCSV.onclick = function() {
  modal.style.display = "block";
};

// When the user clicks on <span> (x), close the modal
span.onclick = function() {
  modal.style.display = "none";
};

// When the user clicks anywhere outside of the modal, close it
window.onclick = function(event) {
  if (event.target == modal) {
    modal.style.display = "none";
  }
};

const removeX = file => file.split(".")[0];

var fileName = "";

const handleFiles = file => {
  fileName = removeX(file[0].name);
  console.log(fileName);
};

let myHeaders = new Headers();
myHeaders.append("Content-Type", "application/json");

btnSend.onclick = e =>
  fetch("https://qxwfstxqv5.execute-api.us-east-1.amazonaws.com/dev/check", {
    method: "POST",
    mode: "cors",
    cache: "default",
    headers: myHeaders,
    body: JSON.stringify({ file: fileName })
  }).then(response => {
    if (response.ok) {
      alert("OK");
    } else {
      alert("NOT OK");
    }
    console.log(response.json());
  });
