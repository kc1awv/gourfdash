import{Q as r}from"./QTable.b0282ea8.js";import{a as c,l}from"./axios.df14c455.js";import"./compiler-core.esm-bundler.f37e5c05.js";import{_ as m,A as d,r as p,B as u,Q as f,G as s,F as h}from"./index.2bdbffaa.js";import"./QBtn.f4e1706d.js";import"./scroll.4afce4cd.js";const _=[{name:"callsign",label:"Callsign",field:"callsign",sortable:!0},{name:"ip",label:"IP",field:"ip",sortable:!0},{name:"linkedmodule",label:"Linked Module",field:"linkedmodule",sortable:!0},{name:"protocol",label:"Protocol",field:"protocol",sortable:!0},{name:"connecttime",label:"Connect Time",field:"connecttime",sortable:!0},{name:"lastheardtime",label:"Last Heard Time",field:"lastheardtime",sortable:!0}],b=d({data(){return{name:"Links",link_columns:_,link_rows:[],pagination:p({rowsPerPage:0}),intervalId:null}},mounted(){this.fetchLinks(),this.intervalId=setInterval(()=>{this.fetchLinks()},2e4)},beforeRouteLeave(e,t,a){clearInterval(this.intervalId),a()},methods:{fetchLinks(){let e="/json/links";c.get(e).then(t=>{let a=t.data;a.forEach(function(o,n){o.id=n+1,o.connecttime=l(o.connecttime),o.lastheardtime=l(o.lastheardtime),a[n]=o}),this.link_rows=a}).catch(t=>{console.error(t)})}}}),g={class:"q-pa-sm"},k={class:"row"},v={class:"col-xs-12 col-md-12"},w={class:"q-pa-md"};function L(e,t,a,o,n,I){return u(),f("div",g,[s("div",k,[s("div",v,[s("div",w,[h(r,{rows:e.link_rows,columns:e.link_columns,"row-key":"name","virtual-scroll":"",pagination:e.pagination,"onUpdate:pagination":t[0]||(t[0]=i=>e.pagination=i),"rows-per-page-options":[0]},null,8,["rows","columns","pagination"])])])])])}var q=m(b,[["render",L]]);export{q as default};
