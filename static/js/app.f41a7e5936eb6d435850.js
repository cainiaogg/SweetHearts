webpackJsonp([0],[,,function(t,e,a){"use strict";var n=a(0),i=a(29),s=a(24),o=a.n(s),r=a(23),l=a.n(r);n.default.use(i.a),e.a=new i.a({routes:[{path:"/",name:"friendCircle",component:o.a},{path:"/add",name:"Add",component:l.a}]})},function(t,e){},function(t,e){},function(t,e){},function(t,e){},,,,function(t,e,a){a(19);var n=a(1)(a(12),a(27),null,null);t.exports=n.exports},,function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default={name:"app"}},function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default={name:"GHeader",props:{hname:{default:"admin"}},data:function(){return{ishoney:0}},methods:{Post:function(t){console.log(t),0==t&&(this.ishoney=0),1==t&&(this.ishoney=1),this.$emit("post",this.ishoney),console.log("被触发")},reback:function(){this.$router.push("/?username"+this.hname),console.log("返回了")}}}},function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var n=a(22),i=a.n(n);e.default={components:{GHeader:i.a},created:function(){var t=this;console.log("wer"),this.$http.get("http://118.89.200.173/circle/get_id").then(function(e){console.log(e),t.addParam.id=e.data,t.param.id=e.data});var e=/username=.*$/g,a=location.href.match(e)[0].split("=")[1];this.addParam.username=this.param.username=a},data:function(){return{defaultList:[],imgName:"",visible:!1,uploadList:[],param:{username:"123",id:"1"},addParam:{username:"",id:"",ishoney:"0"},content:""}},methods:{report:function(t){var e=this;console.log(t),this.addParam.ishoney=t;var a=this;this.$http.post("http://118.89.200.173/circle/add_friend_circle",{username:this.addParam.username,id:this.addParam.id,content:this.content,ishoney:this.addParam.ishoney}).then(function(t){console.log(t),a.$router.push("/?username="+e.addParam.username)},function(t){console.log(t)})},handleView:function(t){this.imgName=t,this.visible=!0},handleRemove:function(t){var e=this.$refs.upload.fileList;this.$refs.upload.fileList.splice(e.indexOf(t),1)},handleSuccess:function(t,e){e.url=t[1],e.name=t[0],console.log(t),console.log(e),console.log(this.defaultList)},handleFormatError:function(t){this.$Notice.warning({title:"文件格式不正确",desc:"文件 "+t.name+" 格式不正确，请上传 jpg 或 png 格式的图片。"})},handleMaxSize:function(t){this.$Notice.warning({title:"超出文件大小限制",desc:"文件 "+t.name+" 太大，不能超过 2M。"})},handleBeforeUpload:function(){this.imgName=((new Date).getTime()+Math.random()).toString(),console.log(this.name);var t=this.uploadList.length<5;return t||this.$Notice.warning({title:"最多只能上传 5 张图片。"}),t}},mounted:function(){this.uploadList=this.$refs.upload.fileList}}},function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default={created:function(){var t=/username=.*$/g;console.log(location.href.match(t));var e=location.href.match(t)[0].split("=")[1];this.username=e,this.initData(1)},methods:{goAdd:function(){this.$router.push("/add?username="+this.username)},initData:function(t){var e=this.username,a=t||this.page,n="http://118.89.200.173/circle/get_friend_circle?username="+e+"&page="+a,i=this;1==a&&(this.friendCircleList=[]),this.$http.get(n).then(function(t){console.log(t.body),t.body.forEach(function(t,e){var a={};a.Content=t.Content,a.picList=t.ImagePath,a.UserName=t.UserName,t.TimeStamp=1e3*t.TimeStamp,a.TimeStamp=new Date(t.TimeStamp).getMonth()+1+"月"+new Date(t.TimeStamp).getDate()+"日"+new Date(t.TimeStamp).getHours()+":"+(new Date(t.TimeStamp).getMinutes()<10?"0"+new Date(t.TimeStamp).getMinutes():new Date(t.TimeStamp).getMinutes()),a.headerUrl=t.HeadImage?t.HeadImage:"https://o5wwk8baw.qnssl.com/a42bdcc1178e62b4694c830f028db5c0/avatar",a.ishoney=t.IsHoney,i.friendCircleList.push(a)}),i.loading=!1})},loadTop:function(){this.$refs.loadmore.onTopLoaded(),this.initData(1)},loadBottom:function(){this.loading=!0;var t=this.page++;this.initData(t)}},data:function(){return{friendCircleList:[],username:"admin",page:1,allLoaded:!1}}}},function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var n=a(4),i=(a.n(n),a(8)),s=a.n(i),o=a(5),r=(a.n(o),a(9)),l=a.n(r),c=a(0),d=a(7),u=a.n(d),m=a(10),p=a.n(m),h=a(2),f=a(11),v=a(6),g=(a.n(v),a(3));a.n(g);c.default.use(u.a),c.default.use(f.a),c.default.component(l.a.name,l.a),c.default.use(s.a),c.default.config.productionTip=!1,new c.default({el:"#app",router:h.a,template:"<App/>",components:{App:p.a}})},function(t,e){},function(t,e){},function(t,e){},function(t,e){},function(t,e){},function(t,e,a){a(21),a(20);var n=a(1)(a(13),a(28),"data-v-bb66213e",null);t.exports=n.exports},function(t,e,a){a(17);var n=a(1)(a(14),a(25),"data-v-13052efb",null);t.exports=n.exports},function(t,e,a){a(18);var n=a(1)(a(15),a(26),"data-v-17546a84",null);t.exports=n.exports},function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"container"},[a("g-header",{attrs:{hname:this.addParam.username},on:{post:t.report}}),t._v(" "),a("textarea",{directives:[{name:"model",rawName:"v-model",value:t.content,expression:"content"}],staticClass:"content",attrs:{name:"content",placeholder:"说点什么吧..."},domProps:{value:t.content},on:{input:function(e){e.target.composing||(t.content=e.target.value)}}}),t._v(" "),t._l(t.uploadList,function(e){return a("div",{staticClass:"demo-upload-list"},["finished"===e.status?[a("img",{attrs:{src:e.url}}),t._v(" "),a("div",{staticClass:"demo-upload-list-cover"},[a("Icon",{attrs:{type:"ios-eye-outline"},nativeOn:{click:function(a){t.handleView(e.name)}}}),t._v(" "),a("Icon",{attrs:{type:"ios-trash-outline"},nativeOn:{click:function(a){t.handleRemove(e)}}})],1)]:[e.showProgress?a("Progress",{attrs:{percent:e.percentage,"hide-info":""}}):t._e()]],2)}),t._v(" "),a("Upload",{ref:"upload",staticStyle:{float:"left",width:"58px"},attrs:{"show-upload-list":!1,"default-file-list":t.defaultList,"on-success":t.handleSuccess,format:["jpg","jpeg","png"],"max-size":2048,"on-format-error":t.handleFormatError,"on-exceeded-size":t.handleMaxSize,"before-upload":t.handleBeforeUpload,name:"image",multiple:"",type:"drag",action:"http://118.89.200.173/circle/upload_picture",data:t.param}},[a("div",{staticStyle:{width:"58px",height:"58px","line-height":"58px"}},[a("Icon",{attrs:{type:"camera",size:"20"}})],1)]),t._v(" "),a("Modal",{attrs:{title:"查看图片"},model:{value:t.visible,callback:function(e){t.visible=e},expression:"visible"}},[t.visible?a("img",{staticStyle:{width:"100%"},attrs:{src:"https://o5wwk8baw.qnssl.com/"+t.imgName+"/large"}}):t._e()])],2)},staticRenderFns:[]}},function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"page"},[a("header",[a("span",{on:{click:function(e){t.goAdd()}}},[t._v("发表")]),t._v(" "),a("span",{staticClass:"friendship"},[t._v("好友圈")]),t._v(" "),a("span")]),t._v(" "),a("mt-loadmore",{ref:"loadmore",attrs:{"top-method":t.loadTop}},[a("ul",{directives:[{name:"infinite-scroll",rawName:"v-infinite-scroll",value:t.loadBottom,expression:"loadBottom"}],staticClass:"page",attrs:{"infinite-scroll-disabled":"loading","infinite-scroll-distance":"10"}},t._l(t.friendCircleList,function(e){return a("li",{staticClass:"circle-item"},[a("div",{staticClass:"header"},[a("div",{staticClass:"headerImg"},[a("img",{attrs:{src:e.headerUrl,alt:"头像"}})]),t._v(" "),a("div",{staticClass:"intro"},[a("span",{staticStyle:{display:"inline-block"}},[t._v(t._s(e.UserName))]),t._v(" "),1==e.ishoney?a("i",{staticClass:"tags"}):t._e(),t._v(" "),a("span",{staticClass:"time"},[t._v(t._s(e.TimeStamp))])])]),t._v(" "),a("div",{staticClass:"content"},[t._v("\n                    "+t._s(e.Content)+"\n                ")]),t._v(" "),a("ul",{staticClass:"picList"},t._l(e.picList,function(t){return a("li",[a("img",{attrs:{src:t,alt:"图片"}})])}))])}))])],1)},staticRenderFns:[]}},function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"app"}},[a("router-view")],1)},staticRenderFns:[]}},function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"header"},[a("span",{staticClass:"left",on:{click:function(e){t.reback()}}},[t._v("取消")]),t._v(" "),a("span",[t._v("动态")]),t._v(" "),a("div",{staticClass:"left"},[a("Dropdown",{staticStyle:{"margin-left":"20px"},attrs:{trigger:"click",placement:"bottom-end"},on:{"on-click":t.Post}},[a("a",{attrs:{href:"javascript:void(0)",id:"postText"}},[t._v("\n                发表\n                "),a("Icon",{attrs:{type:"arrow-down-b"}})],1),t._v(" "),a("Dropdown-menu",{slot:"list"},[a("Dropdown-item",{attrs:{name:"0"}},[t._v("仅情侣可见")]),t._v(" "),a("Dropdown-item",{attrs:{name:"1"}},[t._v("开放")])],1)],1)],1)])},staticRenderFns:[]}},,,,function(t,e){}],[16]);
//# sourceMappingURL=app.f41a7e5936eb6d435850.js.map