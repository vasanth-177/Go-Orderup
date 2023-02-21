

var selectedHotelName
var selectedItemsList=[]
var userName="vasanth"
function onChangeOption(oEvent) {
    var selectedHotelType = document.getElementById('hotelType').value;
    getHotels(selectedHotelType);
}

function onHotelChange(oEvent) {
    selectedHotelName = document.getElementById("hotelDropDown").value;
}


function getHotels(selectedHotelType) {
    oTBody0.innerHTML = "";
    document.getElementById("hotelDropDown").innerHTML = "";
    fetch("http://localhost:8081/" + selectedHotelType,
        {
            method: 'GET',
            mode: 'cors',
            headers: {
                // 'Access-Control-Allow-Origin': '*'
            }
        }).then((response) => response.json())
        .then((result) => {
            var table = document.getElementById('oTBody0');
            const hotelDropDown = document.getElementById("hotelDropDown");
            result.forEach((eachResult) => {
                var row = table.insertRow(table.rows.length);
                var cell1 = row.insertCell(0);
                var cell2 = row.insertCell(1);
                cell1.innerHTML = eachResult.Hotelname;
                cell2.innerHTML = eachResult.Type;

                let option = document.createElement("option");
                option.setAttribute('value', eachResult.Hotelname);

                let optionText = document.createTextNode(eachResult.Hotelname);
                option.appendChild(optionText);

                hotelDropDown.appendChild(option);
            })
        })


}

function getItems() {
    selectedHotelName = document.getElementById('hotelDropDown').value;
    if(selectedHotelName){
    oTBody1.innerHTML = "";
    document.getElementById("itemDropDown").innerHTML = "";
    
    // var raw = JSON.stringify({
    //     "option": selectedHotelName
    // });

    var requestOptions = {
        method: 'GET',
        headers: {
            // 'Content-Type': 'application/json'
            // 'Access-Control-Allow-Origin': 'http://localhost:8082',
            // 'Access-Control-Request-Method': 'POST'
        },
        mode: 'cors'
    };

    fetch("http://localhost:8082/getItem?type="+selectedHotelName, requestOptions)
        .then((response) => response.json())
        .then((result) => {
            // gItems=result;
            var table = document.getElementById('oTBody1');
            const itemDropDown = document.getElementById("itemDropDown");
            result.forEach((eachResult) => {
                var row = table.insertRow(table.rows.length);
                var cell1 = row.insertCell(0);
                var cell2 = row.insertCell(1);
                cell1.innerHTML = eachResult.item;
                cell2.innerHTML = eachResult.price;

                let option = document.createElement("option");
                option.setAttribute('value', eachResult.item);

                let optionText = document.createTextNode(eachResult.item);
                option.appendChild(optionText);

                itemDropDown.appendChild(option);
            })

        })
        .catch(error => console.log('error', error));

    for (var i = 1; i <= 50; i++) {
        var select = document.getElementById("quantityDropDown");
        var option = document.createElement("OPTION");
        select.options.add(option);
        option.text = i;
        option.value = i;
    }

    }
}

function add_items() {
    var selectedItem=document.getElementById("itemDropDown").value;
    var selectedQuantity=document.getElementById("quantityDropDown").value;
    var jsonRequestString = '{"item":"'+selectedItem+'","h_name":"'+selectedHotelName+'","quantity":"'+selectedQuantity+'"}';
    var obj = JSON.parse(jsonRequestString);
    selectedItemsList.push(obj);
    console.log(JSON.stringify(selectedItemsList));
}

function generate_bill(){
    let len = selectedItemsList.length;
  if(len>0)
  {
    
    var raw = JSON.stringify({
        "name": userName,
        "orderDetails":selectedItemsList
    });
    console.log(raw)
    var requestOptions = {
        method: 'POST',
        headers: {
            // 'Content-Type': 'application/json'
            // 'Access-Control-Allow-Origin': 'http://localhost:8082',
            // 'Access-Control-Request-Method': 'POST'
        },
        body: raw,
        mode: 'cors',
        responseType: 'arraybuffer'
    };

    fetch("http://localhost:8083/generateBill", requestOptions)
     .then((response) => response.arrayBuffer())
     .then((data) => {
       
        // fs.writeFileSync('pdfs/attachment.pdf', data);
        console.log(data);
       
     });

     
  
// .then(response => response.text())
// .then(result => console.log(result))
// .catch(error => console.log('error', error));

//   window.location.replace("info.html");
//   setTimeout("location.href='info.html';",3000);
  }
}


