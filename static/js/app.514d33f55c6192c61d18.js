webpackJsonp([0],[,,function(e,t,a){"use strict";var n=a(0),i=a(28),s=a(23),o=a.n(s),r=a(22),l=a.n(r);n.default.use(i.a),t.a=new i.a({routes:[{path:"/",name:"friendCircle",component:o.a},{path:"/add",name:"Add",component:l.a}]})},function(e,t){},function(e,t){},function(e,t){},function(e,t){},,,,function(e,t,a){a(19);var n=a(1)(a(12),a(26),null,null);e.exports=n.exports},,function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0}),t.default={name:"app"}},function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0}),t.default={name:"GHeader",props:{hname:{default:"admin"}},methods:{Post:function(){this.$emit("post"),console.log("被触发")},reback:function(){this.$router.push("/?username"+this.hname),console.log("返回了")}}}},function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var n=a(21),i=a.n(n);t.default={components:{GHeader:i.a},created:function(){var e=this;console.log("wer"),this.$http.get("http://172.18.76.42:8080/circle/get_id").then(function(t){console.log(t),e.addParam.id=t.data,e.param.id=t.data});var t=location.href.match(/username=.*$/g)[0].split("=")[1];this.addParam.username=this.param.username=t},data:function(){return{defaultList:[],imgName:"",visible:!1,uploadList:[],param:{username:"123",id:"1"},addParam:{username:"",id:"",ishoney:"0"},content:""}},methods:{report:function(){var e=this,t=this;this.$http.post("http://172.18.76.42:8080/circle/add_friend_circle",{username:this.addParam.username,id:this.addParam.id,content:this.content,ishoney:this.addParam.ishoney}).then(function(a){console.log(a),t.$router.push("/?username"+e.addParam.username)},function(e){console.log(e)})},handleView:function(e){this.imgName=e,this.visible=!0},handleRemove:function(e){var t=this.$refs.upload.fileList;this.$refs.upload.fileList.splice(t.indexOf(e),1)},handleSuccess:function(e,t){t.url=e[1],t.name=e[0],console.log(e),console.log(t),console.log(this.defaultList)},handleFormatError:function(e){this.$Notice.warning({title:"文件格式不正确",desc:"文件 "+e.name+" 格式不正确，请上传 jpg 或 png 格式的图片。"})},handleMaxSize:function(e){this.$Notice.warning({title:"超出文件大小限制",desc:"文件 "+e.name+" 太大，不能超过 2M。"})},handleBeforeUpload:function(){this.imgName=((new Date).getTime()+Math.random()).toString(),console.log(this.name);var e=this.uploadList.length<5;return e||this.$Notice.warning({title:"最多只能上传 5 张图片。"}),e}},mounted:function(){this.uploadList=this.$refs.upload.fileList}}},function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0}),t.default={created:function(){var e=/username=.*$/g;console.log(location.href.match(e));var t=location.href.match(e)[0].split("=")[1];this.username=t,this.initData(1)},methods:{goAdd:function(){this.$router.push("/add?username="+this.username)},initData:function(e){var t=this.username,a=e||this.page,n="http://172.18.76.42:8080/circle/get_friend_circle?username="+t+"&page="+a,i=this;1==a&&(this.friendCircleList=[]),this.$http.get(n).then(function(e){console.log(e.body),e.body.forEach(function(e,t){var a={};a.Content=e.Content,a.picList=e.ImagePath,a.UserName=e.UserName,a.TimeStamp=new Date(e.TimeStamp).getMonth()+1+"月"+new Date(e.TimeStamp).getDate()+"日"+new Date(e.TimeStamp).getHours()+":"+(new Date(e.TimeStamp).getMinutes()<10?"0"+new Date(e.TimeStamp).getMinutes():new Date(e.TimeStamp).getMinutes()),a.headerUrl=e.HeadImage?e.HeadImage:"https://o5wwk8baw.qnssl.com/a42bdcc1178e62b4694c830f028db5c0/avatar",a.ishoney=e.IsHoney,i.friendCircleList.push(a)}),i.loading=!1})},loadTop:function(){this.$refs.loadmore.onTopLoaded(),this.initData(1)},loadBottom:function(){this.loading=!0;var e=this.page++;this.initData(e)}},data:function(){return{friendCircleList:[],username:"admin",page:1,allLoaded:!1}}}},function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var n=a(4),i=(a.n(n),a(8)),s=a.n(i),o=a(5),r=(a.n(o),a(9)),l=a.n(r),c=a(0),d=a(7),u=a.n(d),p=a(10),m=a.n(p),f=a(2),h=a(11),v=a(6),g=(a.n(v),a(3));a.n(g);c.default.use(u.a),c.default.use(h.a),c.default.component(l.a.name,l.a),c.default.use(s.a),c.default.config.productionTip=!1,new c.default({el:"#app",router:f.a,template:"<App/>",components:{App:m.a}})},function(e,t){},function(e,t){},function(e,t){},function(e,t){},function(e,t,a){a(20);var n=a(1)(a(13),a(27),"data-v-bb66213e",null);e.exports=n.exports},function(e,t,a){a(17);var n=a(1)(a(14),a(24),"data-v-13052efb",null);e.exports=n.exports},function(e,t,a){a(18);var n=a(1)(a(15),a(25),"data-v-17546a84",null);e.exports=n.exports},function(e,t){e.exports={render:function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"container"},[a("g-header",{attrs:{hname:this.addParam.username},on:{post:e.report}}),e._v(" "),a("textarea",{directives:[{name:"model",rawName:"v-model",value:e.content,expression:"content"}],staticClass:"content",attrs:{name:"content",placeholder:"说点什么吧..."},domProps:{value:e.content},on:{input:function(t){t.target.composing||(e.content=t.target.value)}}}),e._v(" "),e._l(e.uploadList,function(t){return a("div",{staticClass:"demo-upload-list"},["finished"===t.status?[a("img",{attrs:{src:t.url}}),e._v(" "),a("div",{staticClass:"demo-upload-list-cover"},[a("Icon",{attrs:{type:"ios-eye-outline"},nativeOn:{click:function(a){e.handleView(t.name)}}}),e._v(" "),a("Icon",{attrs:{type:"ios-trash-outline"},nativeOn:{click:function(a){e.handleRemove(t)}}})],1)]:[t.showProgress?a("Progress",{attrs:{percent:t.percentage,"hide-info":""}}):e._e()]],2)}),e._v(" "),a("Upload",{ref:"upload",staticStyle:{float:"left",width:"58px"},attrs:{"show-upload-list":!1,"default-file-list":e.defaultList,"on-success":e.handleSuccess,format:["jpg","jpeg","png"],"max-size":2048,"on-format-error":e.handleFormatError,"on-exceeded-size":e.handleMaxSize,"before-upload":e.handleBeforeUpload,name:"image",multiple:"",type:"drag",action:"http://172.18.76.42:8080/circle/upload_picture",data:e.param}},[a("div",{staticStyle:{width:"58px",height:"58px","line-height":"58px"}},[a("Icon",{attrs:{type:"camera",size:"20"}})],1)]),e._v(" "),a("Modal",{attrs:{title:"查看图片"},model:{value:e.visible,callback:function(t){e.visible=t},expression:"visible"}},[e.visible?a("img",{staticStyle:{width:"100%"},attrs:{src:"https://o5wwk8baw.qnssl.com/"+e.imgName+"/large"}}):e._e()])],2)},staticRenderFns:[]}},function(e,t){e.exports={render:function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"page"},[a("header",[a("span",{on:{click:function(t){e.goAdd()}}},[e._v("发表")]),e._v(" "),a("span",{staticClass:"friendship"},[e._v("好友圈")]),e._v(" "),a("span")]),e._v(" "),a("mt-loadmore",{ref:"loadmore",attrs:{"top-method":e.loadTop}},[a("ul",{directives:[{name:"infinite-scroll",rawName:"v-infinite-scroll",value:e.loadBottom,expression:"loadBottom"}],staticClass:"page",attrs:{"infinite-scroll-disabled":"loading","infinite-scroll-distance":"10"}},e._l(e.friendCircleList,function(t){return a("li",{staticClass:"circle-item"},[a("div",{staticClass:"header"},[a("div",{staticClass:"headerImg"},[a("img",{attrs:{src:t.headerUrl,alt:"头像"}})]),e._v(" "),a("div",{staticClass:"intro"},[a("span",{staticStyle:{display:"inline-block"}},[e._v(e._s(t.UserName))]),e._v(" "),1==t.ishoney?a("i",{staticClass:"tags"}):e._e(),e._v(" "),a("span",{staticClass:"time"},[e._v(e._s(t.TimeStamp))])])]),e._v(" "),a("div",{staticClass:"content"},[e._v("\n                    "+e._s(t.Content)+"\n                ")]),e._v(" "),a("ul",{staticClass:"picList"},e._l(t.picList,function(e){return a("li",[a("img",{attrs:{src:e,alt:"图片"}})])}))])}))])],1)},staticRenderFns:[]}},function(e,t){e.exports={render:function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{attrs:{id:"app"}},[a("router-view")],1)},staticRenderFns:[]}},function(e,t){e.exports={render:function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"header"},[a("span",{staticClass:"left",on:{click:function(t){e.reback()}}},[e._v("取消")]),e._v(" "),a("span",[e._v("动态")]),e._v(" "),a("span",{on:{click:e.Post}},[e._v("发表")])])},staticRenderFns:[]}},,,,function(e,t){}],[16]);
//# sourceMappingURL=app.514d33f55c6192c61d18.js.map