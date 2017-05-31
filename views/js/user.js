var UPMS = {}
UPMS.Post = function (url, data, callback) {
    $.ajax({
        url: url,
        type: "POST",
        data: data,
        timeout: 1000,
        dataType: "json",
        success: callback
    });
}

UPMS.Get = function (url, data, callback) {
    $.ajax({
        url: url,
        type: "Get",
        data: data,
        timeout: 1000,
        dataType: "json",
        success: callback
    });
}

var upms = angular.module("upms", [])
upms.controller("UserController", function ($scope) {
    $scope.createuserdata = {};
    $scope.deleteuser = function () {
        var deleteids = [];
        var users = $scope.users;
        for (var i = 0; i < users.length; i++) {
            if (users[i].checked) {
                deleteids.push(users[i].ID)
            }
        }
        if (deleteids.length > 0) {
            console.debug(deleteids);
            var postdata = {};
            postdata.ids = deleteids;
            console.debug(deleteids.toLocaleString())
            UPMS.Post("/upms/v1/user/delete", deleteids.toLocaleString(), function (data, status) {
                if ("success" == status) {
                    console.debug(data);
                    if (data.status == 0) {
                        alert("用户删除成功！");
                        $.getJSON("/upms/v1/user", function (data, status) {
                            if ("success" == status) {
                                $scope.users = data;
                                for (var i = 0; i < data.length; i++) {
                                    console.debug(data[i])
                                }
                                $scope.$apply();
                            }
                        })
                    } else {
                        alert("用户删除失败！" + data.info);
                    }
                } else {
                    alert("用户删除失败！" + data.info);
                }
            });
        }
    };
    $scope.createuser = function () {
        UPMS.Post("/upms/v1/user/create", $scope.createuserdata, function (data, status) {
            if ("success" == status) {
                console.debug(data);
                if (data.status == 0) {
                    alert("保存成功！");
                    $('#createUser').modal('hide')
                    $.getJSON("/upms/v1/user", function (data, status) {
                        if ("success" == status) {
                            $scope.users = data;
                            for (var i = 0; i < data.length; i++) {
                                console.debug(data[i])
                            }
                            $scope.$apply();
                        }
                    })
                } else {
                    alert("保存失败！" + data.info);
                }
            } else {
                alert("保存失败！" + data.info);
            }
        })
    };
    $.getJSON("/upms/v1/user", function (data, status) {
        if ("success" == status) {
            $scope.users = data;
            for (var i = 0; i < data.length; i++) {
                console.debug(data[i])
            }
            $scope.$apply();
        }
    });
});

$(function () {
    var selected_platforms = []

    $('#roleselect').multiSelect();
    //$('#platformselect').multiSelect();

    $('#platformselect').multiSelect({
        afterSelect: function (values) {
            selected_platforms = selected_platforms.concat(values)
            console.debug(selected_platforms);
            $("#roleselect").empty();
            $('#roleselect').multiSelect('refresh');
            $.post("/upms/v1/roles", selected_platforms.toLocaleString(), function (data, status) {
                if (status == "success" && data.status == 0 && data.data != null) {
                    for (var i = 0; i < data.data.length; i++) {
                        $('#roleselect').multiSelect('addOption', {
                            value: data.data[i].ID,
                            text: data.data[i].Role + "(" + data.data[i].Platform + ")",
                            index: 0
                        });
                    }
                }
            });
        },
        afterDeselect: function (values) { 
            for (var i = 0; i < values.length; i++) {
                selected_platforms.splice(jQuery.inArray(values[i], selected_platforms), 1);
            } 
            $("#roleselect").empty();
            $('#roleselect').multiSelect('refresh');
            $.post("/upms/v1/roles", selected_platforms.toLocaleString(), function (data, status) {
                if (status == "success" && data.status == 0 && data.data != undefined && data.data != null) {
                    console.debug(data)
                    for (var i = 0; i < data.data.length; i++) {
                        $('#roleselect').multiSelect('addOption', {
                            value: data.data[i].ID,
                            text: data.data[i].Role + "(" + data.data[i].Platform + ")",
                            index: 0
                        });
                    }
                }
            });
        }
    });
    $.get("/upms/v1/platforms", function (data, status) {
        console.debug(data)
        if (status == "success" && data.status == 0) {
            for (var i = 0; i < data.data.length; i++) {
                $('#platformselect').multiSelect('addOption', {
                    value: data.data[i],
                    text: data.data[i],
                    index: 0
                });
            }
        }
    });
})

function newRole() {
    var role = $("#newRoleText").val();
    $('#roleselect').multiSelect('addOption', {
        value: role,
        text: role,
        index: 0
    });
};

function newPlatform() {
    var platform = $("#newPlatformText").val();
    $('#platformselect').multiSelect('addOption', {
        value: platform,
        text: platform,
        index: 0
    });
};
