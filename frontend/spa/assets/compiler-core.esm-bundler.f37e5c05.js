import{c as o,b as v,e as E,f as I,g as $}from"./QBtn.f4e1706d.js";import{c as n,h as r,g as b,r as y,b as p,s as R}from"./index.2bdbffaa.js";const m={dark:{type:Boolean,default:null}};function f(e,a){return n(()=>e.dark===null?a.dark.isActive:e.dark)}const D=["horizontal","vertical","cell","none"];var A=o({name:"QMarkupTable",props:{...m,dense:Boolean,flat:Boolean,bordered:Boolean,square:Boolean,wrapCells:Boolean,separator:{type:String,default:"horizontal",validator:e=>D.includes(e)}},setup(e,{slots:a}){const l=b(),u=f(e,l.proxy.$q),i=n(()=>`q-markup-table q-table__container q-table__card q-table--${e.separator}-separator`+(u.value===!0?" q-table--dark q-table__card--dark q-dark":"")+(e.dense===!0?" q-table--dense":"")+(e.flat===!0?" q-table--flat":"")+(e.bordered===!0?" q-table--bordered":"")+(e.square===!0?" q-table--square":"")+(e.wrapCells===!1?" q-table--no-wrap":""));return()=>r("div",{class:i.value},[r("table",{class:"q-table"},v(a.default))])}}),T=o({name:"QItemLabel",props:{overline:Boolean,caption:Boolean,header:Boolean,lines:[Number,String]},setup(e,{slots:a}){const l=n(()=>parseInt(e.lines,10)),u=n(()=>"q-item__label"+(e.overline===!0?" q-item__label--overline text-overline":"")+(e.caption===!0?" q-item__label--caption text-caption":"")+(e.header===!0?" q-item__label--header":"")+(l.value===1?" ellipsis":"")),i=n(()=>e.lines!==void 0&&l.value>1?{overflow:"hidden",display:"-webkit-box","-webkit-box-orient":"vertical","-webkit-line-clamp":l.value}:null);return()=>r("div",{style:i.value,class:u.value},v(a.default))}}),M=o({name:"QItemSection",props:{avatar:Boolean,thumbnail:Boolean,side:Boolean,top:Boolean,noWrap:Boolean},setup(e,{slots:a}){const l=n(()=>`q-item__section column q-item__section--${e.avatar===!0||e.side===!0||e.thumbnail===!0?"side":"main"}`+(e.top===!0?" q-item__section--top justify-start":" justify-center")+(e.avatar===!0?" q-item__section--avatar":"")+(e.thumbnail===!0?" q-item__section--thumbnail":"")+(e.noWrap===!0?" q-item__section--nowrap":""));return()=>r("div",{class:l.value},v(a.default))}}),N=o({name:"QItem",props:{...m,...E,tag:{type:String,default:"div"},active:{type:Boolean,default:null},clickable:Boolean,dense:Boolean,insetLevel:Number,tabindex:[String,Number],focused:Boolean,manualFocus:Boolean},emits:["click","keyup"],setup(e,{slots:a,emit:l}){const{proxy:{$q:u}}=b(),i=f(e,u),{hasLink:q,linkAttrs:_,linkClass:h,linkTag:B,navigateOnClick:w}=I(),c=y(null),d=y(null),k=n(()=>e.clickable===!0||q.value===!0||e.tag==="label"),s=n(()=>e.disable!==!0&&k.value===!0),x=n(()=>"q-item q-item-type row no-wrap"+(e.dense===!0?" q-item--dense":"")+(i.value===!0?" q-item--dark":"")+(q.value===!0&&e.active===null?h.value:e.active===!0?` q-item--active${e.activeClass!==void 0?` ${e.activeClass}`:""}`:"")+(e.disable===!0?" disabled":"")+(s.value===!0?" q-item--clickable q-link cursor-pointer "+(e.manualFocus===!0?"q-manual-focusable":"q-focusable q-hoverable")+(e.focused===!0?" q-manual-focusable--focused":""):"")),C=n(()=>{if(e.insetLevel===void 0)return null;const t=u.lang.rtl===!0?"Right":"Left";return{["padding"+t]:16+e.insetLevel*56+"px"}});function L(t){s.value===!0&&(d.value!==null&&(t.qKeyEvent!==!0&&document.activeElement===c.value?d.value.focus():document.activeElement===d.value&&c.value.focus()),w(t))}function Q(t){if(s.value===!0&&p(t,13)===!0){R(t),t.qKeyEvent=!0;const g=new MouseEvent("click",t);g.qKeyEvent=!0,c.value.dispatchEvent(g)}l("keyup",t)}function S(){const t=$(a.default,[]);return s.value===!0&&t.unshift(r("div",{class:"q-focus-helper",tabindex:-1,ref:d})),t}return()=>{const t={ref:c,class:x.value,style:C.value,role:"listitem",onClick:L,onKeyup:Q};return s.value===!0?(t.tabindex=e.tabindex||"0",Object.assign(t,_.value)):k.value===!0&&(t["aria-disabled"]="true"),r(B.value,t,S())}}}),P=o({name:"QList",props:{...m,bordered:Boolean,dense:Boolean,separator:Boolean,padding:Boolean,tag:{type:String,default:"div"}},setup(e,{slots:a}){const l=b(),u=f(e,l.proxy.$q),i=n(()=>"q-list"+(e.bordered===!0?" q-list--bordered":"")+(e.dense===!0?" q-list--dense":"")+(e.separator===!0?" q-list--separator":"")+(u.value===!0?" q-list--dark":"")+(e.padding===!0?" q-list--padding":""));return()=>r(e.tag,{class:i.value},v(a.default))}});new RegExp("\\b"+"do,if,for,let,new,try,var,case,else,with,await,break,catch,class,const,super,throw,while,yield,delete,export,import,return,switch,default,extends,finally,continue,debugger,function,arguments,typeof,void".split(",").join("\\b|\\b")+"\\b");export{A as Q,f as a,T as b,P as c,N as d,M as e,m as u};
