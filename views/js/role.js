var zTree;

var setting = {
    view: {
        showLine: true,
        checkable: true
    },
    data: {
        key: {
            name: "Name",
        },
        simpleData: {
            enable: true,
            idKey: "ID",
            pIdKey: "ParentID",
            rootPId: ""
        }
    },
    callback: {
        beforeClick: function (treeId, treeNode) {
            if (treeNode.Roles != undefined) {
                console.debug(treeNode.Roles);
                zTree.removeChildNodes(treeNode);
                for (var i = 0; i < treeNode.Roles.length; i++) {
                    $("#users-in-group").append("<li class='list-group-item' style='text-align:center'>"+treeNode.Roles[i].Name+"</li>")
                    zTree.addNodes(treeNode, 0, treeNode.Roles[i], false);
                }
//                $.getJSON("/upms/v1/user/groupid", function (data, status) {
//                    var zNodes = data.data;
//                    t = $.fn.zTree.init($("#tree"), setting, zNodes);
//                });
            }else{
                $("#users-in-group").empty();
            }
        }
    }
};

$(document).ready(function () {
    $.getJSON("/upms/v1/rolegroups", function (data, status) {
        console.debug(data.data);
        var zNodes = data.data;
        zTree = $.fn.zTree.init($("#tree"), setting, zNodes);
    });

});
