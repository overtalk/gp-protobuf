import { Msg } from "./protocol";

function post(){
    var data = getData()
    console.log(data)
    send(data)
}

function getData() {
    var json = "JSON.parse(JSON.stringify(this.model.list))"
    return json
}


function send(data) {
    $.ajax({
        type: "POST",
        url: "http://172.26.164.74:9090/newgun",
        contentType: "application/json; charset=utf-8",
        data: data,
        dataType: "text",
        async: true,
        error: function(request) {
            alert("send fail")
        },
        success: function(data, statu) {
            alert("send success")
        }
    });
}

