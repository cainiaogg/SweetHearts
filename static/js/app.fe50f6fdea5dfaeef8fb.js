webpackJsonp([1,0],[function(e,t,n){"use strict";function a(e){return e&&e.__esModule?e:{default:e}}var s=n(3),i=a(s),o=n(2),r=a(o);n(8);var c=n(11),l=a(c),u=n(12),f=a(u);i.default.use(r.default),i.default.use(l.default),new i.default({el:"#app",template:"<App/>",components:{App:f.default}})},,,,,function(e,t){},function(e,t){},function(e,t){},function(e,t){},,,,function(e,t,n){n(6);var a=n(1)(n(18),n(16),null,null);e.exports=a.exports},function(e,t,n){n(7);var a=n(1)(n(19),n(17),null,null);e.exports=a.exports},function(e,t,n){n(5);var a=n(1)(n(20),n(15),"data-v-1499b4dc",null);e.exports=a.exports},function(e,t){e.exports={render:function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("section",{staticClass:"top"},[n("img",{staticClass:"head",attrs:{src:e.friendnameImgHead}}),e._v(" "),n("span",{staticClass:"name"},[e._v(e._s(e.friendname))]),e._v(" "),n("span",{staticClass:"voice"},[n("i",{staticClass:"fa",class:e.vociceOpen?"fa-volume-up":"fa-volume-down",attrs:{"aria-hidden":"true"},on:{click:function(t){e.vociceOpen=!e.vociceOpen}}})])])},staticRenderFns:[]}},function(e,t){e.exports={render:function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"app"}},[n("Heads"),e._v(" "),n("ChatList")],1)},staticRenderFns:[]}},function(e,t){e.exports={render:function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("section",{staticClass:"chatlist",class:e.showSelBox>0?"chatlist-bottom-collapse":"chatlist-bottom"},[n("mt-loadmore",{ref:"loadmore",attrs:{"top-method":e.loadTop,"top-pull-text":"加载更多","top-drop-text":"释放加载"},on:{"top-status-change":e.handleTopChange}},[n("ul",[e._l(e.records,function(t){return[1==t.type?n("li",{staticClass:"chat-mine"},[n("div",{staticClass:"chat-user"},[n("img",{attrs:{src:e.usernameImgHead}})]),e._v(" "),n("div",{staticClass:"time"},[n("cite",[n("i",[e._v(e._s(t.time))]),e._v(e._s(t.name))])]),e._v(" "),n("div",{staticClass:"chat-text",domProps:{innerHTML:e._s(e.replaceFace(t.content))}})]):e._e(),e._v(" "),2==t.type?n("li",[n("div",{staticClass:"chat-user"},[n("img",{attrs:{src:e.friendnameImgHead}})]),e._v(" "),n("div",{staticClass:"time"},[n("cite",[e._v(e._s(t.name)),n("i",[e._v(e._s(t.time))])])]),e._v(" "),n("div",{staticClass:"chat-text",domProps:{innerHTML:e._s(e.replaceFace(t.content))}})]):e._e()]})],2)])],1),e._v(" "),n("section",{staticClass:"foot"},[n("mt-field",{staticClass:"con",attrs:{id:"txtContent",placeholder:"请输入消息"},model:{value:e.content,callback:function(t){e.content=t},expression:"content"}}),e._v(" "),n("span",{staticClass:"btn-face",on:{click:function(t){e.showSelBox=1==e.showSelBox?0:1}}},[n("i",{staticClass:"fa fa-smile-o",attrs:{"aria-hidden":"true"}})]),e._v(" "),n("span",{staticClass:"btn-plus",on:{click:function(t){e.showSelBox=2==e.showSelBox?0:2}}},[n("i",{staticClass:"fa",class:2==e.showSelBox?"fa-minus-circle":"fa-plus-circle",attrs:{"aria-hidden":"true"}})]),e._v(" "),n("span",{staticClass:"btn btn-send",on:{click:e.sendMsg}},[e._v("发送")]),e._v(" "),n("section",{staticClass:"selbox",class:e.showSelBox>0?"show":"hide"},[n("section",{directives:[{name:"show",rawName:"v-show",value:1==e.showSelBox,expression:"showSelBox==1"}],staticClass:"face-box"},[n("mt-swipe",{attrs:{auto:0,continuous:!1}},e._l(Math.ceil(e.EXPS.length/18),function(t){return n("mt-swipe-item",e._l(e.getEXP(t,18),function(t,a){return n("li",[n("img",{attrs:{src:"static/emotion/"+t.file,alt:"",data:t.code},on:{click:function(n){e.content+=t.code}}})])}))}))],1),e._v(" "),n("div",{directives:[{name:"show",rawName:"v-show",value:2==e.showSelBox,expression:"showSelBox==2"}]},[e._v(e._s(e.selOther))])])],1)])},staticRenderFns:[]}},function(e,t,n){"use strict";function a(e){return e&&e.__esModule?e:{default:e}}Object.defineProperty(t,"__esModule",{value:!0});var s=n(14),i=a(s),o=n(13),r=a(o);t.default={name:"app",components:{Heads:i.default,ChatList:r.default}}},function(e,t,n){"use strict";function a(e){return e&&e.__esModule?e:{default:e}}Object.defineProperty(t,"__esModule",{value:!0});var s=n(21),i=a(s),o=n(2),r=null,c=(new Date).getTime();t.default={name:"chatlist",created:function(){var e=/username=.*$/g,t=location.href.match(e)[0].split("&"),n=t[0].split("=")[1],a=t[1].split("=")[1];this.username=n,this.friendname=a,r=new WebSocket("ws://192.168.253.18:8080/chat?username="+n+"&friendname="+a),this.usernameImgHead="http://192.168.253.18:8080/image/head_image/"+n+".jpg",this.friendnameImgHead="http://192.168.253.18:8080/image/head_image/"+a+".jpg",this.loadTop()},data:function(){return{username:"",friendname:"",usernameImgHead:"",friendnameImgHead:"",showSelBox:0,selFace:"表情",selOther:"其他功能",content:"",topStatus:"",records:[],EXPS:[{file:"100.gif",code:"/::)",title:"微笑",reg:/\/::\)/g},{file:"101.gif",code:"/::~",title:"伤心",reg:/\/::~/g},{file:"102.gif",code:"/::B",title:"美女",reg:/\/::B/g},{file:"103.gif",code:"/::|",title:"发呆",reg:/\/::\|/g},{file:"104.gif",code:"/:8-)",title:"墨镜",reg:/\/:8-\)/g},{file:"105.gif",code:"/::<",title:"哭",reg:/\/::</g},{file:"106.gif",code:"/::$",title:"羞",reg:/\/::\$/g},{file:"107.gif",code:"/::X",title:"哑",reg:/\/::X/g},{file:"108.gif",code:"/::Z",title:"睡",reg:/\/::Z/g},{file:"109.gif",code:"/::'(",title:"哭",reg:/\/::'\(/g},{file:"110.gif",code:"/::-|",title:"囧",reg:/\/::-\|/g},{file:"111.gif",code:"/::@",title:"怒",reg:/\/::@/g},{file:"112.gif",code:"/::P",title:"调皮",reg:/\/::P/g},{file:"113.gif",code:"/::D",title:"笑",reg:/\/::D/g},{file:"114.gif",code:"/::O",title:"惊讶",reg:/\/::O/g},{file:"115.gif",code:"/::(",title:"难过",reg:/\/::\(/g},{file:"116.gif",code:"/::+",title:"酷",reg:/\/::\+/g},{file:"117.gif",code:"/:--b",title:"汗",reg:/\/:--b/g},{file:"118.gif",code:"/::Q",title:"抓狂",reg:/\/::Q/g},{file:"119.gif",code:"/::T",title:"吐",reg:/\/::T/g},{file:"120.gif",code:"/:,@P",title:"笑",reg:/\/:,@P/g},{file:"121.gif",code:"/:,@-D",title:"快乐",reg:/\/:,@-D/g},{file:"122.gif",code:"/::d",title:"奇",reg:/\/::d/g},{file:"123.gif",code:"/:,@o",title:"傲",reg:/\/:,@o/g},{file:"124.gif",code:"/::g",title:"饿",reg:/\/::g/g},{file:"125.gif",code:"/:|-)",title:"累",reg:/\/:\|-\)/g},{file:"126.gif",code:"/::!",title:"吓",reg:/\/::!/g}]}},methods:{getEXP:function(e,t){return this.EXPS.slice((e-1)*t,t*e)},sendMsg:function(){var e=this;return""==this.content?void(0,o.Toast)("请输入消息"):(r.onopen=i.default.send(r,this.content),this.records.push({type:1,time:i.default.formatDate.format(new Date,"yyyy-MM-dd hh:mm:ss"),name:this.username,content:this.content}),r.onmessage=function(t){console.log(e.records);var n=JSON.parse(t.data)[1];console.log(i.default.formatDate.myformat(n.TimeStamp)),e.records.push({type:2,time:i.default.formatDate.myformat(n.TimeStamp),name:e.friendname,content:n.Content})},this.content="",void this.scrollToBottom())},focusTxtContent:function(){document.querySelector("#txtContent input").focus()},scrollToBottom:function(){setTimeout(function(){var e=document.getElementsByClassName("chatlist")[0];e.scrollTop=e.scrollHeight},100)},replaceFace:function(e){for(var t=this.EXPS,n=0;n<t.length;n++)e=e.replace(t[n].reg,'<img src="static/emotion/'+t[n].file+'"  alt="" />');return e},handleTopChange:function(e){this.topStatus=e},loadTop:function(e){var t=this,n=this;console.log(c),setTimeout(function(){var a=document.getElementsByClassName("chatlist")[0],s=a.scrollHeight;t.$http.get("http://192.168.253.18:8080/chat/get_chat_record?username="+t.username+"&friendname="+t.friendname+"&timestamp="+c).then(function(e){console.log(e);var t=e.body;t&&t.forEach(function(e,t){var a={};a.content=e[1].Content,a.type=e[0],a.time=i.default.formatDate.myformat(e[1].TimeStamp),a.name=n.username,n.records.unshift(a),console.log(e),c=e[1].TimeStamp})}),setTimeout(function(){var e=a.scrollHeight;a.scrollTop=e-s},100),t.$refs.loadmore.onTopLoaded(e)},1500)}},mounted:function(){this.scrollToBottom(),this.focusTxtContent()}}},function(e,t){"use strict";Object.defineProperty(t,"__esModule",{value:!0}),t.default={name:"heads",created:function(){var e=/username=.*$/g,t=location.href.match(e)[0].split("&"),n=t[0].split("=")[1],a=t[1].split("=")[1];this.username=n,this.friendname=a,this.usernameImgHead="http://192.168.253.18:8080/image/head_image/"+n+".jpg",this.friendnameImgHead="http://192.168.253.18:8080/image/head_image/"+a+".jpg"},data:function(){return{vociceOpen:!0,username:"",friendname:"",usernameImgHead:"",friendnameImgHead:""}}}},function(e,t){"use strict";function n(e,t){for(var t=t-(e+"").length,n=0;n<t;n++)e="0"+e;return e}Object.defineProperty(t,"__esModule",{value:!0});var a=/([yMdhsm])(\1*)/g,s="yyyy-MM-dd";t.default={getQueryStringByName:function(e){var t=new RegExp("(^|&)"+e+"=([^&]*)(&|$)","i"),n=window.location.search.substr(1).match(t),a="";return null!=n&&(a=n[2]),t=null,n=null,null==a||""==a||"undefined"==a?"":a},formatDate:{myformat:function(e){return new Date(e).getFullYear()+"-"+(new Date(e).getMonth()+1)+"-"+new Date(e).getDate()+" "+new Date(e).getHours()+":"+(new Date(e).getMinutes()<10?"0"+new Date(e).getMinutes():new Date(e).getMinutes())+":"+(new Date(e).getSeconds()<10?"0"+new Date(e).getSeconds():new Date(e).getSeconds())},format:function(e,t){return t=t||s,t.replace(a,function(t){switch(t.charAt(0)){case"y":return n(e.getFullYear(),t.length);case"M":return n(e.getMonth()+1,t.length);case"d":return n(e.getDate(),t.length);case"w":return e.getDay()+1;case"h":return n(e.getHours(),t.length);case"m":return n(e.getMinutes(),t.length);case"s":return n(e.getSeconds(),t.length)}})},parse:function(e,t){var n=t.match(a),s=e.match(/(\d)+/g);if(n.length==s.length){for(var i=new Date(1970,0,1),o=0;o<n.length;o++){var r=parseInt(s[o]),c=n[o];switch(c.charAt(0)){case"y":i.setFullYear(r);break;case"M":i.setMonth(r-1);break;case"d":i.setDate(r);break;case"h":i.setHours(r);break;case"m":i.setMinutes(r);break;case"s":i.setSeconds(r)}}return i}return null}},send:function(e,t){e.send(t)}}},function(e,t){}]);
//# sourceMappingURL=app.fe50f6fdea5dfaeef8fb.js.map