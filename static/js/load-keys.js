$(document).ready(function () {
    $.ajax({
        type: "get",
        url: "/get-omise-public-key",

        contentType: "application/json;charset=utf-8",
        success: function (data) {
            alert("data");
        },
        error: function (data) {
            alert(data);
        }
    });
});

