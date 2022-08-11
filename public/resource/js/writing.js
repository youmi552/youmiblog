var CONTENT_KEY = "CACHE_CONTENT"; // 编辑器内容缓存key
var TITLE_KEY = "CACHE_TITLE"; // 标题缓存key
var AUTO_SAVE_TIME = 5000; // 自动保存时间
var cos = null;
var MdEditor = null;
var headInput = null;
var ArticleItem = {};

function setAjaxToken(xhr) {
  xhr.setRequestHeader("Authorization", localStorage.getItem("AUTH_TOKEN"));
}
function initEditor() {
  if (ArticleItem.PostMore){
    initEditor2()
    return
  }

  // 取默认标题
    headInput.val(ArticleItem.title);
  // 初始化编辑器
  MdEditor = editormd("editormd", {
    width: "99.5%",
    height: window.innerHeight - 78,
    syncScrolling: "single",
    editorTheme: "default",
    path: CNDURL + "/lib/",
    placeholder: "",
    appendMarkdown: ArticleItem.markdown,
    codeFold: true,
    saveHTMLToTextarea: true,
    // tocm: true,
    imageUpload: true,
    taskList: true,
    // emoji: true,
    imageFormats: ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
    // imageUploadURL: "/api/v1/uploadfile",
    // imageUploadCalback: function (files, cb) {
    //   uploadImage(files[0], cb);
    // },
    imageUploadCalback:function (files,cb){
      uploadImage2(files[0],cb);
    },
  });
}
function uploadImage2(file,cb){
  //将form的数据放入formData中
  let formData = new FormData()
  formData.append("f1",file)
  $.ajax({
    url:'/uploadfile',
    method:"post",
    data:formData,
    // 告诉jQuery不要去处理发送的数据，用于对data参数进行序列化处理 这里必须false
    processData : false,
    // 告诉jQuery不要去设置Content-Type请求头
    contentType : false,
    success:function (res) {
      if (res.code !== 200) return alert(res.msg);
      cb(res.data)
    }
  })
}
function uploadImage(file, cb) {
  const config = {
    useCdnDomain: true,
    //自行去改七牛云的空间区域的配置
    region: qiniu.region.z1
  };
  const putExtra = {
  };
  // 异步获取临时密钥
  $.ajax({
    url: "/api/v1/qiniu/token",
    type: "GET",
    contentType: "application/json",
    success: function (res) {
      if (res.code !== 200) return alert(res.error);
      const token = res.data;
      const observable = qiniu.upload(file, "goblog/upload/"+Date.now() + "_" + file.name, token, putExtra, config)
      const observer = {
        next(res){
          // ...
        },
        error(err){
          // ...
        },
        complete(res){
          console.log(res)
          cb("https://static.mszlu.com/" + res.key)
        }
      }
      const subscription = observable.subscribe(observer) // 上传开始

    },
    beforeSend: setAjaxToken,
  });

}


function getArticleItem(id) {
  $.ajax({
    url: "/writing/" + id,
    type: "GET",
    contentType: "application/json",
    success: function (res) {
      if (res.code != 200) {
        initEditor();
        return alert("初始化失败！");
      }
      ArticleItem = res.data || {};
      ArticleItem.pid=ArticleItem.PostMore.pid
      ArticleItem.categoryId=ArticleItem.PostMore.Category[0].categoryId
      initActive();
      initEditor();
    },
    beforeSend: setAjaxToken,
  });
}
function initActive() {
  $(".category li[value=" + ArticleItem.PostMore.Category[0].categoryId + "]")
    .addClass("active")
    .siblings()
    .removeClass("active");
  $(".type-box li[value=" + ArticleItem.PostMore.type + "]")
    .addClass("active")
    .siblings()
    .removeClass("active");
  $(".slug-input").val(ArticleItem.PostMore.slug);
}
function initCache() {
  headInput = $(".header-input");
  var query = new URLSearchParams(location.search);
  var _id = query.get("id");
  if (_id) return getArticleItem(_id);
  // 取本地缓存
  ArticleItem.title = window.localStorage.getItem(TITLE_KEY);
  ArticleItem.markdown = window.localStorage.getItem(CONTENT_KEY) || "";
  // initEditor
  initEditor();
  // 自动保存
  setInterval(() => saveHandler, AUTO_SAVE_TIME);
}

function saveHandler() {
  window.localStorage.setItem(TITLE_KEY, headInput.val());
  window.localStorage.setItem(CONTENT_KEY, MdEditor.getMarkdown());
}
function clearHandler() {
  window.localStorage.removeItem(TITLE_KEY);
  window.localStorage.removeItem(CONTENT_KEY);
}

// 发布
function publishHandler() {
  if (!ArticleItem.categoryId) return $(".publish-tip").text("请选择分类");
  ArticleItem.slug = $(".slug-input").val();
  if (ArticleItem.type == 1 && !ArticleItem.slug)
    return $(".publish-tip").text("请输入自定义链接");
  ArticleItem.title = headInput.val();
  if (!ArticleItem.title) return $(".publish-tip").text("请输入标题");
  ArticleItem.markdown = MdEditor.getMarkdown();
  if (!ArticleItem.markdown) return $(".publish-tip").text("正文");
  ArticleItem.content = MdEditor.getPreviewedHTML();
  $.ajax({
    url: "/writing/",
    type: ArticleItem.pid ? "PUT" : "POST",
    contentType: "application/json",
    data: JSON.stringify(ArticleItem),
    success: function (res) {
      if (res.code==200)alert(res.msg)
      if (res.code !== 200) return alert(res.msg);
      if (ArticleItem.pid) return $(".publish-tip").text(res.msg);
      ArticleItem = res.data || {};
      if (!ArticleItem.pid) {
        clearHandler();
      }
      location.search = "?id=" + ArticleItem.PostMore.pid;
    },
    beforeSend: setAjaxToken,
  });
}

$(function () {
  // 初始化缓存
  initCache();
  // 返回首页
  var back = $(".home-btn");
  back.click(function () {
    saveHandler();
    location.href = ArticleItem.pid ? "/p/" + ArticleItem.pid : "/";
  });
  if (location.search) back.text("返回");
  // 保存
  $(".save-btn").click(saveHandler);
  var drop = $(".publish-drop");
  // 显示
  $(".publish-show").click(function () {
    drop.show();
  });
  // 隐藏
  $(".publish-close").click(function () {
    drop.hide();
  });
  $(".cancel-btn").click(function () {
    drop.hide();
  });
  // 发布逻辑
  $(".publish-btn").click(publishHandler);
  // 选择分类
  $(".category").on("click", "li", function (event) {
    var target = $(event.target);
    target.addClass("active").siblings().removeClass("active");
    ArticleItem.categoryId = target.attr("value");
    $(".publish-tip").text("");
  });
  // 选择类型
  $(".type-box").on("click", "li", function (event) {
    var target = $(event.target);
    target.addClass("active").siblings().removeClass("active");
    ArticleItem.type = Number(target.attr("value") || 0);
  });
});
function initEditor2() {
  // 取默认标题
  headInput.val(ArticleItem.PostMore.title);
  // 初始化编辑器
  MdEditor = editormd("editormd", {
    width: "99.5%",
    height: window.innerHeight - 78,
    syncScrolling: "single",
    editorTheme: "default",
    path: CNDURL + "/lib/",
    placeholder: "",
    appendMarkdown: ArticleItem.PostMore.markdown,
    codeFold: true,
    saveHTMLToTextarea: true,
    // tocm: true,
    imageUpload: true,
    taskList: true,
    // emoji: true,
    imageFormats: ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
    // imageUploadURL: "/api/v1/uploadfile",
    imageUploadCalback: function (files, cb) {
      uploadImage2(files[0], cb);
    },
  });
}