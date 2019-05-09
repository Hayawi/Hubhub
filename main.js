function toggleCheck(element){
    var check = element.checked;
	var request = new XMLHttpRequest();
	request.open("POST", ("?device=" + element + "&checked=" + check), true);
	request.onreadystatechange = function() {
		submitToggle(request);
	};
	request.send();
}

function submitToggle(request){
    if ((request.readyState == 4) && (request.status == 200)){
		
	}
}

function getDevices(){
    var request = new XMLHttpRequest();

	request.open("GET", ("?checks=true"), true);
	request.onreadystatechange = function(){
		deviceHandler(request);
	}
	request.send();
}

function deviceHandler(request){
    if ((request.readyState == 4) && (request.status == 200)){
        var devices = document.getElementById("devices");
        devices.innerHTML = "";
        var responseArray = request.responseText;
        if (responseArray != ""){
            var elements = responseArray.split(',');
            for (var qty in elements){
                var toBeUsed = elements[qty].trim();
                var info = toBeUsed.split('|');
                devices.innerHTML += '<div id="' + info[0] +'" class="device">'
                +' <label class="name" id="' + info[0] +'">' + info[1] + '</label>'
                + '<label class="switch">'
                + '<label class="off">Off</label>'
                + '<input type="checkbox" id="' + info[0] +'" onclick="toggleCheck(this)">'
                + '<span class="slider"></span>'
                + '<label class="on">On</label>'
                + '</label>'
                + '</div>'
            }
        }
    }
}