var DOMAIN = window.location.protocol+'//'+window.location.host+'/';
$(function () {
    initialMenuSel();
    $('.menu li').click(function () {
        setMenuSel($(this));
        window.location.href = DOMAIN+$(this).attr('data-c');
    });
    function initialMenuSel() {
        var path = window.location.pathname;
        var menu = ['index','note','download','about','sarscov'];
        var flag = false;
        for (var i in menu){
            if(path.substring(1,menu[i].length+1).indexOf(menu[i]) !== -1){
                flag = true;
                break;
            }
        }
        setMenuSel($('.menu li').eq((flag?i:0)));
    }
    function setMenuSel(_this) {
        $('.menu li').removeClass('sel');
        _this.addClass('sel');
    }
});