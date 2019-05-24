function load(){
    getDevices();
    getPresets();
}

function toggleCheck(element) {
    var check = element.checked;
    var request = new XMLHttpRequest();
    request.open("POST", ("/Main/toggleDevice/?device=" + element.id + "&checked=" + check), true);
    request.onreadystatechange = function () {
        submitToggle(request);
    };
    request.send();
}

function submitToggle(request) {
    if ((request.readyState == 4) && (request.status == 200)) {

    }
}

function getDevices() {
    var request = new XMLHttpRequest();

    request.open("GET", ("/Main/getDevices"), true);
    request.onreadystatechange = function () {
        deviceHandler(request);
    }
    request.send();
}

function deviceHandler(request) {
    if ((request.readyState == 4) && (request.status == 200)) {
        var devices = document.getElementById("devices");
        devices.innerHTML = "";
        var responseArray = request.responseText;
        responseArray = responseArray.substring(1, responseArray.length);
        if (responseArray != "") {
            var elements = responseArray.split(',');
            for (var qty in elements) {
                var toBeUsed = elements[qty].trim();
                var info = toBeUsed.split('|');
                if (info[2] != "") {
                    devices.innerHTML += '<div class="device">' +
                        ' <label class="name">' + info[1] + '</label>' +
                        '<label class="switch">' +
                        '<label class="off">Off</label>' +
                        '<input type="checkbox" id="' + info[0] + '" onclick="toggleCheck(this)"' +
                        'checked>' +
                        '<span class="slider"></span>' +
                        '<label class="on">On</label>' +
                        '</label>' +
                        '</div>';
                } else {
                    devices.innerHTML += '<div class="device">' +
                        ' <label class="name" >' + info[1] + '</label>' +
                        '<label class="switch">' +
                        '<label class="off">Off</label>' +
                        '<input type="checkbox" id="' + info[0] + '" onclick="toggleCheck(this)"' +
                        '>' +
                        '<span class="slider"></span>' +
                        '<label class="on">On</label>' +
                        '</label>' +
                        '</div>';

                }


            }
        }
    }
}

function getPresets(){
    var request = new XMLHttpRequest();

    request.open("GET", ("/Main/presets"), true);
    request.onreadystatechange = function () {
        presetHandler(request);
    }
    request.send();
}

function presetHandler(request){
    if ((request.readyState == 4) && (request.status == 200)) {
        var presets = document.getElementById("presets");
        presets.innerHTML = "";
        var responseArray = request.responseText;

        if (responseArray != "") {
            response = responseArray.split(" ");
            
            for (var ind in response){
                presets.innerHTML += '<input type="button" id="'+ response[ind].trim() +'" value="'+ response[ind].trim() +'" onclick="preset(this)"/>';
            }
        }
    }
}

function preset(element){
    var request = new XMLHttpRequest();
    request.open("GET", ("/Main/presetPicked/?preset=" + element.id), true);
    request.onreadystatechange = function () {
        activatePreset(request);
    };
    request.send();
}

function activatePreset(request){
    if ((request.readyState == 4) && (request.status == 200)) {
        var responseArray = request.responseText;
        responseArray = responseArray.substring(1, responseArray.length-1);
        if (responseArray != "") {
            var elements = responseArray.split(" ");
            
            for (var qty in elements) {
                var toBeUsed = elements[qty].trim();
                var info = toBeUsed.split('|');
                var id = info[0];
                var checked = info[1];

                if (checked === "on"){
                    document.getElementById(id).checked = true;
                }else{
                    document.getElementById(id).checked = false;
                }
            }
        }
    }
}

function savePreset(){
    var request = new XMLHttpRequest();
    var ids = getCheckedIds();
    var presetName = prompt("What would you like to call this preset?", "Dhuhr");
    if (presetName != null && presetName != ""){
        var queryString = "?presetName=" + presetName;
        for (var ind in ids){
            queryString += "&ids=" + ids[ind];
        }
        
        request.open("POST", ("/Main/savePreset/" + queryString), true);
        request.onreadystatechange = function () {
            if ((request.readyState == 4) && (request.status == 200)) {
    
            }
        };
        request.send();
    }
}


function getCheckedIds(){
    return Array.from(document.querySelectorAll('input[type="checkbox"]'))
        .filter((checkbox) => checkbox.checked)
        .map((checkbox) => checkbox.id);
}