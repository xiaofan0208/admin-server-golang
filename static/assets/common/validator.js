

    function SDValidator (form  , fields , redirectUrl){
            $(form).bootstrapValidator({
                message: 'This value is not valid',
                //提供输入验证图标提示
                feedbackIcons: {
                    valid: 'glyphicon glyphicon-ok',
                    invalid: 'glyphicon glyphicon-remove',
                    validating: 'glyphicon glyphicon-refresh'
                },
                fields : fields
            }).on('success.form.bv', function(e) {//点击提交之后
                e.preventDefault();
                var $form = $(e.target);
                var bv = $form.data('bootstrapValidator');
        
      
                var jsondata = $form.serializeJSON()
                var formdata1 = new FormData()
                for(var p in jsondata){
                    formdata1.append(p ,jsondata[p] )
                }
                for (var [a, b] of formdata1.entries()) {
                    console.log("name = " + a, " ;  value = " + b);
                }   
 
                var action = $form.attr("action")
                $.ajax({
                    type:"post",
                    url: action ,
                    data:formdata1,
                    async: false,//同步上传
                    cache: false,//上传文件无需缓存
                    processData: false,  // 不处理数据
                    contentType: false, // 不设置内容类型
                    success:function (data) {
                        console.log("------success = " + JSON.stringify(data));
                        if(data["code"] == 0 ){
                            window.location.href = redirectUrl
                        }else{
                            layer.alert( data["msg"] , { icon: 2, title: "失败" })
                        }
                    },
                    error:function (data) {
                        layer.msg('请求失败，请刷新重试');
                    }
                })
                
            })
    }
