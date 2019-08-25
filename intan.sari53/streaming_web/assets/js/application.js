require("bootstrap/dist/js/bootstrap.bundle.js");

$(()=>{
    $(window).on('resize',function(){
        if (this.window.innerWidth>991) {
            $(".drop-navigation.drop-navigation").css("display","block");
        }
    });
})