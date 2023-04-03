var mainUrl ="http://127.0.0.1:5555/"
var currentDirectory

// 
var viewDAta = function(data){
  console.log(data);
}
var viewDataNub = function (data) {
    document.getElementById("xAxis").innerHTML = "<option> none </option>"
    document.getElementById("yAxis").innerHTML = "<option> none </option>"
    
    let body = ""
    let head = "<tr>"
    for (let i=0; i< data.Header.length; i++) {
      head += "<th style='text-align: center'> " + data.Header[i] + "</th>"
      document.getElementById("xAxis").innerHTML += "<option>" + data.Header[i] + "</option>"
      document.getElementById("yAxis").innerHTML += "<option>" + data.Header[i] + "</option>"
    }
    head += "</tr>"
    document.getElementById("pHeader").innerHTML = head
    for (let i = 0; i < data.Data.length; i++) {
        body += "<tr>"
        for (let key in data.Data[i]) {
        body += "<td>" + data.Data[i][key] + "</td>"
      }
      body += "</tr>"
    }
    document.getElementById("pData").innerHTML = body
  }

var getData = function (){
  let xAxis = document.getElementById("xAxis")
  let yAxis = document.getElementById("yAxis")
  if ((xAxis.value== "none") || (yAxis.value=="none")){
    console.log("none");
    return
  };  
  console.log(xAxis.value,yAxis.value);
}
var processDirectory = function (data) {
    currentDirectory = data
    document.getElementById("directoryLabel").innerHTML = "Directory: " + currentDirectory.Home
    document.getElementById("directory").innerHTML = "<option>..</option>"
    for (i in currentDirectory.Sub_Directories) {
      document.getElementById("directory").innerHTML += "<option>" + currentDirectory.Sub_Directories[i] + "</option>"
    }

    document.getElementById("files").innerHTML = ""
    for (i in currentDirectory.Files) {
      if (currentDirectory.Files[i].endsWith(".csv")) {
        document.getElementById("files").innerHTML += "<option>" + currentDirectory.Files[i] + "</option>"
      }
    }
  }

  var selectDirectory = function () {
    var x = document.getElementById("directory").value
    if (x == "..") {
      let pos = currentDirectory.Home.lastIndexOf("/")
      getDirectory(currentDirectory.Home.slice(0, pos));
    } else {
      getDirectory(currentDirectory.Home + "/" + x)
    }
  }  

  var selectFile = function () {
    var x = document.getElementById("files").value
    getFile(currentDirectory.Home + "/" + x)
  }

  var getFile = function (dir) {
    var url = new URL(mainUrl+"getFile");
    url.searchParams.append("directory", dir);
    fetch(url)
      // Handle success
      .then((response) => response.json()) // convert to json
      .then((json) => viewDataNub(json)) //print data to console
      .catch((err) => console.log("Request Failed", err)); // Catch errors
  };
var getDirectory = function (dir) {
    var url = new URL(mainUrl+"getDirectory");
    url.searchParams.append("directory", dir);
    fetch(url)
      // Handle success
      .then((response) => response.json()) // convert to json
      .then((json) => processDirectory(json)) //print data to console
      .catch((err) => console.log("Request Failed", err)); // Catch errors
  };
  getDirectory("./")