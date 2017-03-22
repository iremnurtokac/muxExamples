$(document).ready(function () {
    $("#button1").click(function (event) {
        $.ajax({
            type: 'GET',
            crossOrigin: true,
            url: 'http://localhost:8080/users/Tictactoe',
            success: function (result) {
                $("#target").html('<table class="table table-bordered"><tr>' +
                "<td>"+ result.field[0] + "</td>" +
                "<td>"+ result.field[1] + "</td>" +
                "<td>"+ result.field[2] + "</td><tr>" +
                "<td>"+ result.field[3] + "</td>" +
                "<td>"+ result.field[4] + "</td>" +
                "<td>"+ result.field[5] + "</td><tr>" +
                "<td>"+ result.field[6] + "</td>" +
                "<td>"+ result.field[7] + "</td>" +
                "<td>"+ result.field[8] + "</td>"
                +"</tr> </table><div>"+ result.winner +"</div>" 
                 );
            },
            error: function () {

            }
        });
    });
});


/* 
$(document).ready(function () {
    $("#button1").click(function (event) {
        $.ajax({
            type: 'GET',
            crossOrigin: true,
            url: 'http://localhost:8080/users/Tictactoe',
            success: function (result) {
                $("#target").html('<table class="table table-bordered"><tr>' +
                "<td>"+ result.field[0] + "</td>" +
                "<td>"+ result.field[1] + "</td>" +
                "<td>"+ result.field[2] + "</td><tr>" +
                "<td>"+ result.field[3] + "</td>" +
                "<td>"+ result.field[4] + "</td>" +
                "<td>"+ result.field[5] + "</td><tr>" +
                "<td>"+ result.field[6] + "</td>" +
                "<td>"+ result.field[7] + "</td>" +
                "<td>"+ result.field[8] + "</td>"
                //+"</tr> </table><div>"+ result.player +"</div>" 
                 );
            },
            error: function () {

            }
        });
    });
});

*/