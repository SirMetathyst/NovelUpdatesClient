package NovelUpdatesClient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc
// Capture Date: DD/MM/YYYY 07/04/22 @ 8:19 PM BST
var TestSuccessRequester = DummyRequester(`
<!DOCTYPE HTML>
<html class="" lang="en-US">
<head>
	<meta charset="UTF-8">
	<title>Series Finder - Novel Updates</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
	<meta name="SKYPE_TOOLBAR" content="SKYPE_TOOLBAR_PARSER_COMPATIBLE" />
    <meta name="referrer" content="unsafe-url">

	    <link rel="apple-touch-icon" href="/appicon.png"/>
   
<!-- UniConsent Cookies Consent Notice start -->
<script src="https://cmp.uniconsent.com/v2/stub.min.js"></script>
<script async src='https://cmp.uniconsent.com/v2/69a34e0934/cmp.js'></script>
<!-- UniConsent Cookies Consent Notice end -->
	
	<link rel='dns-prefetch' href='//www.novelupdates.com' />
<link rel='dns-prefetch' href='//www.google.com' />
<link rel='dns-prefetch' href='//cdn.novelupdates.com' />
<link rel='dns-prefetch' href='//ajax.googleapis.com' />
<link rel='dns-prefetch' href='//fonts.googleapis.com' />
<link rel='dns-prefetch' href='//s.w.org' />
<link rel="alternate" type="application/rss+xml" title="Novel Updates &raquo; Feed" href="https://www.novelupdates.com/feed/" />
<link rel="alternate" type="application/rss+xml" title="Novel Updates &raquo; Comments Feed" href="https://www.novelupdates.com/comments/feed/" />
		<script type="text/javascript">
			window._wpemojiSettings = {"baseUrl":"https:\/\/s.w.org\/images\/core\/emoji\/12.0.0-1\/72x72\/","ext":".png","svgUrl":"https:\/\/s.w.org\/images\/core\/emoji\/12.0.0-1\/svg\/","svgExt":".svg","source":{"concatemoji":"https:\/\/www.novelupdates.com\/wp-includes\/js\/wp-emoji-release.min.js?ver=5.2.7"}};
			!function(a,b,c){function d(a,b){var c=String.fromCharCode;l.clearRect(0,0,k.width,k.height),l.fillText(c.apply(this,a),0,0);var d=k.toDataURL();l.clearRect(0,0,k.width,k.height),l.fillText(c.apply(this,b),0,0);var e=k.toDataURL();return d===e}function e(a){var b;if(!l||!l.fillText)return!1;switch(l.textBaseline="top",l.font="600 32px Arial",a){case"flag":return!(b=d([55356,56826,55356,56819],[55356,56826,8203,55356,56819]))&&(b=d([55356,57332,56128,56423,56128,56418,56128,56421,56128,56430,56128,56423,56128,56447],[55356,57332,8203,56128,56423,8203,56128,56418,8203,56128,56421,8203,56128,56430,8203,56128,56423,8203,56128,56447]),!b);case"emoji":return b=d([55357,56424,55356,57342,8205,55358,56605,8205,55357,56424,55356,57340],[55357,56424,55356,57342,8203,55358,56605,8203,55357,56424,55356,57340]),!b}return!1}function f(a){var c=b.createElement("script");c.src=a,c.defer=c.type="text/javascript",b.getElementsByTagName("head")[0].appendChild(c)}var g,h,i,j,k=b.createElement("canvas"),l=k.getContext&&k.getContext("2d");for(j=Array("flag","emoji"),c.supports={everything:!0,everythingExceptFlag:!0},i=0;i<j.length;i++)c.supports[j[i]]=e(j[i]),c.supports.everything=c.supports.everything&&c.supports[j[i]],"flag"!==j[i]&&(c.supports.everythingExceptFlag=c.supports.everythingExceptFlag&&c.supports[j[i]]);c.supports.everythingExceptFlag=c.supports.everythingExceptFlag&&!c.supports.flag,c.DOMReady=!1,c.readyCallback=function(){c.DOMReady=!0},c.supports.everything||(h=function(){c.readyCallback()},b.addEventListener?(b.addEventListener("DOMContentLoaded",h,!1),a.addEventListener("load",h,!1)):(a.attachEvent("onload",h),b.attachEvent("onreadystatechange",function(){"complete"===b.readyState&&c.readyCallback()})),g=c.source||{},g.concatemoji?f(g.concatemoji):g.wpemoji&&g.twemoji&&(f(g.twemoji),f(g.wpemoji)))}(window,document,window._wpemojiSettings);
		</script>
		<style type="text/css">
img.wp-smiley,
img.emoji {
	display: inline !important;
	border: none !important;
	box-shadow: none !important;
	height: 1em !important;
	width: 1em !important;
	margin: 0 .07em !important;
	vertical-align: -0.1em !important;
	background: none !important;
	padding: 0 !important;
}
</style>
	<link rel='stylesheet' id='dashicons-css'  href='https://www.novelupdates.com/wp-includes/css/dashicons.min.css?ver=5.2.7' type='text/css' media='all' />
<link rel='stylesheet' id='theme-my-login-css'  href='https://www.novelupdates.com/wp-content/plugins/theme-my-login/theme-my-login.css?ver=6.4.10' type='text/css' media='all' />
<link rel='stylesheet' id='wp-block-library-css'  href='https://www.novelupdates.com/wp-includes/css/dist/block-library/style.min.css?ver=5.2.7' type='text/css' media='all' />
<link rel='stylesheet' id='bbspoiler-css'  href='https://www.novelupdates.com/wp-content/plugins/bbspoiler/inc/bbspoiler.css?ver=5.2.7' type='text/css' media='all' />
<link rel='stylesheet' id='cstmsrch_stylesheet-css'  href='https://www.novelupdates.com/wp-content/plugins/custom-search-plugin/css/style.css?ver=5.2.7' type='text/css' media='all' />
<link rel='stylesheet' id='yasrcss-css'  href='https://www.novelupdates.com/wp-content/plugins/yet-another-stars-rating/css/yasr.css' type='text/css' media='all' />
<link rel='stylesheet' id='yasrcsslightscheme-css'  href='https://www.novelupdates.com/wp-content/plugins/yet-another-stars-rating/css/yasr-table-light.css' type='text/css' media='all' />
<link rel='stylesheet' id='breadcrumbs-nu-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/breadcrumbs.css?ver=1.2.5' type='text/css' media='' />
<link rel='stylesheet' id='global-extras-css'  href='//cdn.novelupdates.com/js/global_extras.css?ver=1.0.6' type='text/css' media='' />
<link rel='stylesheet' id='ajax-search-posts-css-css'  href='//cdn.novelupdates.com/js/ajax_search.post.css?ver=1.3.6' type='text/css' media='' />
<link rel='stylesheet' id='alert-nu-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/alert_nu.css?ver=1.4.1' type='text/css' media='' />
<link rel='stylesheet' id='select2-css-css'  href='//cdn.novelupdates.com/js/select2.min.css?ver=1.0.1' type='text/css' media='all' />
<link rel='stylesheet' id='select2-bootstrap-css-css'  href='//cdn.novelupdates.com/js/select2-bootstrap.min.css?ver=1.0.1' type='text/css' media='all' />
<link rel='stylesheet' id='pagination-nu-css'  href='//cdn.novelupdates.com/js/pg/pagination.css?ver=1.1.1' type='text/css' media='' />
<link rel='stylesheet' id='us-font-1-css'  href='https://fonts.googleapis.com/css?family=Noto+Sans%3A300%2C400%2C700&#038;subset=latin&#038;ver=5.2.7' type='text/css' media='all' />
<link rel='stylesheet' id='us-font-2-css'  href='https://fonts.googleapis.com/css?family=Open+Sans%3A400%2C400italic%2C700%2C700italic%2C300&#038;subset=latin&#038;ver=5.2.7' type='text/css' media='all' />
<link rel='stylesheet' id='us-motioncss-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates/css/motioncss.css?ver=4.4.8' type='text/css' media='all' />
<link rel='stylesheet' id='us-motioncss-responsive-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates/css/motioncss-responsive.css?ver=4.4.8' type='text/css' media='all' />
<link rel='stylesheet' id='us-font-awesome-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates/css/font-awesome.css?ver=4.7.1' type='text/css' media='all' />
<link rel='stylesheet' id='us-style-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates/css/style.css?ver=4.4.8' type='text/css' media='all' />
<link rel='stylesheet' id='us-responsive-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates/css/responsive.css?ver=4.4.8' type='text/css' media='all' />
<link rel='stylesheet' id='ndupdates-style-css'  href='//www.novelupdates.com/wp-content/themes/ndupdates-child/style.css?ver=4.4.8' type='text/css' media='all' />
<script type='text/javascript' src='//ajax.googleapis.com/ajax/libs/jquery/1.8.3/jquery.min.js'></script>
<script type='text/javascript'>
/* <![CDATA[ */
var title = {"unfolded":"Expand","folded":"Collapse"};
/* ]]> */
</script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/plugins/bbspoiler/inc/bbspoiler.js?ver=5.2.7'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-includes/js/jquery/ui/core.min.js?ver=1.11.4'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-includes/js/jquery/ui/widget.min.js?ver=1.11.4'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-includes/js/jquery/ui/mouse.min.js?ver=1.11.4'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-includes/js/jquery/ui/sortable.min.js?ver=1.11.4'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/plugins/custom-search-plugin/js/script.js?ver=5.2.7'></script>
<script type='text/javascript' src='https://www.google.com/recaptcha/api.js?hl=en-US&#038;ver=5.2.7'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/plugins/theme-my-login/modules/themed-profiles/themed-profiles.js?ver=5.2.7'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/nu_extras_js.js?ver=1.0.4'></script>
<script type='text/javascript' src='//cdn.novelupdates.com/js/ajax_search_post.js?ver=1.3.6'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/alert_nu.js?ver=1.4.1'></script>
<script type='text/javascript' src='//cdn.novelupdates.com/js/select2.min.js?ver=1.0.1'></script>
<script type='text/javascript' src='//cdn.novelupdates.com/js/popoverlay/jquery.popupoverlay.js?ver=1.0.1'></script>
<script type='text/javascript' src='//cdn.novelupdates.com/js/pg/pagination.js?ver=1.0.2'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/seriesfinder.js?ver=1.0.8'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/chosen.jquery.min.js?ver=1.0.2'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/jquery.tablesorter.min.js?ver=1.0.1'></script>
<link rel='https://api.w.org/' href='https://www.novelupdates.com/wp-json/' />
<link rel="EditURI" type="application/rsd+xml" title="RSD" href="https://www.novelupdates.com/xmlrpc.php?rsd" />
<link rel="wlwmanifest" type="application/wlwmanifest+xml" href="https://www.novelupdates.com/wp-includes/wlwmanifest.xml" /> 
<meta name="generator" content="WordPress 5.2.7" />
<link rel="canonical" href="https://www.novelupdates.com/series-finder/" />
<link rel='shortlink' href='https://www.novelupdates.com/?p=3370' />
<link rel="alternate" type="application/json+oembed" href="https://www.novelupdates.com/wp-json/oembed/1.0/embed?url=https%3A%2F%2Fwww.novelupdates.com%2Fseries-finder%2F" />
<link rel="alternate" type="text/xml+oembed" href="https://www.novelupdates.com/wp-json/oembed/1.0/embed?url=https%3A%2F%2Fwww.novelupdates.com%2Fseries-finder%2F&#038;format=xml" />
 
   
    

	<style id="us_fonts_inline">
body {
		font-family: 'Open Sans';
		font-size: 15px;
	line-height: 26px;
	font-weight: 400;
	}
.page-template-page-blank-php .l-main {
	font-size: 15px;
	}
	
.l-header .menu-item-language,
.l-header .w-nav-item {
		font-family: 'Open Sans';
		font-weight: 300;
	}
.type_desktop .menu-item-language > a,
.l-header .type_desktop .w-nav-anchor.level_1,
.type_desktop [class*="columns"] .menu-item-has-children .w-nav-anchor.level_2 {
	font-size: 16px;
	}
.type_desktop .submenu-languages .menu-item-language > a,
.l-header .type_desktop .w-nav-anchor.level_2,
.l-header .type_desktop .w-nav-anchor.level_3,
.l-header .type_desktop .w-nav-anchor.level_4 {
	font-size: 15px;
	}
.type_mobile .menu-item-language > a,
.l-header .type_mobile .w-nav-anchor.level_1 {
	font-size: 16px;
	}
.l-header .type_mobile .w-nav-anchor.level_2,
.l-header .type_mobile .w-nav-anchor.level_3,
.l-header .type_mobile .w-nav-anchor.level_4 {
	font-size: 15px;
	}

h1, h2, h3, h4, h5, h6,
.w-counter-number,
.w-logo-title,
.w-pricing-item-title,
.w-pricing-item-price,
.w-shortblog-entry-date-day,
.ult_price_figure,
.ult_countdown-amount,
.ultb3-box .ultb3-title,
.stats-block .stats-desc .stats-number {
		font-family: 'Noto Sans';
		font-weight: 300;
	}
h1 {
	font-size: 38px;
	}
h2 {
	font-size: 32px;
	}
h3 {
	font-size: 26px;
	}
h4,
.widgettitle,
.comment-reply-title,
.ultb3-box .ultb3-title,
.flip-box-wrap .flip-box .ifb-face h3,
.aio-icon-box .aio-icon-header h3.aio-icon-title {
	font-size: 20px;
	}
h5,
.w-portfolio-item-title {
	font-size: 20px;
	}
h6 {
	font-size: 18px;
	}
@media only screen and (max-width: 767px) {
body {
	font-size: 13px;
	line-height: 23px;
	}
h1 {
	font-size: 30px;
	}
h2 {
	font-size: 26px;
	}
h3 {
	font-size: 22px;
	}
h4,
.widgettitle,
.comment-reply-title,
.ultb3-box .ultb3-title,
.flip-box-wrap .flip-box .ifb-face h3,
.aio-icon-box .aio-icon-header h3.aio-icon-title {
	font-size: 20px;
	}
h5,
.w-portfolio-item-title {
	font-size: 18px;
	}
h6 {
	font-size: 16px;
	}
}

.l-body,
.headerpos_fixed .l-header {
	min-width: 1240px;
	}
.l-canvas.type_boxed,
.l-canvas.type_boxed .l-subheader,
.l-canvas.type_boxed ~ .l-footer .l-subfooter {
	max-width: 1240px;
	}
.l-subheader-h,
.l-submain-h,
.l-subfooter-h {
	max-width: 1140px;
	}
.l-sidebar {
	width: 25%;
	}
.l-content {
	width: 70%;
	}
@media only screen and (max-width: 767px) {
.g-cols.type_boxed,
.g-cols.type_boxed > div {
	display: block;
	}
.g-cols > div {
	width: 100% !important;
	margin-left: 0 !important;
	margin-bottom: 30px;
	}
.g-cols.offset_none > div,
.g-cols > div:last-child {
	margin-bottom: 0 !important;
	}
}

@media only screen and (min-width: 900px) {
.w-logo-img {
	height: 30px;
	}
.w-logo.with_transparent .w-logo-img > img.for_default {
	margin-bottom: -30px;
	}
.l-header.sticky .w-logo-img {
	height: 30px;
	}
.l-header.sticky .w-logo.with_transparent .w-logo-img > img.for_default {
	margin-bottom: -30px;
	}
}
@media only screen and (min-width: 600px) and (max-width: 899px) {
.w-logo-img {
	height: 30px;
	}
.w-logo.with_transparent .w-logo-img > img.for_default {
	margin-bottom: -30px;
	}
}
@media only screen and (max-width: 599px) {
.w-logo-img {
	height: 30px;
	}
.w-logo.with_transparent .w-logo-img > img.for_default {
	margin-bottom: -30px;
	}
}
</style>
<style id="us_colors_inline">
/*************************** BODY ***************************/

/* Body Background Color */
.l-body {
	background-color: #2c3e50;
	}

/*************************** HEADER ***************************/

/* Header Background Color */
.l-subheader.at_middle,
.l-subheader.at_middle .w-lang-list,
.l-subheader.at_middle .type_mobile .w-nav-list.level_1 {
	background-color: #2c3e50;
	}

/* Header Text Color */
.l-subheader.at_middle,
.transparent .l-subheader.at_middle .type_mobile .w-nav-list.level_1 {
	color: #edf0f2;
	}

/* Header Text Hover Color */
.no-touch .w-logo-link:hover,
.no-touch .l-subheader.at_middle .w-contacts-item-value a:hover,
.no-touch .l-subheader.at_middle .w-lang-item:hover,
.no-touch .transparent .l-subheader.at_middle .w-lang.active .w-lang-item:hover,
.no-touch .l-subheader.at_middle .w-cart:hover .w-cart-link,
.no-touch .l-subheader.at_middle .w-search-show:hover,
.l-subheader.at_middle .w-cart-quantity {
	color: #fc4349;
	}

/* Extended Header Background Color */
.l-subheader.at_top,
.l-subheader.at_top .w-lang-list,
.l-subheader.at_bottom,
.l-subheader.at_bottom .type_mobile .w-nav-list.level_1 {
	background-color: #384b5f;
	}

/* Extended Header Text Color */
.l-subheader.at_top,
.l-subheader.at_bottom,
.transparent .l-subheader.at_bottom .type_mobile .w-nav-list.level_1,
.w-lang.active .w-lang-item {
	color: #d3d8db;
	}

/* Extended Header Text Hover Color */
.no-touch .l-subheader.at_top .w-contacts-item-value a:hover,
.no-touch .l-subheader.at_top .w-lang-item:hover,
.no-touch .transparent .l-subheader.at_top .w-lang.active .w-lang-item:hover,
.no-touch .l-subheader.at_bottom .w-cart:hover .w-cart-link,
.no-touch .l-subheader.at_bottom .w-search-show:hover,
.l-subheader.at_bottom .w-cart-quantity {
	color: #fff;
	}
	
/* Transparent Header Text Color */
.l-header.transparent .l-subheader {
	color: #fff;
	}
	
/* Transparent Header Text Hover Color */
.no-touch .l-header.transparent .type_desktop .menu-item-language > a:hover,
.no-touch .l-header.transparent .type_desktop .menu-item-language:hover > a,
.no-touch .l-header.transparent .w-logo-link:hover,
.no-touch .l-header.transparent .l-subheader .w-contacts-item-value a:hover,
.no-touch .l-header.transparent .l-subheader .w-lang-item:hover,
.no-touch .l-header.transparent .l-subheader .w-search-show:hover,
.no-touch .l-header.transparent .l-subheader .w-cart-link:hover,
.l-header.transparent .l-subheader .w-cart-quantity,
.no-touch .l-header.transparent .type_desktop .w-nav-item.level_1:hover .w-nav-anchor.level_1 {
	color: #fc4349;
	}
.l-header.transparent .w-nav-title:after {
	background-color: #fc4349;
	}
	
/* Search Screen Background Color */
.l-subheader .w-search-form:before {
	background-color: #6dbcdb;
	}

/* Search Screen Text Color */
.l-subheader .w-search-form {
	color: #fff;
	}

/*************************** MAIN MENU ***************************/

/* Menu Hover Background Color */
.no-touch .l-header .menu-item-language > a:hover,
.no-touch .type_desktop .menu-item-language:hover > a,
.no-touch .l-header .w-nav-item.level_1:hover .w-nav-anchor.level_1 {
	background-color: #2c3e50;
	}

/* Menu Hover Text Color */
.no-touch .l-header .menu-item-language > a:hover,
.no-touch .type_desktop .menu-item-language:hover > a,
.no-touch .l-header .w-nav-item.level_1:hover .w-nav-anchor.level_1 {
	color: #fc4349;
	}
.w-nav-title:after {
	background-color: #fc4349;
	}

/* Menu Active Background Color */
.l-header .w-nav-item.level_1.current-menu-item .w-nav-anchor.level_1,
.l-header .w-nav-item.level_1.current-menu-ancestor .w-nav-anchor.level_1 {
	background-color: #2c3e50;
	}

/* Menu Active Text Color */
.l-header .w-nav-item.level_1.current-menu-item .w-nav-anchor.level_1,
.l-header .w-nav-item.level_1.current-menu-ancestor .w-nav-anchor.level_1 {
	color: #6dbcdb;
	}
	
/* Transparent Menu Active Text Color */
.l-header.transparent .type_desktop .w-nav-item.level_1.active .w-nav-anchor.level_1,
.l-header.transparent .type_desktop .w-nav-item.level_1.current-menu-item .w-nav-anchor.level_1,
.l-header.transparent .type_desktop .w-nav-item.level_1.current-menu-ancestor .w-nav-anchor.level_1 {
	color: #6dbcdb;
	}

/* Dropdown Background Color */
.type_desktop .submenu-languages,
.l-header .w-nav-list.level_2,
.l-header .w-nav-list.level_3,
.l-header .w-nav-list.level_4 {
	background-color: #2c3e50;
	}

/* Dropdown Text Color */
.type_desktop .submenu-languages,
.l-header .w-nav-anchor.level_2,
.l-header .w-nav-anchor.level_3,
.l-header .w-nav-anchor.level_4,
.type_desktop [class*="columns"] .w-nav-item.menu-item-has-children.current-menu-item .w-nav-anchor.level_2,
.type_desktop [class*="columns"] .w-nav-item.menu-item-has-children.current-menu-ancestor .w-nav-anchor.level_2,
.no-touch .type_desktop [class*="columns"] .w-nav-item.menu-item-has-children:hover .w-nav-anchor.level_2 {
	color: #edf0f2;
	}

/* Dropdown Hover Background Color */
.no-touch .type_desktop .submenu-languages .menu-item-language:hover > a,
.no-touch .l-header .w-nav-item.level_2:hover .w-nav-anchor.level_2,
.no-touch .l-header .w-nav-item.level_3:hover .w-nav-anchor.level_3,
.no-touch .l-header .w-nav-item.level_4:hover .w-nav-anchor.level_4 {
	background-color: #2c3e50;
	}

/* Dropdown Hover Text Color */
.no-touch .type_desktop .submenu-languages .menu-item-language:hover > a,
.no-touch .l-header .w-nav-item.level_2:hover .w-nav-anchor.level_2,
.no-touch .l-header .w-nav-item.level_3:hover .w-nav-anchor.level_3,
.no-touch .l-header .w-nav-item.level_4:hover .w-nav-anchor.level_4 {
	color: #fc4349;
	}

/* Dropdown Active Background Color */
.l-header .w-nav-item.level_2.current-menu-item .w-nav-anchor.level_2,
.l-header .w-nav-item.level_2.current-menu-ancestor .w-nav-anchor.level_2,
.l-header .w-nav-item.level_3.current-menu-item .w-nav-anchor.level_3,
.l-header .w-nav-item.level_3.current-menu-ancestor .w-nav-anchor.level_3,
.l-header .w-nav-item.level_4.current-menu-item .w-nav-anchor.level_4,
.l-header .w-nav-item.level_4.current-menu-ancestor .w-nav-anchor.level_4 {
	background-color: #2c3e50;
	}

/* Dropdown Active Text Color */
.l-header .w-nav-item.level_2.current-menu-item .w-nav-anchor.level_2,
.l-header .w-nav-item.level_2.current-menu-ancestor .w-nav-anchor.level_2,
.l-header .w-nav-item.level_3.current-menu-item .w-nav-anchor.level_3,
.l-header .w-nav-item.level_3.current-menu-ancestor .w-nav-anchor.level_3,
.l-header .w-nav-item.level_4.current-menu-item .w-nav-anchor.level_4,
.l-header .w-nav-item.level_4.current-menu-ancestor .w-nav-anchor.level_4 {
	color: #6dbcdb;
	}

/* Menu Button Background Color */
.btn.w-nav-item .w-nav-anchor.level_1,
.headerpos_fixed .transparent .btn.w-nav-item .w-nav-anchor.level_1 {
	background-color: #6dbcdb !important;
	}

/* Menu Button Text Color */
.btn.w-nav-item .w-nav-anchor.level_1 {
	color: #fff !important;
	}

/* Menu Button Hover Background Color */
.no-touch .btn.w-nav-item .w-nav-anchor.level_1:before {
	background-color: #fc4349 !important;
	}

/* Menu Button Hover Text Color */
.no-touch .btn.w-nav-item .w-nav-anchor.level_1:hover {
	color: #fff !important;
	}

/*************************** MAIN CONTENT ***************************/

/* Background Color */
.l-canvas,
.w-blog.type_masonry .w-blog-entry-preview:before,
.w-cart-dropdown,
.w-filters-item.active,
.no-touch .w-filters-item.active:hover,
.w-portfolio-item-anchor,
.w-tabs-item.active,
.no-touch .w-tabs-item.active:hover,
.w-timeline-item,
.w-timeline-section-title,
.w-timeline.type_vertical .w-timeline-section-content,
#lang_sel ul ul,
#lang_sel_click ul ul,
#lang_sel_footer,
.woocommerce div.product .woocommerce-tabs .tabs li.active,
.no-touch .woocommerce div.product .woocommerce-tabs .tabs li.active:hover,
.woocommerce .stars span:after,
.woocommerce .stars span a:after,
#bbp-user-navigation li.current,
.gform_wrapper .chosen-container .chosen-drop,
.gform_wrapper .chosen-container-multi .chosen-choices li.search-choice {
	background-color: #fff;
	}
.g-btn.color_contrast,
.no-touch .g-btn.color_contrast:hover,
.no-touch .g-btn.color_contrast.outlined:hover,
.w-icon.color_border.type_solid .w-icon-link,
.w-iconbox.color_contrast.type_solid .w-iconbox-icon {
	color: #fff;
	}

/* Alternate Background Color */
input[type="text"],
input[type="password"],
input[type="email"],
input[type="url"],
input[type="tel"],
input[type="number"],
input[type="date"],
input[type="search"],
textarea,
select,
.w-actionbox.color_alternate,
.w-author,
.w-blog.imgpos_atleft .w-blog-entry-preview-icon,
.w-filters,
.w-icon.color_text.type_solid .w-icon-link,
.w-icon.color_fade.type_solid .w-icon-link,
.w-pricing-item-title,
.w-pricing-item-price,
.w-progbar-bar,
.w-tabs-list,
.no-touch .widget_nav_menu .menu-item a:hover,
.no-touch #lang_sel a:hover,
.no-touch #lang_sel_click a:hover,
.woocommerce .quantity .plus,
.woocommerce .quantity .minus,
.woocommerce div.product .woocommerce-tabs .tabs,
.woocommerce table.shop_table .actions .coupon .input-text,
.woocommerce #payment .payment_box,
#bbp-user-navigation,
#subscription-toggle,
#favorite-toggle,
.bbp-row-actions #favorite-toggle a,
.bbp-row-actions #subscription-toggle a,
.gform_wrapper .chosen-container-single .chosen-single,
.gform_wrapper .chosen-container-multi .chosen-choices,
.select2-container a.select2-choice,
.smile-icon-timeline-wrap .timeline-wrapper .timeline-block,
.smile-icon-timeline-wrap .timeline-feature-item.feat-item {
	background-color: #f2f4f5;
	}
.woocommerce #payment .payment_box:after,
.timeline-wrapper .timeline-post-right .ult-timeline-arrow l,
.timeline-wrapper .timeline-post-left .ult-timeline-arrow l,
.timeline-feature-item.feat-item .ult-timeline-arrow l {
	border-color: #f2f4f5;
	}

/* Border Color */
.l-submain,
.g-cols > div,
.w-blog-entry,
.w-bloglist,
.w-blognav,
.w-comments,
.w-comments-item,
.w-pricing-item-h,
.w-profile,
.w-sharing.type_simple .w-sharing-item,
.w-tabs.layout_accordion,
.w-tabs.layout_accordion .w-tabs-section,
.w-timeline.type_vertical .w-timeline-section-content,
#wp-calendar thead th,
#wp-calendar tbody td,
#wp-calendar tfoot td,
.widget_nav_menu .menu-item a,
#lang_sel a,
#lang_sel_click a,
.woocommerce table th,
.woocommerce table td,
.woocommerce .login,
.woocommerce .checkout_coupon,
.woocommerce .register,
.woocommerce .cart.variations_form,
.woocommerce .commentlist,
.woocommerce .commentlist li,
.woocommerce .comment-respond,
.woocommerce .related,
.woocommerce .upsells,
.woocommerce .cross-sells,
.woocommerce .checkout #order_review,
.woocommerce ul.order_details li,
.woocommerce .shop_table.my_account_orders,
.widget_price_filter .ui-slider-handle,
.widget_layered_nav ul,
.widget_layered_nav ul li,
#bbpress-forums .bbp-body ul.forum,
#bbpress-forums .bbp-forums li.bbp-header,
#bbpress-forums .bbp-body ul.topic,
#bbpress-forums .bbp-topics li.bbp-header,
.bbp-replies .bbp-body,
div.bbp-forum-header,
div.bbp-topic-header,
div.bbp-reply-header,
.bbp-pagination-links a,
.bbp-pagination-links span.current,
span.bbp-topic-pagination a.page-numbers,
.bbp-logged-in,
fieldset,
.form_saved_message,
.gform_wrapper .gsection,
.gform_wrapper .gf_page_steps,
.gform_wrapper li.gfield_creditcard_warning,
.smile-icon-timeline-wrap .timeline-line,
.ult_pricing_table_wrap.ult_design_6 .ult_pricing_table {
	border-color: #e4e8eb;
	}
.g-hr-h i,
.w-icon.color_border .w-icon-link,
.w-iconbox.color_light .w-iconbox-icon {
	color: #e4e8eb;
	}
.g-hr-h:before,
.g-hr-h:after,
.g-btn.color_default,
.g-btn.color_default.outlined:before,
.w-icon.color_border.type_solid .w-icon-link,
.w-iconbox.color_light.type_solid .w-iconbox-icon,
.w-timeline-list:before,
.woocommerce .button,
.no-touch .woocommerce .quantity .plus:hover,
.no-touch .woocommerce .quantity .minus:hover,
.widget_price_filter .ui-slider,
.gform_wrapper .gform_page_footer .gform_previous_button {
	background-color: #e4e8eb;
	}
.g-btn.color_default.outlined,
.pagination .page-numbers,
.w-iconbox.color_light.type_outlined .w-iconbox-icon,
.w-socials-item-link,
.w-tags-item-link,
.w-team-links-item,
.w-testimonial,
.woocommerce-pagination a,
.woocommerce-pagination span {
	box-shadow: 0 0 0 2px #e4e8eb inset;
	}

/* Heading Color */
h1, h2, h3, h4, h5, h6,
input[type="text"],
input[type="password"],
input[type="email"],
input[type="url"],
input[type="tel"],
input[type="number"],
input[type="date"],
input[type="search"],
textarea,
select,
.w-form-field > i,
.no-touch .g-btn.color_default:hover,
.no-touch .g-btn.color_default.outlined:hover,
.g-btn.color_contrast.outlined,
.w-blog-entry-title,
.w-counter-number,
.w-iconbox.color_contrast .w-iconbox-icon,
.w-portfolio-item-anchor,
.no-touch a.w-portfolio-item-anchor:hover,
.l-submain.color_primary a.w-portfolio-item-anchor,
.l-submain.color_secondary a.w-portfolio-item-anchor,
.w-pricing-item-title,
.w-pricing-item-price,
.woocommerce .product .price {
	color: #292e33;
	}
.g-btn.color_contrast,
.g-btn.color_contrast.outlined:before,
.w-iconbox.color_contrast.type_solid .w-iconbox-icon,
.w-progbar.color_contrast .w-progbar-bar-h {
	background-color: #292e33;
	}
.g-btn.color_contrast.outlined,
.w-iconbox.color_contrast.type_outlined .w-iconbox-icon {
	box-shadow: 0 0 0 2px #292e33 inset;
	}

/* Text Color */
.l-canvas,
.g-btn.color_default,
.g-btn.color_default.outlined,
.w-cart-dropdown,
.w-icon.color_text .w-icon-link,
.w-iconbox.color_light.type_solid .w-iconbox-icon,
.color_primary .w-icon.color_text.type_solid .w-icon-link,
.woocommerce .button,
.l-subfooter.at_top .woocommerce .button {
	color: #5c6166;
	}

/* Primary Color */
a,
.g-html .highlight,
.g-btn.color_primary.outlined,
.no-touch .w-blog-entry-link:hover .w-blog-entry-title-h,
.no-touch .w-blog-entry-link:hover,
.l-main .w-contacts-item i,
.w-counter.color_primary .w-counter-number,
.w-filters-item.active,
.no-touch .w-filters-item.active:hover,
.w-form-field > input:focus + i,
.w-form-field > textarea:focus + i,
.w-icon.color_primary .w-icon-link,
.w-iconbox.color_primary .w-iconbox-icon,
.no-touch .w-iconbox-link:hover .w-iconbox-title,
.no-touch .w-pagehead-nav-item:hover,
.w-tabs-item.active,
.no-touch .w-tabs-item.active:hover,
.w-tabs.layout_accordion .w-tabs-section.active .w-tabs-section-header,
.no-touch .w-tags-item-link:hover,
.w-team-link .w-team-name,
.no-touch .w-clients .slick-prev:hover,
.no-touch .w-clients .slick-next:hover,
.woocommerce .products .product .button,
.no-touch .woocommerce .products .product .button.loading:hover,
.woocommerce .star-rating span:before,
.woocommerce-breadcrumb a,
.woocommerce div.product .woocommerce-tabs .tabs li.active,
.no-touch .woocommerce div.product .woocommerce-tabs .tabs li.active:hover,
.woocommerce .stars span a:after,
#subscription-toggle span.is-subscribed:before,
#favorite-toggle span.is-favorite:before {
	color: #6dbcdb;
	}
.l-submain.color_primary,
.w-actionbox.color_primary,
.g-btn.color_primary,
.g-btn.color_primary.outlined:before,
button,
input[type="submit"],
.no-touch .pagination .page-numbers:before,
.pagination .page-numbers.current,
.no-touch .w-iconbox.type_outlined .w-iconbox-icon:before,
.no-touch .w-iconbox.type_solid .w-iconbox-icon:before,
.w-iconbox.color_primary.type_solid .w-iconbox-icon,
.no-touch .w-filters-item:hover,
.w-icon.color_primary.type_solid .w-icon-link,
.w-pricing-item.type_featured .w-pricing-item-title,
.w-pricing-item.type_featured .w-pricing-item-price,
.w-progbar.color_primary .w-progbar-bar-h,
.no-touch .w-team-links,
.w-timeline-item:before,
.w-timeline.type_vertical .w-timeline-section:before,
.w-timeline-section-title:before,
.w-timeline-section.active .w-timeline-section-title:before,
.no-touch .w-toplink.active:hover,
.no-touch .tp-leftarrow.custom:before,
.no-touch .tp-rightarrow.custom:before,
.widget_nav_menu .menu-item.current-menu-item > a,
.no-touch .widget_nav_menu .menu-item.current-menu-item > a:hover,
p.demo_store,
.woocommerce .button.alt,
.woocommerce .button.checkout,
.no-touch .woocommerce .products .product .button:hover,
.no-touch .woocommerce-pagination a:hover,
.woocommerce-pagination span.current,
.woocommerce .onsale,
.widget_price_filter .ui-slider-range,
.widget_layered_nav ul li.chosen,
.widget_layered_nav_filters ul li a,
.no-touch .bbp-pagination-links a:hover,
.bbp-pagination-links span.current,
.no-touch span.bbp-topic-pagination a.page-numbers:hover,
.gform_wrapper .gf_progressbar_percentage,
.gform_wrapper .gform_page_footer .gform_next_button,
.gform_wrapper .chosen-container .chosen-results li.highlighted,
.smile-icon-timeline-wrap .timeline-separator-text .sep-text,
.smile-icon-timeline-wrap .timeline-wrapper .timeline-dot,
.smile-icon-timeline-wrap .timeline-feature-item .timeline-dot {
	background-color: #6dbcdb;
	}
.g-html blockquote,
.w-blog-entry.sticky,
.no-touch .w-clients-item-h:hover,
.w-filters-item.active,
.w-tabs-item.active,
.no-touch .w-tabs-item.active:hover,
.widget_nav_menu .menu-item.current-menu-item > a,
.fotorama__thumb-border,
.woocommerce div.product .woocommerce-tabs .tabs li.active,
.widget_layered_nav ul li.chosen,
.no-touch .bbp-pagination-links a:hover,
.bbp-pagination-links span.current,
.no-touch span.bbp-topic-pagination a.page-numbers:hover,
#bbp-user-navigation li.current {
	border-color: #6dbcdb;
	}
input[type="text"]:focus,
input[type="password"]:focus,
input[type="email"]:focus,
input[type="url"]:focus,
input[type="tel"]:focus,
input[type="number"]:focus,
input[type="date"]:focus,
input[type="search"]:focus,
textarea:focus,
select:focus,
#bbpress-forums div.bbp-the-content-wrapper textarea:focus {
	box-shadow: 0 0 0 2px #6dbcdb;
	}
.g-btn.color_primary.outlined,
.l-main .w-contacts-item i,
.w-iconbox.color_primary.type_outlined .w-iconbox-icon,
.no-touch .w-pagehead-nav-item:hover,
.no-touch .w-tags-item-link:hover,
.no-touch .w-testimonial:hover,
.w-timeline-item,
.w-timeline-section-title,
.no-touch .w-clients .slick-prev:hover,
.no-touch .w-clients .slick-next:hover,
.woocommerce .products .product .button,
.no-touch .woocommerce .products .product .button.loading:hover {
	box-shadow: 0 0 0 2px #6dbcdb inset;
	}

/* Secondary Color */
.no-touch a:hover,
.g-btn.color_secondary.outlined,
.no-touch .w-blog.type_masonry .w-blog-meta a:hover,
.no-touch .w-blognav-prev:hover .w-blognav-title,
.no-touch .w-blognav-next:hover .w-blognav-title,
.w-counter.color_secondary .w-counter-number,
.w-icon.color_secondary .w-icon-link,
.w-iconbox.color_secondary .w-iconbox-icon,
.no-touch .w-team-link:hover .w-team-name,
.no-touch .widget_archive ul li a:hover,
.no-touch .widget_categories ul li a:hover,
.no-touch .widget_tag_cloud .tagcloud a:hover,
.no-touch .woocommerce-breadcrumb a:hover,
.no-touch .widget_product_tag_cloud .tagcloud a:hover,
.no-touch .bbp_widget_login a.button.logout-link:hover {
	color: #fc4349;
	}
.l-submain.color_secondary,
.w-actionbox.color_secondary,
.g-btn.color_secondary,
.g-btn.color_secondary.outlined:before,
.w-icon.color_secondary.type_solid .w-icon-link,
.w-iconbox.color_secondary.type_solid .w-iconbox-icon,
.w-progbar.color_secondary .w-progbar-bar-h,
.no-touch .button:hover,
.no-touch input[type="submit"]:hover,
.no-touch .woocommerce .button:hover,
.no-touch .woocommerce .button.alt:hover,
.no-touch .woocommerce .button.checkout:hover,
.no-touch .woocommerce .shop_table.cart .product-remove a:hover,
.no-touch .widget_layered_nav_filters ul li a:hover,
.no-touch .bbp-row-actions #favorite-toggle a:hover,
.no-touch .bbp-row-actions #subscription-toggle a:hover {
	background-color: #fc4349;
	}
.g-btn.color_secondary.outlined,
.w-iconbox.color_secondary.type_outlined .w-iconbox-icon {
	box-shadow: 0 0 0 2px #fc4349 inset;
	}

/* Fade Elements Color */
.w-blog-meta,
.w-blog-meta a,
.w-icon.color_fade .w-icon-link,
.no-touch .w-icon.color_fade.type_solid .w-icon-link:hover,
.w-profile-link.for_logout,
.widget_tag_cloud .tagcloud a,
.woocommerce-breadcrumb,
.woocommerce .star-rating:before,
.woocommerce .stars span:after,
.woocommerce table.shop_table .product-remove a.remove,
.widget_product_tag_cloud .tagcloud a,
p.bbp-topic-meta,
.bbp_widget_login a.button.logout-link {
	color: #a4abb3;
	}
.w-shortblog-entry-date {
	box-shadow: 0 0 0 2px #a4abb3 inset;
	}

/*************************** ALTERNATE CONTENT ***************************/

/* Background Color */
.l-submain.color_alternate,
.color_alternate .w-blog.type_masonry .w-blog-entry-preview:before,
.color_alternate .w-filters-item.active,
.no-touch .color_alternate .w-filters-item.active:hover,
.color_alternate .w-tabs-item.active,
.no-touch .color_alternate .w-tabs-item.active:hover,
.color_alternate .w-timeline-item,
.color_alternate .w-timeline-section-title,
.color_alternate .w-timeline.type_vertical .w-timeline-section-content {
	background-color: #f2f4f5;
	}
.color_alternate .g-btn.color_contrast,
.no-touch .color_alternate .g-btn.color_contrast:hover,
.no-touch .color_alternate .g-btn.color_contrast.outlined:hover,
.color_alternate .w-icon.color_border.type_solid .w-icon-link {
	color: #f2f4f5;
	}

/* Alternate Background Color */
.color_alternate input[type="text"],
.color_alternate input[type="password"],
.color_alternate input[type="email"],
.color_alternate input[type="url"],
.color_alternate input[type="tel"],
.color_alternate input[type="number"],
.color_alternate input[type="date"],
.color_alternate input[type="search"],
.color_alternate textarea,
.color_alternate select,
.color_alternate .w-blog.imgpos_atleft .w-blog-entry-preview-icon,
.color_alternate .w-filters,
.color_alternate .w-icon.color_text.type_solid .w-icon-link,
.color_alternate .w-icon.color_fade.type_solid .w-icon-link,
.color_alternate .w-pricing-item-title,
.color_alternate .w-pricing-item-price,
.color_alternate .w-tabs-list {
	background-color: #fff;
	}

/* Border Color */
.color_alternate .w-blog-entry,
.color_alternate .w-bloglist,
.color_alternate .w-comments-item,
.color_alternate .w-pricing-item-h,
.color_alternate .w-tabs.layout_accordion,
.color_alternate .w-tabs.layout_accordion .w-tabs-section,
.color_alternate .w-timeline.type_vertical .w-timeline-section-content {
	border-color: #e3e6e8;
	}
.color_alternate .g-hr-h i,
.color_alternate .page-404 i,
.color_alternate .w-icon.color_border .w-icon-link {
	color: #e3e6e8;
	}
.color_alternate .g-hr-h:before,
.color_alternate .g-hr-h:after,
.color_alternate .g-btn.color_default,
.color_alternate .g-btn.color_default.outlined:before,
.color_alternate .w-icon.color_border.type_solid .w-icon-link,
.color_alternate .w-timeline-list:before {
	background-color: #e3e6e8;
	}
.color_alternate .g-btn.color_default.outlined,
.color_alternate .pagination .page-numbers,
.color_alternate .w-socials-item-link,
.color_alternate .w-tags-item-link,
.color_alternate .w-team-links-item,
.color_alternate .w-testimonial {
	box-shadow: 0 0 0 2px #e3e6e8 inset;
	}

/* Heading Color */
.color_alternate h1,
.color_alternate h2,
.color_alternate h3,
.color_alternate h4,
.color_alternate h5,
.color_alternate h6,
.color_alternate input[type="text"],
.color_alternate input[type="password"],
.color_alternate input[type="email"],
.color_alternate input[type="url"],
.color_alternate input[type="tel"],
.color_alternate input[type="number"],
.color_alternate input[type="date"],
.color_alternate textarea,
.color_alternate select,
.color_alternate .w-form-field > i,
.no-touch .color_alternate .g-btn.color_default:hover,
.no-touch .color_alternate .g-btn.color_default.outlined:hover,
.color_alternate .g-btn.color_contrast.outlined,
.color_alternate .w-blog-entry-title,
.color_alternate .w-counter-number,
.color_alternate .w-pricing-item-title,
.color_alternate .w-pricing-item-price {
	color: #292e33;
	}
.color_alternate .g-btn.color_contrast,
.color_alternate .g-btn.color_contrast.outlined:before {
	background-color: #292e33;
	}
.color_alternate .g-btn.color_contrast.outlined {
	box-shadow: 0 0 0 2px #292e33 inset;
	}

/* Text Color */
.l-submain.color_alternate,
.color_alternate .g-btn.color_default,
.color_alternate .g-btn.color_default.outlined,
.color_alternate .w-icon.color_text .w-icon-link {
	color: #5c6166;
	}

/* Primary Color */
.color_alternate a,
.color_alternate .g-btn.color_primary.outlined,
.no-touch .color_alternate .w-blog-entry-link:hover .w-blog-entry-title-h,
.no-touch .color_alternate .w-blog-entry-link:hover,
.color_alternate .l-main .w-contacts-item i,
.color_alternate .w-counter.color_primary .w-counter-number,
.color_alternate .w-filters-item.active,
.no-touch .color_alternate .w-filters-item.active:hover,
.color_alternate .w-icon.color_primary .w-icon-link,
.color_alternate .w-iconbox-icon,
.no-touch .color_alternate .w-iconbox-link:hover .w-iconbox-title,
.no-touch .color_alternate .w-pagehead-nav-item:hover,
.color_alternate .w-tabs-item.active,
.no-touch .color_alternate .w-tabs-item.active:hover,
.color_alternate .w-tabs.layout_accordion .w-tabs-section.active .w-tabs-section-header,
.no-touch .color_alternate .w-tags-item-link:hover,
.color_alternate .w-team-link .w-team-name {
	color: #6dbcdb;
	}
.color_alternate .g-btn.color_primary,
.color_alternate .g-btn.color_primary.outlined:before,
.color_alternate input[type="submit"],
.no-touch .color_alternate .pagination .page-numbers:before,
.color_alternate .pagination .page-numbers.current,
.no-touch .color_alternate .w-filters-item:hover,
.color_alternate .w-icon.color_primary.type_solid .w-icon-link,
.color_alternate .w-pricing-item.type_featured .w-pricing-item-title,
.color_alternate .w-pricing-item.type_featured .w-pricing-item-price,
.no-touch .color_alternate .w-team-links,
.color_alternate .w-timeline-item:before,
.color_alternate .w-timeline.type_vertical .w-timeline-section:before,
.color_alternate .w-timeline-section-title:before,
.color_alternate .w-timeline-section.active .w-timeline-section-title:before {
	background-color: #6dbcdb;
	}
.color_alternate .g-html blockquote,
.color_alternate .w-blog-entry.sticky,
.color_alternate .w-filters-item.active,
.color_alternate .w-tabs-item.active,
.no-touch .color_alternate .w-tabs-item.active:hover {
	border-color: #6dbcdb;
	}
.color_alternate input[type="text"]:focus,
.color_alternate input[type="password"]:focus,
.color_alternate input[type="email"]:focus,
.color_alternate input[type="url"]:focus,
.color_alternate input[type="tel"]:focus,
.color_alternate input[type="number"]:focus,
.color_alternate input[type="date"]:focus,
.color_alternate input[type="search"]:focus,
.color_alternate textarea:focus,
.color_alternate select:focus {
	box-shadow: 0 0 0 2px #6dbcdb;
	}
.color_alternate .g-btn.color_primary.outlined,
.color_alternate .l-main .w-contacts-item i,
.no-touch .color_alternate .w-pagehead-nav-item:hover,
.no-touch .color_alternate .w-tags-item-link:hover,
.no-touch .color_alternate .w-testimonial:hover,
.color_alternate .w-timeline-item:before,
.color_alternate .w-timeline-section-title:before {
	box-shadow: 0 0 0 2px #6dbcdb inset;
	}

/* Secondary Color */
.no-touch .color_alternate a:hover,
.no-touch .color_alternate a:active,
.color_alternate .g-btn.color_secondary.outlined,
.no-touch .color_alternate .w-blog.type_masonry .w-blog-meta a:hover,
.color_alternate .w-counter.color_secondary .w-counter-number,
.color_alternate .w-icon.color_secondary .w-icon-link,
.no-touch .color_alternate .w-team-link:hover .w-team-name {
	color: #fc4349;
	}
.color_alternate .g-btn.color_secondary,
.color_alternate .g-btn.color_secondary.outlined:before,
.color_alternate .w-icon.color_secondary.type_solid .w-icon-link {
	background-color: #fc4349;
	}
.color_alternate .g-btn.color_secondary.outlined {
	box-shadow: 0 0 0 2px #fc4349 inset;
	}

/* Fade Elements Color */
.color_alternate .w-blog-meta,
.color_alternate .w-blog-meta a,
.color_alternate .w-bloglist-entry-date,
.color_alternate .w-bloglist-entry-author,
.color_alternate .w-icon.color_fade .w-icon-link {
	color: #a4abb3;
	}
.color_alternate .w-shortblog-entry-date {
	box-shadow: 0 0 0 2px #a4abb3 inset;
	}

/*************************** SUBFOOTER ***************************/

/* Background Color */
.l-subfooter.at_top,
.l-subfooter.at_top #lang_sel ul ul,
.l-subfooter.at_top #lang_sel_click ul ul {
	background-color: #2c3e50;
	}
	
/* Alternate Background Color */
.l-subfooter.at_top input[type="text"],
.l-subfooter.at_top input[type="password"],
.l-subfooter.at_top input[type="email"],
.l-subfooter.at_top input[type="url"],
.l-subfooter.at_top input[type="tel"],
.l-subfooter.at_top input[type="number"],
.l-subfooter.at_top input[type="date"],
.l-subfooter.at_top input[type="search"],
.l-subfooter.at_top textarea,
.l-subfooter.at_top select,
.no-touch .l-subfooter.at_top #lang_sel a:hover,
.no-touch .l-subfooter.at_top #lang_sel_click a:hover {
	background-color: #384b5f;
	}

/* Border Color */
.l-subfooter.at_top,
.l-subfooter.at_top #wp-calendar thead th,
.l-subfooter.at_top #wp-calendar tbody td,
.l-subfooter.at_top #wp-calendar tfoot td,
.l-subfooter.at_top #lang_sel a,
.l-subfooter.at_top #lang_sel_click a,
.l-subfooter.at_top .widget_nav_menu .menu-item a {
	border-color: #384b5f;
	}
.l-subfooter.at_top .w-socials-item-link {
	box-shadow: 0 0 0 2px #384b5f inset;
	}

/* Text Color */
.l-subfooter.at_top,
.l-subfooter.at_top .w-socials-item-link {
	color: #939da3;
	}

/* Heading Color */
.l-subfooter.at_top h1,
.l-subfooter.at_top h2,
.l-subfooter.at_top h3,
.l-subfooter.at_top h4,
.l-subfooter.at_top h5,
.l-subfooter.at_top h6,
.l-subfooter.at_top .w-form-field > i,
.l-subfooter.at_top input[type="text"],
.l-subfooter.at_top input[type="password"],
.l-subfooter.at_top input[type="email"],
.l-subfooter.at_top input[type="url"],
.l-subfooter.at_top input[type="tel"],
.l-subfooter.at_top input[type="number"],
.l-subfooter.at_top input[type="date"],
.l-subfooter.at_top input[type="search"],
.l-subfooter.at_top textarea,
.l-subfooter.at_top select {
	color: #d3d8db;
	}

/* Link Color */
.l-subfooter.at_top a,
.l-subfooter.at_top .widget_tag_cloud .tagcloud a,
.l-subfooter.at_top .widget_product_tag_cloud .tagcloud a {
	color: #6dbcdb;
	}

/* Link Hover Color */
.no-touch .l-subfooter.at_top a:hover,
.no-touch .l-subfooter.at_top .widget_tag_cloud .tagcloud a:hover,
.no-touch .l-subfooter.at_top .widget_product_tag_cloud .tagcloud a:hover {
	color: #fc4349;
	}

/*************************** FOOTER ***************************/

/* Background Color */
.l-subfooter.at_bottom {
	background-color: #ffffff;
	}

/* Text Color */
.l-subfooter.at_bottom {
	color: #939da3;
	}

/* Link Color */
.l-subfooter.at_bottom a {
	color: #d3d8db;
	}

/* Link Hover Color */
.no-touch .l-subfooter.at_bottom a:hover {
	color: #fc4349;
	}
</style>
</head>


<link rel='stylesheet' id='mesh-nightmode-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates-child/mystyles/Default.css?ver=2.1.4' type='text/css' media='all' /><body class="page-template page-template-pg-seriesfinder page-template-pg-seriesfinder-php page page-id-3370 l-body us-theme_nd_4-4-8">


<!-- CANVAS -->

<div class="l-canvas  col_contside">

	<!-- HEADER -->
	<div class="l-header">

		<div class="l-subheader at_top" style="line-height: 36px; ">
			<div class="l-subheader-h i-cf">
			</div>
		</div>
				<div class="l-subheader at_middle"  style="line-height: 60px;">
			<div class="l-subheader-h i-widgets i-cf">

				<div class="w-logo  with_title">
					<a class="w-logo-link" href="https://www.novelupdates.com/">
												<span class="w-logo-title">Novel Updates</span>
					</a>
				</div>

				                
                   
			    <script>$( document ).ready(function() {$.get( "https://forum.novelupdates.com/nufonline_guest.php" );});</script>                
                <span class="menu_username_right" id="menu_username_right"><span class="menu_right_icons notloggedin"><a class="menu_right_links" href="/register/">Register</a></span><span class="menu_right_icons notloggedin"><a class="menu_right_links" href="/login/">Login</a></span><span id="gs_menu" class="menu_right_icons search guest" onclick="toggleUserMenu();"><i id="gs_menu" class="fa fa-user-circle menu" aria-hidden="true" dp="yes"></i><i id="gs_menu" class="fa fa-angle-down menu" aria-hidden="true" dp="yes"></i></span><div title="Guest Menu" class="lo_main_themain guest" style="display: inline; position:relative;">
						<div class="lo_main">
						<div class="arrow-up_toc_username"></div>
						<div class="lo_box">
						<span><span class="user_menu_links atpmn" id="user_menu_links">Theme <select class="wi_themes" id="wi_themes" dp="yes" onchange="toggleTheme(this.value);"><option value="Default" dp="yes">Default</option><option value="Dark" dp="yes">Dark</option><option value="Dark - Crisp" dp="yes">Dark Crisp</option></select></span></span>
						</div></div>
						</div><span class="menu_right_icons search" title="Search" onclick="show_search_bar(this);"><i class="fa fa-search menu"></i></span><span class="menu_right_icons mobile" style="display:none;" onclick="toggleMobileMenu();"><i class="fa fa-bars menu" aria-hidden="true"></i></span><span class="nu_menu_search"><div class="ab_title" dp="yes">Search</div><form id="search_nu_novel" action="https://www.novelupdates.com/"><span class="w-search-input-h">
						<div class="srh_menu">
						<input id="searchme-rl newmenu" class="nusearchnovel" autocomplete="off" onkeyup="showSearch(this.value)" type="text" value="" name="s" placeholder="Search...">
						</div>
						<input type="hidden" name="post_type" value="seriesplans">
							
						<div class="w-search-submit">
						<span class="search_menu_right"><i class="fa fa-search btn" onclick="search_btn_click();"></i></span>
						
						<div class ="search_type">			
						<span style="padding-right: 35px;">
						<span onclick="g_checkmark(this);" id="rl_checkbox_main" class="s_series islbl" value="1"><i dp="yes" class="fa fa-check" value="1" aria-hidden="true"></i></span>
						<span onclick="g_check_lbl(this);" style="position: relative; top: -1px ;padding-left: 2px; cursor:pointer;">Series</span></span>
					
						<span>
						<span onclick="g_checkmark(this);" id="rl_checkbox_main" class="s_users islbl" value="0">
						<i class="fa fa-check icon-invisible" aria-hidden="true" value="0"></i></span>
						<span onclick="g_check_lbl(this);" style="position: relative; top: -1px ;padding-left: 2px; cursor:pointer;">Group</span></span>
						</div>
						
						<div style="font-size: 13px; padding-right:5px; position: relative; top: 15px;"><a href="https://www.novelupdates.com/series-finder/">Series Finder</a></div>
						</div>
							
						</span><div class="livesearchresult" style="display: none; z-index: 2;"></div></form></span></span>                
            
				<!-- NAV -->
				<nav class="w-nav layout_hor animation_">
					<ul class="w-nav-list level_1 ">

                        <div class="menu_items_display" style="display:none;"><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item menu-item-73"><a class="w-nav-anchor level_1" href="/login/"><span class="w-nav-title">Login</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item menu-item-73"><a class="w-nav-anchor level_1" href="/register/"><span class="w-nav-title">Register</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="http://forum.novelupdates.com/"><span class="w-nav-title">Forum</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="//www.novelupdates.com/random-novel/"><span class="w-nav-title">Random Novel</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="//www.novelupdates.com/series-finder/"><span class="w-nav-title">Series Finder</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="//www.novelupdates.com/novelslisting/"><span class="w-nav-title">Series Listing</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="//www.novelupdates.com/series-ranking/"><span class="w-nav-title">Series Ranking</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="//www.novelupdates.com/latest-series/"><span class="w-nav-title">Latest Series</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="https://www.novelupdates.com/recommendation-lists/"><span class="w-nav-title">Rec Lists</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><span class="w-an_cus">Theme</span> <span class="c_theme"><select class="wi_themes" id="wi_themes" dp="yes" onchange="toggleTheme(this.value);"><option value="Default" dp="yes">Default</option><option value="Dark" dp="yes">Dark</option><option value="Dark - Crisp" dp="yes">Dark Crisp</option></select></span></li></div>                    </ul>
				</nav><!-- /NAV -->
				
							</div>
		</div>


	</div>
	<!-- /HEADER -->
    <div class="ol_account_main" style="position:relative; display: none;">
			<div class="ol_account"><span class="ol_title">Your changes has been saved.</span></div></div>	<!-- MAIN -->
	<div class="l-main">




<link rel="stylesheet" href="https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/gadient/style.css"><link rel="stylesheet" href="https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/genre_search.css"><link rel="stylesheet" href="https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/chosen.min.css"><link rel="stylesheet" href="//code.jquery.com/ui/1.11.4/themes/smoothness/jquery-ui.css"><script src="//ajax.googleapis.com/ajax/libs/jqueryui/1.11.1/jquery-ui.min.js"></script>

<style>
.select2
{
	width:100% !important;
}
input.select2-search__field
{
	box-shadow:none !important;
}
.select2-container--bootstrap .select2-selection--multiple .select2-selection__choice
{
	margin: 10px 0 0 6px !important;
}
.select2-container--bootstrap .select2-selection--multiple .select2-search--inline .select2-search__field
{
	margin-top: 3px !important;
}
</style> 

	<div class="l-submain">
		<div class="l-submain-h g-html i-cf other">
			<div class="l-content">
				<div class="w-blog post-3370 page type-page status-publish hentry">
              

 	 	 
					<div class="w-blog-content other">

     
<h3 class="mypage bc top">Series Finder</h3>
<div class="breadcrumb_nu">
		<div id="Breadcrumb_Top" style="max-width:1420px; margin:0px auto;">
				<dl>
					<dd><a href="//www.novelupdates.com/" title="Home">Home</a>&nbsp;&gt;&nbsp;</dd>
          			<dd>Series Finder</dd>
				</dl>
	
		</div>	
</div>
<div class="ricon sf bc" style="margin-top: -72px;""> <i class="fa fa-cogs sfinder" style="margin-right: -2px;" aria-hidden="true"></i> <span class="ftext sf"> [ <a href="#" class="rlhidesme" datahide="no" onclick="hiderk(this)">Filters</a> ] </span></div>


<form style="padding-bottom: 20px;" id="rankfilter" name="rankfilter" action="/series-finder/?sf=1&sort=sdate&order=desc" method="post">
<div class="sf_helpinfo">
<ul style="margin-bottom: 5px;">
<li><b>Chapters</b> - This is the number of releases (chapters).</li>
	<ul style="margin-bottom: 5px;">
	<li>Min - Minimum amount of chapters.</li>
	<li>Max - Maximum amount of chapters.</li>
	</ul>
<li><b>Release Frequency</b> - The release frequency of a novel. Higher frequency means the novel is updated more often.</li>
	<ul style="margin-bottom: 5px;">
	<li>Min - Minimum release frequency. Lower frequency = slower updates. </li>
	<li>Max - Maximum release frequency. Higher frequency = quicker updates.</li>
	</ul>
<li><b>Rating</b> - Novel rating on a scale of 1 to 5.</li>
	<ul style="margin-bottom: 5px;">
	<li>Min - Minimum rating (1 to 5)</li>
	<li>Max - Maximum rating (1 to 5)</li>
	</ul>
<li><b>Number of Ratings</b> - The amount of ratings for a novel.</li>
	<ul style="margin-bottom: 5px;">
	<li>Min - Minimum number of ratings.</li>
	<li>Max - Maximum number of ratings.</li>
	</ul>
<li><b>Readers</b> - The number of readers a novel has.</li>
	<ul style="margin-bottom: 5px;">
	<li>Min - Minimum readers.</li>
	<li>Max - Maximum readers.</li>
	</ul>
<li><b>Last Release Date</b> - The last release date of the latest novel.</li> 
	<ul style="margin-bottom: 5px;">
	<li>Min - Minimum last release date.</li>
	<li>Max - Maximum last release date.</li>
	</ul>
<li><b>Genre</b> - The genres of the novel.</li>
	<ul style="margin-bottom: 5px;">
	<li>Include - Click a genre once to include it in your search.</li>
	<li>Exclude - Click a genre twice to exclude it in your search.</li>
	</ul>
<li><b>Tags</b> - The tags of a novel.</li>
	<ul style="margin-bottom: 5px;">
	<li>OR - Uses the 'OR' condition when searching for tags.</li>
	<li>AND - Uses the 'AND' condition when searching for tags.</li>
	</ul>
<li><b>Hide Series from... (Members Only)</b> - Prevents novels from showing up from the selected list. This is for members using the reading list feature.</li> 
</ul>
</div>

<div class="rankfl">
<span class="rktext sf bc" style="padding-top: 15px;">Current Filters - <span class="clist list"><b>Order By:</b></span> Last Updated, Descending</span>
<span class="sftitle sfpad new">Novel Type <i onclick="sf_getnovelinfo('noveltype', 'Novel Type');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span>   

<div class="g-cols wpb_row offset_default"><div class="one-third" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			
<ul class="genrematch rank"><li><a class="typerank" href="javascript:void(0);" genreid="2443" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Light Novel</a></li></ul></div> 
	</div> </div><div class="one-third" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			

<ul class="genrematch rank"><li></i><a class="typerank" href="javascript:void(0);" genreid="26874" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Published Novel</a></li></ul></div> 
	</div> </div><div class="one-third" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			

<ul class="genrematch rank"><li></i><a class="typerank" href="javascript:void(0);" genreid="2444" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Web Novel</a></li></ul></div> 
	</div> </div></div>
                    
                  
<span class="sftitle sfpad new">Language <i onclick="sf_getnovelinfo('language', 'Language');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span>   
 

<div class="g-cols wpb_row offset_default"><div class="one-third" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			
<ul class="genrematch rank"><li><a class="langrank" href="javascript:void(0);" genreid="495" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Chinese</a></li><li><a class="langrank" href="javascript:void(0);" genreid="9181" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Filipino</a></li><li><a class="langrank" href="javascript:void(0);" genreid="9179" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Indonesian</a></li></ul></div> 
	</div> </div><div class="one-third" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			

<ul class="genrematch rank"><li></i><a class="langrank" href="javascript:void(0);" genreid="496" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Japanese</a></li><li></i><a class="langrank" href="javascript:void(0);" genreid="18657" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Khmer</a></li><li></i><a class="langrank" href="javascript:void(0);" genreid="497" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Korean</a></li></ul></div> 
	</div> </div><div class="one-third" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			

<ul class="genrematch rank"><li></i><a class="langrank" href="javascript:void(0);" genreid="9183" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Malaysian</a></li><li></i><a class="langrank" href="javascript:void(0);" genreid="9954" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Thai</a></li><li></i><a class="langrank" href="javascript:void(0);" genreid="9177" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Vietnamese</a></li></ul></div> 
	</div> <div class="g-hr type_invisible no_icon">
						<span class="g-hr-h">
							<i class="fa fa-"></i>
						</span>
					</div></div></div>
                    
 


<span class="sftitle sfpad new">Chapters [<select name="releases_mm" id="r_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>] <i onclick="sf_getnovelinfo('releases', 'Chapters');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
<input autocomplete="off" class="inputrank finder" type="text" name="rk_releases" id="rk_releases" value="">

<span class="sftitle sfpad new">Release Frequency [<select name="releases_mm" id="rf_mm" class="select sf">
						<option value="min">min</option>
						<option value="max" selected>max</option>
					</select>]<i onclick="sf_getnovelinfo('frequency', 'Release Frequency');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
<input autocomplete="off" class="inputrank finder" type="text" name="rk_rfreq" id="rk_rfreq" value=""> 

<span class="sftitle sfpad new">Reviews [<select name="releases_mm" id="rvc_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>]</span> 
<input autocomplete="off" class="inputrank finder" type="text" name="rk_rreviews" id="rk_rreviews" value=""> 

<span class="sftitle sfpad new">Rating  [<select name="releases_mm" id="rt_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>]<i onclick="sf_getnovelinfo('rating', 'Ratings');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
<input autocomplete="off" class="inputrank finder" type="text" name="rk_sr" id="rk_sr" value=""> 
<span class="sftitle sfpad new">Number of Ratings [<select name="releases_mm" id="rtc_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>] <i onclick="sf_getnovelinfo('numberratings', 'Number of Ratings');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
<input autocomplete="off" class="inputrank finder" type="text" name="rk_rcount" id="rk_rcount" value="">  
<span class="sftitle sfpad new">Readers [<select name="releases_mm" id="rct_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>] <i onclick="sf_getnovelinfo('readers', 'Readers');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
<input autocomplete="off" class="inputrank finder" type="text" name="rk_sread" id="rk_sread" value="">  


<span class="sftitle sfpad new">First Release Date  [<select name="releases_mm" id="dtf_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>]</span> 
 <p>
 <input autocomplete="off" class="inputrank sf" type="text" name="ardate_first" id="ardate_first" value="">
</p>

<span class="sftitle sfpad new">Last Release Date  [<select name="releases_mm" id="dt_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>] <i onclick="sf_getnovelinfo('lastdate', 'Last Release Date');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
 <p>
 <input autocomplete="off" class="inputrank sf" type="text" name="ardate" id="ardate" value="">
</p>


<span class="sftitle sfpad new">Genre [<select name="releases_mm" id="gi_mm" class="select sf">
						<option value="and" selected>AND</option>
						<option value="or">OR</option>
					</select>]<i onclick="sf_getnovelinfo('genre', 'Genre');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
<div style="font-size:12px;">Click once to include genre. Click the same genre twice to exclude.</div>
<div class="g-cols wpb_row offset_default"><div class="one-quarter" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			
<ul class="genrematch rank"><li><a class="genreme" href="javascript:void(0);" genreid="8" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Action</a></li><li><a class="genreme" href="javascript:void(0);" genreid="280" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Adult</a></li><li><a class="genreme" href="javascript:void(0);" genreid="13" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Adventure</a></li><li><a class="genreme" href="javascript:void(0);" genreid="17" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Comedy</a></li><li><a class="genreme" href="javascript:void(0);" genreid="9" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Drama</a></li><li><a class="genreme" href="javascript:void(0);" genreid="292" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Ecchi</a></li><li><a class="genreme" href="javascript:void(0);" genreid="5" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Fantasy</a></li><li><a class="genreme" href="javascript:void(0);" genreid="168" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Gender Bender</a></li><li><a class="genreme" href="javascript:void(0);" genreid="3" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Harem</a></li></ul></div> 
	</div> </div><div class="one-quarter" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			

<ul class="genrematch rank"><li></i><a class="genreme" href="javascript:void(0);" genreid="330" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Historical</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="343" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Horror</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="324" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Josei</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="14" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Martial Arts</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="4" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Mature</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="10" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Mecha</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="245" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Mystery</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="486" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Psychological</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="15" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Romance</a></li></ul></div> 
	</div> </div><div class="one-quarter" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			
<ul class="genrematch rank"><li></i><a class="genreme" href="javascript:void(0);" genreid="6" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>School Life</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="11" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Sci-fi</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="18" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Seinen</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="157" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Shoujo</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="851" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Shoujo Ai</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="12" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Shounen</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="1692" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Shounen Ai</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="7" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Slice of Life</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="281" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Smut</a></li></ul></div> 
	</div> 
    
  </div>
            
            
       <div class="one-quarter" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			
<ul class="genrematch rank"><li></i><a class="genreme" href="javascript:void(0);" genreid="1357" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Sports</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="16" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Supernatural</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="132" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Tragedy</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="479" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Wuxia</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="480" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Xianxia</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="3954" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Xuanhuan</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="560" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Yaoi</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="922" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Yuri</a></li></ul></div> 
	</div> </div>     
 
                    </div>
				<div class="tagfix" style="top:0px;">           
<span class="sftitle sfpad new"> Tags [<select name="releases_mm" id="tgi_mm" class="select sf">
						<option value="and">AND</option>
						<option value="or" selected>OR</option>
					</select>]<i onclick="sf_getnovelinfo('tags', 'Tags');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 

<div class="extrasf"></div>  
<div class="g-cols wpb_row offset_default">
<div class="one-half">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			 <select multiple id="tags_include" class="chzn-select" data-placeholder="Include..." name="tagsinclude&#091;&#093;"><option value="16185">Abandoned Children</option><option value="4859">Ability Steal</option><option value="1248">Absent Parents</option><option value="11325">Abusive Characters</option><option value="4885">Academy</option><option value="475">Accelerated Growth</option><option value="13403">Acting</option><option value="4976">Adapted from Manga</option><option value="6280">Adapted from Manhua</option><option value="269">Adapted to Anime</option><option value="2999">Adapted to Drama</option><option value="928">Adapted to Drama CD</option><option value="891">Adapted to Game</option><option value="270">Adapted to Manga</option><option value="182">Adapted to Manhua</option><option value="2721">Adapted to Manhwa</option><option value="1133">Adapted to Movie</option><option value="1334">Adapted to Visual Novel</option><option value="858">Adopted Children</option><option value="603">Adopted Protagonist</option><option value="4974">Adultery</option><option value="3108">Adventurers</option><option value="3802">Affair</option><option value="357">Age Progression</option><option value="216">Age Regression</option><option value="13720">Aggressive Characters</option><option value="234">Alchemy</option><option value="621">Aliens</option><option value="598">All-Girls School</option><option value="7909">Alternate World</option><option value="625">Amnesia</option><option value="774">Amusement Park</option><option value="4594">Anal</option><option value="476">Ancient China</option><option value="2662">Ancient Times</option><option value="13263">Androgynous Characters</option><option value="892">Androids</option><option value="1327">Angels</option><option value="255">Animal Characteristics</option><option value="7323">Animal Rearing</option><option value="724">Anti-Magic</option><option value="1732">Anti-social Protagonist</option><option value="4297">Antihero Protagonist</option><option value="203">Antique Shop</option><option value="513">Apartment Life</option><option value="12913">Apathetic Protagonist</option><option value="217">Apocalypse</option><option value="721">Appearance Changes</option><option value="430">Appearance Different from Actual Age</option><option value="697">Archery</option><option value="6442">Aristocracy</option><option value="9591">Arms Dealers</option><option value="306">Army</option><option value="9522">Army Building</option><option value="2676">Arranged Marriage</option><option value="13427">Arrogant Characters</option><option value="9851">Artifact Crafting</option><option value="4699">Artifacts</option><option value="141">Artificial Intelligence</option><option value="271">Artists</option><option value="81">Assassins</option><option value="8965">Astrologers</option><option value="9776">Autism</option><option value="11491">Automatons</option><option value="12917">Average-looking Protagonist</option><option value="1134">Award-winning Work</option><option value="13016">Awkward Protagonist</option><option value="4975">Bands</option><option value="1346">Based on a Movie</option><option value="1632">Based on a Song</option><option value="6572">Based on a TV Show</option><option value="1442">Based on a Video Game</option><option value="1447">Based on a Visual Novel</option><option value="1230">Based on an Anime</option><option value="825">Battle Academy</option><option value="6308">Battle Competition</option><option value="10668">BDSM</option><option value="5363">Beast Companions</option><option value="5406">Beastkin</option><option value="116">Beasts</option><option value="111">Beautiful Female Lead</option><option value="15625">Bestiality</option><option value="256">Betrayal</option><option value="2573">Bickering Couple</option><option value="13371">Biochip</option><option value="6359">Bisexual Protagonist</option><option value="2943">Black Belly</option><option value="916">Blackmail</option><option value="2734">Blacksmith</option><option value="9280">Blind Dates</option><option value="2347">Blind Protagonist</option><option value="9492">Blood Manipulation</option><option value="5824">Bloodlines</option><option value="2642">Body Swap</option><option value="9907">Body Tempering</option><option value="568">Body-double</option><option value="224">Bodyguards</option><option value="1288">Books</option><option value="13805">Bookworm</option><option value="325">Boss-Subordinate Relationship</option><option value="1710">Brainwashing</option><option value="5607">Breast Fetish</option><option value="9410">Broken Engagement</option><option value="604">Brother Complex</option><option value="8185">Brotherhood</option><option value="3825">Buddhism</option><option value="123">Bullying</option><option value="782">Business Management</option><option value="11786">Businessmen</option><option value="1597">Butlers</option><option value="12910">Calm Protagonist</option><option value="4248">Cannibalism</option><option value="2774">Card Games</option><option value="11345">Carefree Protagonist</option><option value="12951">Caring Protagonist</option><option value="12914">Cautious Protagonist</option><option value="2304">Celebrities</option><option value="257">Character Growth</option><option value="14493">Charismatic Protagonist</option><option value="14985">Charming Protagonist</option><option value="1942">Chat Rooms</option><option value="5708">Cheats</option><option value="1832">Chefs</option><option value="1368">Child Abuse</option><option value="3470">Child Protagonist</option><option value="1391">Childcare</option><option value="85">Childhood Friends</option><option value="917">Childhood Love</option><option value="691">Childhood Promise</option><option value="14986">Childish Protagonist</option><option value="169">Chuunibyou</option><option value="4423">Clan Building</option><option value="7777">Classic</option><option value="14587">Clever Protagonist</option><option value="2356">Clingy Lover</option><option value="2274">Clones</option><option value="611">Clubs</option><option value="14989">Clumsy Love Interests</option><option value="7951">Co-Workers</option><option value="710">Cohabitation</option><option value="14606">Cold Love Interests</option><option value="14605">Cold Protagonist</option><option value="8656">Collection of Short Stories</option><option value="7776">College/University</option><option value="1675">Coma</option><option value="124">Comedic Undertone</option><option value="282">Coming of Age</option><option value="4822">Complex Family Relationships</option><option value="7117">Conditional Power</option><option value="14673">Confident Protagonist</option><option value="1404">Confinement</option><option value="3241">Conflicting Loyalties</option><option value="859">Contracts</option><option value="2506">Cooking</option><option value="352">Corruption</option><option value="9926">Cosmic Wars</option><option value="7844">Cosplay</option><option value="592">Couple Growth</option><option value="5792">Court Official</option><option value="836">Cousins</option><option value="3752">Cowardly Protagonist</option><option value="5944">Crafting</option><option value="2969">Crime</option><option value="443">Criminals</option><option value="769">Cross-dressing</option><option value="1795">Crossover</option><option value="12956">Cruel Characters</option><option value="19556">Cryostasis</option><option value="117">Cultivation</option><option value="4596">Cunnilingus</option><option value="14593">Cunning Protagonist</option><option value="14990">Curious Protagonist</option><option value="579">Curses</option><option value="16457">Cute Children</option><option value="14987">Cute Protagonist</option><option value="5525">Cute Story</option><option value="8248">Dancers</option><option value="10343">Dao Companion</option><option value="6242">Dao Comprehension</option><option value="3824">Daoism</option><option value="4599">Dark</option><option value="2215">Dead Protagonist</option><option value="2142">Death</option><option value="1615">Death of Loved Ones</option><option value="5333">Debts</option><option value="749">Delinquents</option><option value="1362">Delusions</option><option value="5526">Demi-Humans</option><option value="2386">Demon Lord</option><option value="21268">Demonic Cultivation Technique</option><option value="86">Demons</option><option value="14541">Dense Protagonist</option><option value="12964">Depictions of Cruelty</option><option value="1946">Depression</option><option value="3135">Destiny</option><option value="1200">Detectives</option><option value="12959">Determined Protagonist</option><option value="15456">Devoted Love Interests</option><option value="2243">Different Social Status</option><option value="19241">Disabilities</option><option value="258">Discrimination</option><option value="13407">Disfigurement</option><option value="2343">Dishonest Protagonist</option><option value="7949">Distrustful Protagonist</option><option value="2019">Divination</option><option value="4706">Divine Protection</option><option value="2305">Divorce</option><option value="2876">Doctors</option><option value="16412">Dolls/Puppets</option><option value="2663">Domestic Affairs</option><option value="15026">Doting Love Interests</option><option value="15025">Doting Older Siblings</option><option value="14674">Doting Parents</option><option value="509">Dragon Riders</option><option value="897">Dragon Slayers</option><option value="72">Dragons</option><option value="1406">Dreams</option><option value="311">Drugs</option><option value="8126">Druids</option><option value="174">Dungeon Master</option><option value="175">Dungeons</option><option value="3569">Dwarfs</option><option value="1283">Dystopia</option><option value="15108">e-Sports</option><option value="7711">Early Romance</option><option value="1661">Earth Invasion</option><option value="5431">Easy Going Life</option><option value="784">Economics</option><option value="1878">Editors</option><option value="3796">Eidetic Memory</option><option value="13177">Elderly Protagonist</option><option value="9855">Elemental Magic</option><option value="165">Elves</option><option value="13305">Emotionally Weak Protagonist</option><option value="3570">Empires</option><option value="840">Enemies Become Allies</option><option value="113">Enemies Become Lovers</option><option value="263">Engagement</option><option value="2735">Engineer</option><option value="4480">Enlightenment</option><option value="1798">Episodic</option><option value="10593">Eunuch</option><option value="7778">European Ambience</option><option value="10172">Evil Gods</option><option value="5411">Evil Organizations</option><option value="12874">Evil Protagonist</option><option value="7087">Evil Religions</option><option value="2049">Evolution</option><option value="6365">Exhibitionism</option><option value="711">Exorcism</option><option value="200">Eye Powers</option><option value="525">Fairies</option><option value="134">Fallen Angels</option><option value="734">Fallen Nobility</option><option value="12009">Familial Love</option><option value="2152">Familiars</option><option value="641">Family</option><option value="2013">Family Business</option><option value="8664">Family Conflict</option><option value="1833">Famous Parents</option><option value="14555">Famous Protagonist</option><option value="5399">Fanaticism</option><option value="5691">Fanfiction</option><option value="771">Fantasy Creatures</option><option value="99">Fantasy World</option><option value="3068">Farming</option><option value="5068">Fast Cultivation</option><option value="7599">Fast Learner</option><option value="13721">Fat Protagonist</option><option value="15974">Fat to Fit</option><option value="2574">Fated Lovers</option><option value="14675">Fearless Protagonist</option><option value="3953">Fellatio</option><option value="5262">Female Master</option><option value="2879">Female Protagonist</option><option value="4178">Female to Male</option><option value="10919">Feng Shui</option><option value="9340">Firearms</option><option value="1569">First Love</option><option value="4880">First-time Intercourse</option><option value="6061">Flashbacks</option><option value="17020">Fleet Battles</option><option value="923">Folklore</option><option value="1353">Forced into a Relationship</option><option value="606">Forced Living Arrangements</option><option value="3052">Forced Marriage</option><option value="2367">Forgetful Protagonist</option><option value="7405">Former Hero</option><option value="6222">Fox Spirits</option><option value="1576">Friends Become Enemies</option><option value="1476">Friendship</option><option value="797">Fujoshi</option><option value="3468">Futanari</option><option value="2616">Futuristic Setting</option><option value="7643">Galge</option><option value="9425">Gambling</option><option value="93">Game Elements</option><option value="16272">Game Ranking System</option><option value="225">Gamers</option><option value="313">Gangs</option><option value="5382">Gate to Another World</option><option value="9295">Genderless Protagonist</option><option value="13888">Generals</option><option value="17629">Genetic Modifications</option><option value="2014">Genies</option><option value="14581">Genius Protagonist</option><option value="515">Ghosts</option><option value="2044">Gladiators</option><option value="15547">Glasses-wearing Love Interests</option><option value="15548">Glasses-wearing Protagonist</option><option value="11494">Goblins</option><option value="6683">God Protagonist</option><option value="1354">God-human Relationship</option><option value="1355">Goddesses</option><option value="73">Godly Powers</option><option value="177">Gods</option><option value="4437">Golems</option><option value="455">Gore</option><option value="1736">Grave Keepers</option><option value="807">Grinding</option><option value="2005">Guardian Relationship</option><option value="438">Guilds</option><option value="792">Gunfighters</option><option value="1753">Hackers</option><option value="9445">Half-human Protagonist</option><option value="9873">Handjob</option><option value="192">Handsome Male Lead</option><option value="14431">Hard-Working Protagonist</option><option value="14992">Harem-seeking Protagonist</option><option value="8149">Harsh Training</option><option value="2307">Hated Protagonist</option><option value="12141">Healers</option><option value="5267">Heartwarming</option><option value="242">Heaven</option><option value="4770">Heavenly Tribulation</option><option value="4720">Hell</option><option value="15545">Helpful Protagonist</option><option value="808">Herbalist</option><option value="100">Heroes</option><option value="2111">Heterochromia</option><option value="7227">Hidden Abilities</option><option value="6384">Hiding True Abilities</option><option value="10134">Hiding True Identity</option><option value="765">Hikikomori</option><option value="2741">Homunculus</option><option value="1772">Honest Protagonist</option><option value="1387">Hospital</option><option value="13502">Hot-blooded Protagonist</option><option value="12338">Human Experimentation</option><option value="1687">Human Weapon</option><option value="259">Human-Nonhuman Relationship</option><option value="17218">Humanoid Protagonist</option><option value="494">Hunters</option><option value="1711">Hypnotism</option><option value="10075">Identity Crisis</option><option value="3555">Imaginary Friend</option><option value="233">Immortals</option><option value="9033">Imperial Harem</option><option value="626">Incest</option><option value="3398">Incubus</option><option value="2319">Indecisive Protagonist</option><option value="12135">Industrialization</option><option value="506">Inferiority Complex</option><option value="1400">Inheritance</option><option value="5132">Inscriptions</option><option value="550">Insects</option><option value="4539">Interconnected Storylines</option><option value="11229">Interdimensional Travel</option><option value="14993">Introverted Protagonist</option><option value="9064">Investigations</option><option value="3978">Invisibility</option><option value="3687">Jack of All Trades</option><option value="1864">Jealousy</option><option value="6665">Jiangshi</option><option value="7086">Jobless Class</option><option value="5300">JSDF</option><option value="1250">Kidnappings</option><option value="13901">Kind Love Interests</option><option value="3869">Kingdom Building</option><option value="1904">Kingdoms</option><option value="137">Knights</option><option value="571">Kuudere</option><option value="5086">Lack of Common Sense</option><option value="5426">Language Barrier</option><option value="5920">Late Romance</option><option value="4551">Lawyers</option><option value="338">Lazy Protagonist</option><option value="6706">Leadership</option><option value="3430">Legends</option><option value="7443">Level System</option><option value="1402">Library</option><option value="2222">Limited Lifespan</option><option value="3154">Living Abroad</option><option value="1437">Living Alone</option><option value="3715">Loli</option><option value="865">Loneliness</option><option value="651">Loner Protagonist</option><option value="2891">Long Separations</option><option value="1859">Long-distance Relationship</option><option value="1696">Lost Civilizations</option><option value="7151">Lottery</option><option value="1561">Love at First Sight</option><option value="19605">Love Interest Falls in Love First</option><option value="627">Love Rivals</option><option value="607">Love Triangles</option><option value="1616">Lovers Reunited</option><option value="9161">Low-key Protagonist</option><option value="12035">Loyal Subordinates</option><option value="6526">Lucky Protagonist</option><option value="60">Magic</option><option value="6095">Magic Beasts</option><option value="3038">Magic Formations</option><option value="516">Magical Girls</option><option value="3460">Magical Space</option><option value="14900">Magical Technology</option><option value="336">Maids</option><option value="3257">Male Protagonist</option><option value="171">Male to Female</option><option value="557">Male Yandere</option><option value="2604">Management</option><option value="2781">Mangaka</option><option value="12963">Manipulative Characters</option><option value="2084">Manly Gay Couple</option><option value="642">Marriage</option><option value="706">Marriage of Convenience</option><option value="10145">Martial Spirits</option><option value="19604">Masochistic Characters</option><option value="667">Master-Disciple Relationship</option><option value="260">Master-Servant Relationship</option><option value="2501">Masturbation</option><option value="6437">Matriarchy</option><option value="279">Mature Protagonist</option><option value="9021">Medical Knowledge</option><option value="4498">Medieval</option><option value="9573">Mercenaries</option><option value="14516">Merchants</option><option value="589">Military</option><option value="6073">Mind Break</option><option value="456">Mind Control</option><option value="11503">Misandry</option><option value="1516">Mismatched Couple</option><option value="172">Misunderstandings</option><option value="105">MMORPG</option><option value="12487">Mob Protagonist</option><option value="14781">Models</option><option value="2606">Modern Day</option><option value="2666">Modern Knowledge</option><option value="3754">Money Grubber</option><option value="510">Monster Girls</option><option value="8988">Monster Society</option><option value="253">Monster Tamer</option><option value="261">Monsters</option><option value="345">Movies</option><option value="20139">Mpreg</option><option value="7163">Multiple Identities</option><option value="8266">Multiple Personalities</option><option value="14639">Multiple POV</option><option value="441">Multiple Protagonists</option><option value="3706">Multiple Realms</option><option value="5149">Multiple Reincarnated Individuals</option><option value="7802">Multiple Timelines</option><option value="16138">Multiple Transported Individuals</option><option value="885">Murders</option><option value="1231">Music</option><option value="5036">Mutated Creatures</option><option value="68">Mutations</option><option value="1141">Mute Character</option><option value="8440">Mysterious Family Background</option><option value="2224">Mysterious Illness</option><option value="4783">Mysterious Past</option><option value="3537">Mystery Solving</option><option value="8995">Mythical Beasts</option><option value="1474">Mythology</option><option value="14644">Naive Protagonist</option><option value="14994">Narcissistic Protagonist</option><option value="6204">Nationalism</option><option value="130">Near-Death Experience</option><option value="186">Necromancer</option><option value="1728">Neet</option><option value="4126">Netorare</option><option value="11561">Netorase</option><option value="3862">Netori</option><option value="6164">Nightmares</option><option value="1725">Ninjas</option><option value="265">Nobles</option><option value="17199">Non-humanoid Protagonist</option><option value="2514">Non-linear Storytelling</option><option value="1524">Nudity</option><option value="7551">Nurses</option><option value="1477">Obsessive Love</option><option value="9331">Office Romance</option><option value="16202">Older Love Interests</option><option value="14574">Omegaverse</option><option value="10980">Oneshot</option><option value="5122">Online Romance</option><option value="2060">Onmyouji</option><option value="6211">Orcs</option><option value="2645">Organized Crime</option><option value="11994">Orgy</option><option value="125">Orphans</option><option value="277">Otaku</option><option value="466">Otome Game</option><option value="726">Outcasts</option><option value="1222">Outdoor Intercourse</option><option value="875">Outer Space</option><option value="4598">Overpowered Protagonist</option><option value="145">Overprotective Siblings</option><option value="8832">Pacifist Protagonist</option><option value="6049">Paizuri</option><option value="318">Parallel Worlds</option><option value="1300">Parasites</option><option value="4741">Parent Complex</option><option value="2350">Parody</option><option value="628">Part-Time Job</option><option value="220">Past Plays a Big Role</option><option value="6457">Past Trauma</option><option value="15546">Persistent Love Interests</option><option value="488">Personality Changes</option><option value="14643">Perverted Protagonist</option><option value="2710">Pets</option><option value="5623">Pharmacist</option><option value="1799">Philosophical</option><option value="1817">Phobias</option><option value="8790">Phoenixes</option><option value="1371">Photography</option><option value="3019">Pill Based Cultivation</option><option value="9071">Pill Concocting</option><option value="1215">Pilots</option><option value="3025">Pirates</option><option value="1311">Playboys</option><option value="13369">Playful Protagonist</option><option value="7813">Poetry</option><option value="2674">Poisons</option><option value="83">Police</option><option value="14996">Polite Protagonist</option><option value="298">Politics</option><option value="11890">Polyandry</option><option value="2684">Polygamy</option><option value="12909">Poor Protagonist</option><option value="8801">Poor to Rich</option><option value="13481">Popular Love Interests</option><option value="94">Possession</option><option value="12966">Possessive Characters</option><option value="1301">Post-apocalyptic</option><option value="2551">Power Couple</option><option value="673">Power Struggle</option><option value="564">Pragmatic Protagonist</option><option value="2020">Precognition</option><option value="3330">Pregnancy</option><option value="1763">Pretend Lovers</option><option value="1124">Previous Life Talent</option><option value="3534">Priestesses</option><option value="2341">Priests</option><option value="1426">Prison</option><option value="701">Proactive Protagonist</option><option value="6302">Programmer</option><option value="13215">Prophecies</option><option value="1892">Prostitutes</option><option value="19606">Protagonist Falls in Love First</option><option value="167">Protagonist Strong from the Start</option><option value="18652">Protagonist with Multiple Bodies</option><option value="13480">Psychic Powers</option><option value="846">Psychopaths</option><option value="9950">Puppeteers</option><option value="13370">Quiet Characters</option><option value="931">Quirky Characters</option><option value="2738">R-15</option><option value="4074">R-18</option><option value="2903">Race Change</option><option value="3314">Racism</option><option value="431">Rape</option><option value="11714">Rape Victim Becomes Lover</option><option value="5574">Rebellion</option><option value="447">Reincarnated as a Monster</option><option value="9480">Reincarnated as an Object</option><option value="7297">Reincarnated in a Game World</option><option value="6304">Reincarnated in Another World</option><option value="120">Reincarnation</option><option value="15178">Religions</option><option value="179">Reluctant Protagonist</option><option value="1684">Reporters</option><option value="1835">Restaurant</option><option value="1209">Resurrection</option><option value="13303">Returning from Another World</option><option value="121">Revenge</option><option value="558">Reverse Harem</option><option value="4500">Reverse Rape</option><option value="28883">Reversible Couple</option><option value="11448">Rich to Poor</option><option value="7780">Righteous Protagonist</option><option value="614">Rivalry</option><option value="334">Romantic Subplot</option><option value="106">Roommates</option><option value="335">Royalty</option><option value="12916">Ruthless Protagonist</option><option value="13496">Sadistic Characters</option><option value="7288">Saints</option><option value="1189">Salaryman</option><option value="826">Samurai</option><option value="788">Saving the World</option><option value="24959">Schemes And Conspiracies</option><option value="2288">Schizophrenia</option><option value="843">Scientists</option><option value="426">Sculptors</option><option value="732">Sealed Power</option><option value="2571">Second Chance</option><option value="1475">Secret Crush</option><option value="214">Secret Identity</option><option value="827">Secret Organizations</option><option value="1190">Secret Relationship</option><option value="14997">Secretive Protagonist</option><option value="1623">Secrets</option><option value="7536">Sect Development</option><option value="2595">Seduction</option><option value="201">Seeing Things Other Humans Can't</option><option value="1463">Selfish Protagonist</option><option value="1372">Selfless Protagonist</option><option value="14319">Seme Protagonist</option><option value="629">Senpai-Kouhai Relationship</option><option value="13520">Sentient Objects</option><option value="13498">Sentimental Protagonist</option><option value="1328">Serial Killers</option><option value="254">Servants</option><option value="3863">Seven Deadly Sins</option><option value="13993">Seven Virtues</option><option value="4432">Sex Friends</option><option value="656">Sex Slaves</option><option value="3252">Sexual Abuse</option><option value="6245">Sexual Cultivation Technique</option><option value="13802">Shameless Protagonist</option><option value="10602">Shapeshifters</option><option value="7874">Sharing A Body</option><option value="13996">Sharp-tongued Characters</option><option value="17746">Shield User</option><option value="2280">Shikigami</option><option value="3358">Short Story</option><option value="5580">Shota</option><option value="7147">Shoujo-Ai Subplot</option><option value="7146">Shounen-Ai Subplot</option><option value="565">Showbiz</option><option value="13499">Shy Characters</option><option value="1971">Sibling Rivalry</option><option value="8186">Sibling's Care</option><option value="1411">Siblings</option><option value="65">Siblings Not Related by Blood</option><option value="13199">Sickly Characters</option><option value="8631">Sign Language</option><option value="209">Singers</option><option value="1985">Single Parent</option><option value="668">Sister Complex</option><option value="3820">Skill Assimilation</option><option value="7032">Skill Books</option><option value="6753">Skill Creation</option><option value="3936">Slave Harem</option><option value="5420">Slave Protagonist</option><option value="180">Slaves</option><option value="3014">Sleeping</option><option value="3185">Slow Growth at Start</option><option value="76">Slow Romance</option><option value="831">Smart Couple</option><option value="652">Social Outcasts</option><option value="674">Soldiers</option><option value="7828">Soul Power</option><option value="4627">Souls</option><option value="8777">Spatial Manipulation</option><option value="4842">Spear Wielder</option><option value="323">Special Abilities</option><option value="2069">Spies</option><option value="2586">Spirit Advisor</option><option value="11422">Spirit Users</option><option value="202">Spirits</option><option value="1713">Stalkers</option><option value="1373">Stockholm Syndrome</option><option value="13500">Stoic Characters</option><option value="13633">Store Owner</option><option value="1191">Straight Seme</option><option value="1594">Straight Uke</option><option value="675">Strategic Battles</option><option value="4777">Strategist</option><option value="14788">Strength-based Social Hierarchy</option><option value="12881">Strong Love Interests</option><option value="4971">Strong to Stronger</option><option value="1643">Stubborn Protagonist</option><option value="608">Student Council</option><option value="1224">Student-Teacher Relationship</option><option value="645">Succubus</option><option value="898">Sudden Strength Gain</option><option value="9574">Sudden Wealth</option><option value="1743">Suicides</option><option value="2990">Summoned Hero</option><option value="4127">Summoning Magic</option><option value="347">Survival</option><option value="348">Survival Game</option><option value="4302">Sword And Magic</option><option value="18792">Sword Wielder</option><option value="7357">System Administrator</option><option value="1749">Teachers</option><option value="1847">Teamwork</option><option value="16962">Technological Gap</option><option value="8760">Tentacles</option><option value="2225">Terminal Illness</option><option value="2196">Terrorists</option><option value="1360">Thieves</option><option value="1420">Threesome</option><option value="2970">Thriller</option><option value="886">Time Loop</option><option value="2054">Time Manipulation</option><option value="887">Time Paradox</option><option value="360">Time Skip</option><option value="92">Time Travel</option><option value="5085">Timid Protagonist</option><option value="267">Tomboyish Female Lead</option><option value="355">Torture</option><option value="14264">Toys</option><option value="148">Tragic Past</option><option value="16178">Transformation Ability</option><option value="3046">Transmigration</option><option value="4323">Transplanted Memories</option><option value="7663">Transported into a Game World</option><option value="6559">Transported Modern Structure</option><option value="15008">Transported to Another World</option><option value="1279">Trap</option><option value="5825">Tribal Society</option><option value="4856">Trickster</option><option value="795">Tsundere</option><option value="518">Twins</option><option value="10488">Twisted Personality</option><option value="12155">Ugly Protagonist</option><option value="4851">Ugly to Beautiful</option><option value="1595">Unconditional Love</option><option value="12907">Underestimated Protagonist</option><option value="3718">Unique Cultivation Technique</option><option value="13875">Unique Weapon User</option><option value="13874">Unique Weapons</option><option value="38534">Unlimited Flow</option><option value="2314">Unlucky Protagonist</option><option value="4697">Unreliable Narrator</option><option value="1268">Unrequited Love</option><option value="17588">Valkyries</option><option value="149">Vampires</option><option value="11404">Villainess Noble Girls</option><option value="109">Virtual Reality</option><option value="1634">Vocaloid</option><option value="2887">Voice Actors</option><option value="3256">Voyeurism</option><option value="1216">Waiters</option><option value="5128">War Records</option><option value="101">Wars</option><option value="12813">Weak Protagonist</option><option value="71">Weak to Strong</option><option value="13334">Wealthy Characters</option><option value="17533">Werebeasts</option><option value="1338">Wishes</option><option value="1829">Witches</option><option value="634">Wizards</option><option value="21767">World Hopping</option><option value="364">World Travel</option><option value="13198">World Tree</option><option value="548">Writers</option><option value="291">Yandere</option><option value="926">Youkai</option><option value="11677">Younger Brothers</option><option value="16201">Younger Love Interests</option><option value="11678">Younger Sisters</option><option value="350">Zombies</option></select> 
		</div> 
	</div> </div>
<div class="one-half">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			 <select multiple id="tags_exclude" class="chzn-select" data-placeholder="Exclude..." name="tagsexclude&#091;&#093;"><option value="16185">Abandoned Children</option><option value="4859">Ability Steal</option><option value="1248">Absent Parents</option><option value="11325">Abusive Characters</option><option value="4885">Academy</option><option value="475">Accelerated Growth</option><option value="13403">Acting</option><option value="4976">Adapted from Manga</option><option value="6280">Adapted from Manhua</option><option value="269">Adapted to Anime</option><option value="2999">Adapted to Drama</option><option value="928">Adapted to Drama CD</option><option value="891">Adapted to Game</option><option value="270">Adapted to Manga</option><option value="182">Adapted to Manhua</option><option value="2721">Adapted to Manhwa</option><option value="1133">Adapted to Movie</option><option value="1334">Adapted to Visual Novel</option><option value="858">Adopted Children</option><option value="603">Adopted Protagonist</option><option value="4974">Adultery</option><option value="3108">Adventurers</option><option value="3802">Affair</option><option value="357">Age Progression</option><option value="216">Age Regression</option><option value="13720">Aggressive Characters</option><option value="234">Alchemy</option><option value="621">Aliens</option><option value="598">All-Girls School</option><option value="7909">Alternate World</option><option value="625">Amnesia</option><option value="774">Amusement Park</option><option value="4594">Anal</option><option value="476">Ancient China</option><option value="2662">Ancient Times</option><option value="13263">Androgynous Characters</option><option value="892">Androids</option><option value="1327">Angels</option><option value="255">Animal Characteristics</option><option value="7323">Animal Rearing</option><option value="724">Anti-Magic</option><option value="1732">Anti-social Protagonist</option><option value="4297">Antihero Protagonist</option><option value="203">Antique Shop</option><option value="513">Apartment Life</option><option value="12913">Apathetic Protagonist</option><option value="217">Apocalypse</option><option value="721">Appearance Changes</option><option value="430">Appearance Different from Actual Age</option><option value="697">Archery</option><option value="6442">Aristocracy</option><option value="9591">Arms Dealers</option><option value="306">Army</option><option value="9522">Army Building</option><option value="2676">Arranged Marriage</option><option value="13427">Arrogant Characters</option><option value="9851">Artifact Crafting</option><option value="4699">Artifacts</option><option value="141">Artificial Intelligence</option><option value="271">Artists</option><option value="81">Assassins</option><option value="8965">Astrologers</option><option value="9776">Autism</option><option value="11491">Automatons</option><option value="12917">Average-looking Protagonist</option><option value="1134">Award-winning Work</option><option value="13016">Awkward Protagonist</option><option value="4975">Bands</option><option value="1346">Based on a Movie</option><option value="1632">Based on a Song</option><option value="6572">Based on a TV Show</option><option value="1442">Based on a Video Game</option><option value="1447">Based on a Visual Novel</option><option value="1230">Based on an Anime</option><option value="825">Battle Academy</option><option value="6308">Battle Competition</option><option value="10668">BDSM</option><option value="5363">Beast Companions</option><option value="5406">Beastkin</option><option value="116">Beasts</option><option value="111">Beautiful Female Lead</option><option value="15625">Bestiality</option><option value="256">Betrayal</option><option value="2573">Bickering Couple</option><option value="13371">Biochip</option><option value="6359">Bisexual Protagonist</option><option value="2943">Black Belly</option><option value="916">Blackmail</option><option value="2734">Blacksmith</option><option value="9280">Blind Dates</option><option value="2347">Blind Protagonist</option><option value="9492">Blood Manipulation</option><option value="5824">Bloodlines</option><option value="2642">Body Swap</option><option value="9907">Body Tempering</option><option value="568">Body-double</option><option value="224">Bodyguards</option><option value="1288">Books</option><option value="13805">Bookworm</option><option value="325">Boss-Subordinate Relationship</option><option value="1710">Brainwashing</option><option value="5607">Breast Fetish</option><option value="9410">Broken Engagement</option><option value="604">Brother Complex</option><option value="8185">Brotherhood</option><option value="3825">Buddhism</option><option value="123">Bullying</option><option value="782">Business Management</option><option value="11786">Businessmen</option><option value="1597">Butlers</option><option value="12910">Calm Protagonist</option><option value="4248">Cannibalism</option><option value="2774">Card Games</option><option value="11345">Carefree Protagonist</option><option value="12951">Caring Protagonist</option><option value="12914">Cautious Protagonist</option><option value="2304">Celebrities</option><option value="257">Character Growth</option><option value="14493">Charismatic Protagonist</option><option value="14985">Charming Protagonist</option><option value="1942">Chat Rooms</option><option value="5708">Cheats</option><option value="1832">Chefs</option><option value="1368">Child Abuse</option><option value="3470">Child Protagonist</option><option value="1391">Childcare</option><option value="85">Childhood Friends</option><option value="917">Childhood Love</option><option value="691">Childhood Promise</option><option value="14986">Childish Protagonist</option><option value="169">Chuunibyou</option><option value="4423">Clan Building</option><option value="7777">Classic</option><option value="14587">Clever Protagonist</option><option value="2356">Clingy Lover</option><option value="2274">Clones</option><option value="611">Clubs</option><option value="14989">Clumsy Love Interests</option><option value="7951">Co-Workers</option><option value="710">Cohabitation</option><option value="14606">Cold Love Interests</option><option value="14605">Cold Protagonist</option><option value="8656">Collection of Short Stories</option><option value="7776">College/University</option><option value="1675">Coma</option><option value="124">Comedic Undertone</option><option value="282">Coming of Age</option><option value="4822">Complex Family Relationships</option><option value="7117">Conditional Power</option><option value="14673">Confident Protagonist</option><option value="1404">Confinement</option><option value="3241">Conflicting Loyalties</option><option value="859">Contracts</option><option value="2506">Cooking</option><option value="352">Corruption</option><option value="9926">Cosmic Wars</option><option value="7844">Cosplay</option><option value="592">Couple Growth</option><option value="5792">Court Official</option><option value="836">Cousins</option><option value="3752">Cowardly Protagonist</option><option value="5944">Crafting</option><option value="2969">Crime</option><option value="443">Criminals</option><option value="769">Cross-dressing</option><option value="1795">Crossover</option><option value="12956">Cruel Characters</option><option value="19556">Cryostasis</option><option value="117">Cultivation</option><option value="4596">Cunnilingus</option><option value="14593">Cunning Protagonist</option><option value="14990">Curious Protagonist</option><option value="579">Curses</option><option value="16457">Cute Children</option><option value="14987">Cute Protagonist</option><option value="5525">Cute Story</option><option value="8248">Dancers</option><option value="10343">Dao Companion</option><option value="6242">Dao Comprehension</option><option value="3824">Daoism</option><option value="4599">Dark</option><option value="2215">Dead Protagonist</option><option value="2142">Death</option><option value="1615">Death of Loved Ones</option><option value="5333">Debts</option><option value="749">Delinquents</option><option value="1362">Delusions</option><option value="5526">Demi-Humans</option><option value="2386">Demon Lord</option><option value="21268">Demonic Cultivation Technique</option><option value="86">Demons</option><option value="14541">Dense Protagonist</option><option value="12964">Depictions of Cruelty</option><option value="1946">Depression</option><option value="3135">Destiny</option><option value="1200">Detectives</option><option value="12959">Determined Protagonist</option><option value="15456">Devoted Love Interests</option><option value="2243">Different Social Status</option><option value="19241">Disabilities</option><option value="258">Discrimination</option><option value="13407">Disfigurement</option><option value="2343">Dishonest Protagonist</option><option value="7949">Distrustful Protagonist</option><option value="2019">Divination</option><option value="4706">Divine Protection</option><option value="2305">Divorce</option><option value="2876">Doctors</option><option value="16412">Dolls/Puppets</option><option value="2663">Domestic Affairs</option><option value="15026">Doting Love Interests</option><option value="15025">Doting Older Siblings</option><option value="14674">Doting Parents</option><option value="509">Dragon Riders</option><option value="897">Dragon Slayers</option><option value="72">Dragons</option><option value="1406">Dreams</option><option value="311">Drugs</option><option value="8126">Druids</option><option value="174">Dungeon Master</option><option value="175">Dungeons</option><option value="3569">Dwarfs</option><option value="1283">Dystopia</option><option value="15108">e-Sports</option><option value="7711">Early Romance</option><option value="1661">Earth Invasion</option><option value="5431">Easy Going Life</option><option value="784">Economics</option><option value="1878">Editors</option><option value="3796">Eidetic Memory</option><option value="13177">Elderly Protagonist</option><option value="9855">Elemental Magic</option><option value="165">Elves</option><option value="13305">Emotionally Weak Protagonist</option><option value="3570">Empires</option><option value="840">Enemies Become Allies</option><option value="113">Enemies Become Lovers</option><option value="263">Engagement</option><option value="2735">Engineer</option><option value="4480">Enlightenment</option><option value="1798">Episodic</option><option value="10593">Eunuch</option><option value="7778">European Ambience</option><option value="10172">Evil Gods</option><option value="5411">Evil Organizations</option><option value="12874">Evil Protagonist</option><option value="7087">Evil Religions</option><option value="2049">Evolution</option><option value="6365">Exhibitionism</option><option value="711">Exorcism</option><option value="200">Eye Powers</option><option value="525">Fairies</option><option value="134">Fallen Angels</option><option value="734">Fallen Nobility</option><option value="12009">Familial Love</option><option value="2152">Familiars</option><option value="641">Family</option><option value="2013">Family Business</option><option value="8664">Family Conflict</option><option value="1833">Famous Parents</option><option value="14555">Famous Protagonist</option><option value="5399">Fanaticism</option><option value="5691">Fanfiction</option><option value="771">Fantasy Creatures</option><option value="99">Fantasy World</option><option value="3068">Farming</option><option value="5068">Fast Cultivation</option><option value="7599">Fast Learner</option><option value="13721">Fat Protagonist</option><option value="15974">Fat to Fit</option><option value="2574">Fated Lovers</option><option value="14675">Fearless Protagonist</option><option value="3953">Fellatio</option><option value="5262">Female Master</option><option value="2879">Female Protagonist</option><option value="4178">Female to Male</option><option value="10919">Feng Shui</option><option value="9340">Firearms</option><option value="1569">First Love</option><option value="4880">First-time Intercourse</option><option value="6061">Flashbacks</option><option value="17020">Fleet Battles</option><option value="923">Folklore</option><option value="1353">Forced into a Relationship</option><option value="606">Forced Living Arrangements</option><option value="3052">Forced Marriage</option><option value="2367">Forgetful Protagonist</option><option value="7405">Former Hero</option><option value="6222">Fox Spirits</option><option value="1576">Friends Become Enemies</option><option value="1476">Friendship</option><option value="797">Fujoshi</option><option value="3468">Futanari</option><option value="2616">Futuristic Setting</option><option value="7643">Galge</option><option value="9425">Gambling</option><option value="93">Game Elements</option><option value="16272">Game Ranking System</option><option value="225">Gamers</option><option value="313">Gangs</option><option value="5382">Gate to Another World</option><option value="9295">Genderless Protagonist</option><option value="13888">Generals</option><option value="17629">Genetic Modifications</option><option value="2014">Genies</option><option value="14581">Genius Protagonist</option><option value="515">Ghosts</option><option value="2044">Gladiators</option><option value="15547">Glasses-wearing Love Interests</option><option value="15548">Glasses-wearing Protagonist</option><option value="11494">Goblins</option><option value="6683">God Protagonist</option><option value="1354">God-human Relationship</option><option value="1355">Goddesses</option><option value="73">Godly Powers</option><option value="177">Gods</option><option value="4437">Golems</option><option value="455">Gore</option><option value="1736">Grave Keepers</option><option value="807">Grinding</option><option value="2005">Guardian Relationship</option><option value="438">Guilds</option><option value="792">Gunfighters</option><option value="1753">Hackers</option><option value="9445">Half-human Protagonist</option><option value="9873">Handjob</option><option value="192">Handsome Male Lead</option><option value="14431">Hard-Working Protagonist</option><option value="14992">Harem-seeking Protagonist</option><option value="8149">Harsh Training</option><option value="2307">Hated Protagonist</option><option value="12141">Healers</option><option value="5267">Heartwarming</option><option value="242">Heaven</option><option value="4770">Heavenly Tribulation</option><option value="4720">Hell</option><option value="15545">Helpful Protagonist</option><option value="808">Herbalist</option><option value="100">Heroes</option><option value="2111">Heterochromia</option><option value="7227">Hidden Abilities</option><option value="6384">Hiding True Abilities</option><option value="10134">Hiding True Identity</option><option value="765">Hikikomori</option><option value="2741">Homunculus</option><option value="1772">Honest Protagonist</option><option value="1387">Hospital</option><option value="13502">Hot-blooded Protagonist</option><option value="12338">Human Experimentation</option><option value="1687">Human Weapon</option><option value="259">Human-Nonhuman Relationship</option><option value="17218">Humanoid Protagonist</option><option value="494">Hunters</option><option value="1711">Hypnotism</option><option value="10075">Identity Crisis</option><option value="3555">Imaginary Friend</option><option value="233">Immortals</option><option value="9033">Imperial Harem</option><option value="626">Incest</option><option value="3398">Incubus</option><option value="2319">Indecisive Protagonist</option><option value="12135">Industrialization</option><option value="506">Inferiority Complex</option><option value="1400">Inheritance</option><option value="5132">Inscriptions</option><option value="550">Insects</option><option value="4539">Interconnected Storylines</option><option value="11229">Interdimensional Travel</option><option value="14993">Introverted Protagonist</option><option value="9064">Investigations</option><option value="3978">Invisibility</option><option value="3687">Jack of All Trades</option><option value="1864">Jealousy</option><option value="6665">Jiangshi</option><option value="7086">Jobless Class</option><option value="5300">JSDF</option><option value="1250">Kidnappings</option><option value="13901">Kind Love Interests</option><option value="3869">Kingdom Building</option><option value="1904">Kingdoms</option><option value="137">Knights</option><option value="571">Kuudere</option><option value="5086">Lack of Common Sense</option><option value="5426">Language Barrier</option><option value="5920">Late Romance</option><option value="4551">Lawyers</option><option value="338">Lazy Protagonist</option><option value="6706">Leadership</option><option value="3430">Legends</option><option value="7443">Level System</option><option value="1402">Library</option><option value="2222">Limited Lifespan</option><option value="3154">Living Abroad</option><option value="1437">Living Alone</option><option value="3715">Loli</option><option value="865">Loneliness</option><option value="651">Loner Protagonist</option><option value="2891">Long Separations</option><option value="1859">Long-distance Relationship</option><option value="1696">Lost Civilizations</option><option value="7151">Lottery</option><option value="1561">Love at First Sight</option><option value="19605">Love Interest Falls in Love First</option><option value="627">Love Rivals</option><option value="607">Love Triangles</option><option value="1616">Lovers Reunited</option><option value="9161">Low-key Protagonist</option><option value="12035">Loyal Subordinates</option><option value="6526">Lucky Protagonist</option><option value="60">Magic</option><option value="6095">Magic Beasts</option><option value="3038">Magic Formations</option><option value="516">Magical Girls</option><option value="3460">Magical Space</option><option value="14900">Magical Technology</option><option value="336">Maids</option><option value="3257">Male Protagonist</option><option value="171">Male to Female</option><option value="557">Male Yandere</option><option value="2604">Management</option><option value="2781">Mangaka</option><option value="12963">Manipulative Characters</option><option value="2084">Manly Gay Couple</option><option value="642">Marriage</option><option value="706">Marriage of Convenience</option><option value="10145">Martial Spirits</option><option value="19604">Masochistic Characters</option><option value="667">Master-Disciple Relationship</option><option value="260">Master-Servant Relationship</option><option value="2501">Masturbation</option><option value="6437">Matriarchy</option><option value="279">Mature Protagonist</option><option value="9021">Medical Knowledge</option><option value="4498">Medieval</option><option value="9573">Mercenaries</option><option value="14516">Merchants</option><option value="589">Military</option><option value="6073">Mind Break</option><option value="456">Mind Control</option><option value="11503">Misandry</option><option value="1516">Mismatched Couple</option><option value="172">Misunderstandings</option><option value="105">MMORPG</option><option value="12487">Mob Protagonist</option><option value="14781">Models</option><option value="2606">Modern Day</option><option value="2666">Modern Knowledge</option><option value="3754">Money Grubber</option><option value="510">Monster Girls</option><option value="8988">Monster Society</option><option value="253">Monster Tamer</option><option value="261">Monsters</option><option value="345">Movies</option><option value="20139">Mpreg</option><option value="7163">Multiple Identities</option><option value="8266">Multiple Personalities</option><option value="14639">Multiple POV</option><option value="441">Multiple Protagonists</option><option value="3706">Multiple Realms</option><option value="5149">Multiple Reincarnated Individuals</option><option value="7802">Multiple Timelines</option><option value="16138">Multiple Transported Individuals</option><option value="885">Murders</option><option value="1231">Music</option><option value="5036">Mutated Creatures</option><option value="68">Mutations</option><option value="1141">Mute Character</option><option value="8440">Mysterious Family Background</option><option value="2224">Mysterious Illness</option><option value="4783">Mysterious Past</option><option value="3537">Mystery Solving</option><option value="8995">Mythical Beasts</option><option value="1474">Mythology</option><option value="14644">Naive Protagonist</option><option value="14994">Narcissistic Protagonist</option><option value="6204">Nationalism</option><option value="130">Near-Death Experience</option><option value="186">Necromancer</option><option value="1728">Neet</option><option value="4126">Netorare</option><option value="11561">Netorase</option><option value="3862">Netori</option><option value="6164">Nightmares</option><option value="1725">Ninjas</option><option value="265">Nobles</option><option value="17199">Non-humanoid Protagonist</option><option value="2514">Non-linear Storytelling</option><option value="1524">Nudity</option><option value="7551">Nurses</option><option value="1477">Obsessive Love</option><option value="9331">Office Romance</option><option value="16202">Older Love Interests</option><option value="14574">Omegaverse</option><option value="10980">Oneshot</option><option value="5122">Online Romance</option><option value="2060">Onmyouji</option><option value="6211">Orcs</option><option value="2645">Organized Crime</option><option value="11994">Orgy</option><option value="125">Orphans</option><option value="277">Otaku</option><option value="466">Otome Game</option><option value="726">Outcasts</option><option value="1222">Outdoor Intercourse</option><option value="875">Outer Space</option><option value="4598">Overpowered Protagonist</option><option value="145">Overprotective Siblings</option><option value="8832">Pacifist Protagonist</option><option value="6049">Paizuri</option><option value="318">Parallel Worlds</option><option value="1300">Parasites</option><option value="4741">Parent Complex</option><option value="2350">Parody</option><option value="628">Part-Time Job</option><option value="220">Past Plays a Big Role</option><option value="6457">Past Trauma</option><option value="15546">Persistent Love Interests</option><option value="488">Personality Changes</option><option value="14643">Perverted Protagonist</option><option value="2710">Pets</option><option value="5623">Pharmacist</option><option value="1799">Philosophical</option><option value="1817">Phobias</option><option value="8790">Phoenixes</option><option value="1371">Photography</option><option value="3019">Pill Based Cultivation</option><option value="9071">Pill Concocting</option><option value="1215">Pilots</option><option value="3025">Pirates</option><option value="1311">Playboys</option><option value="13369">Playful Protagonist</option><option value="7813">Poetry</option><option value="2674">Poisons</option><option value="83">Police</option><option value="14996">Polite Protagonist</option><option value="298">Politics</option><option value="11890">Polyandry</option><option value="2684">Polygamy</option><option value="12909">Poor Protagonist</option><option value="8801">Poor to Rich</option><option value="13481">Popular Love Interests</option><option value="94">Possession</option><option value="12966">Possessive Characters</option><option value="1301">Post-apocalyptic</option><option value="2551">Power Couple</option><option value="673">Power Struggle</option><option value="564">Pragmatic Protagonist</option><option value="2020">Precognition</option><option value="3330">Pregnancy</option><option value="1763">Pretend Lovers</option><option value="1124">Previous Life Talent</option><option value="3534">Priestesses</option><option value="2341">Priests</option><option value="1426">Prison</option><option value="701">Proactive Protagonist</option><option value="6302">Programmer</option><option value="13215">Prophecies</option><option value="1892">Prostitutes</option><option value="19606">Protagonist Falls in Love First</option><option value="167">Protagonist Strong from the Start</option><option value="18652">Protagonist with Multiple Bodies</option><option value="13480">Psychic Powers</option><option value="846">Psychopaths</option><option value="9950">Puppeteers</option><option value="13370">Quiet Characters</option><option value="931">Quirky Characters</option><option value="2738">R-15</option><option value="4074">R-18</option><option value="2903">Race Change</option><option value="3314">Racism</option><option value="431">Rape</option><option value="11714">Rape Victim Becomes Lover</option><option value="5574">Rebellion</option><option value="447">Reincarnated as a Monster</option><option value="9480">Reincarnated as an Object</option><option value="7297">Reincarnated in a Game World</option><option value="6304">Reincarnated in Another World</option><option value="120">Reincarnation</option><option value="15178">Religions</option><option value="179">Reluctant Protagonist</option><option value="1684">Reporters</option><option value="1835">Restaurant</option><option value="1209">Resurrection</option><option value="13303">Returning from Another World</option><option value="121">Revenge</option><option value="558">Reverse Harem</option><option value="4500">Reverse Rape</option><option value="28883">Reversible Couple</option><option value="11448">Rich to Poor</option><option value="7780">Righteous Protagonist</option><option value="614">Rivalry</option><option value="334">Romantic Subplot</option><option value="106">Roommates</option><option value="335">Royalty</option><option value="12916">Ruthless Protagonist</option><option value="13496">Sadistic Characters</option><option value="7288">Saints</option><option value="1189">Salaryman</option><option value="826">Samurai</option><option value="788">Saving the World</option><option value="24959">Schemes And Conspiracies</option><option value="2288">Schizophrenia</option><option value="843">Scientists</option><option value="426">Sculptors</option><option value="732">Sealed Power</option><option value="2571">Second Chance</option><option value="1475">Secret Crush</option><option value="214">Secret Identity</option><option value="827">Secret Organizations</option><option value="1190">Secret Relationship</option><option value="14997">Secretive Protagonist</option><option value="1623">Secrets</option><option value="7536">Sect Development</option><option value="2595">Seduction</option><option value="201">Seeing Things Other Humans Can't</option><option value="1463">Selfish Protagonist</option><option value="1372">Selfless Protagonist</option><option value="14319">Seme Protagonist</option><option value="629">Senpai-Kouhai Relationship</option><option value="13520">Sentient Objects</option><option value="13498">Sentimental Protagonist</option><option value="1328">Serial Killers</option><option value="254">Servants</option><option value="3863">Seven Deadly Sins</option><option value="13993">Seven Virtues</option><option value="4432">Sex Friends</option><option value="656">Sex Slaves</option><option value="3252">Sexual Abuse</option><option value="6245">Sexual Cultivation Technique</option><option value="13802">Shameless Protagonist</option><option value="10602">Shapeshifters</option><option value="7874">Sharing A Body</option><option value="13996">Sharp-tongued Characters</option><option value="17746">Shield User</option><option value="2280">Shikigami</option><option value="3358">Short Story</option><option value="5580">Shota</option><option value="7147">Shoujo-Ai Subplot</option><option value="7146">Shounen-Ai Subplot</option><option value="565">Showbiz</option><option value="13499">Shy Characters</option><option value="1971">Sibling Rivalry</option><option value="8186">Sibling's Care</option><option value="1411">Siblings</option><option value="65">Siblings Not Related by Blood</option><option value="13199">Sickly Characters</option><option value="8631">Sign Language</option><option value="209">Singers</option><option value="1985">Single Parent</option><option value="668">Sister Complex</option><option value="3820">Skill Assimilation</option><option value="7032">Skill Books</option><option value="6753">Skill Creation</option><option value="3936">Slave Harem</option><option value="5420">Slave Protagonist</option><option value="180">Slaves</option><option value="3014">Sleeping</option><option value="3185">Slow Growth at Start</option><option value="76">Slow Romance</option><option value="831">Smart Couple</option><option value="652">Social Outcasts</option><option value="674">Soldiers</option><option value="7828">Soul Power</option><option value="4627">Souls</option><option value="8777">Spatial Manipulation</option><option value="4842">Spear Wielder</option><option value="323">Special Abilities</option><option value="2069">Spies</option><option value="2586">Spirit Advisor</option><option value="11422">Spirit Users</option><option value="202">Spirits</option><option value="1713">Stalkers</option><option value="1373">Stockholm Syndrome</option><option value="13500">Stoic Characters</option><option value="13633">Store Owner</option><option value="1191">Straight Seme</option><option value="1594">Straight Uke</option><option value="675">Strategic Battles</option><option value="4777">Strategist</option><option value="14788">Strength-based Social Hierarchy</option><option value="12881">Strong Love Interests</option><option value="4971">Strong to Stronger</option><option value="1643">Stubborn Protagonist</option><option value="608">Student Council</option><option value="1224">Student-Teacher Relationship</option><option value="645">Succubus</option><option value="898">Sudden Strength Gain</option><option value="9574">Sudden Wealth</option><option value="1743">Suicides</option><option value="2990">Summoned Hero</option><option value="4127">Summoning Magic</option><option value="347">Survival</option><option value="348">Survival Game</option><option value="4302">Sword And Magic</option><option value="18792">Sword Wielder</option><option value="7357">System Administrator</option><option value="1749">Teachers</option><option value="1847">Teamwork</option><option value="16962">Technological Gap</option><option value="8760">Tentacles</option><option value="2225">Terminal Illness</option><option value="2196">Terrorists</option><option value="1360">Thieves</option><option value="1420">Threesome</option><option value="2970">Thriller</option><option value="886">Time Loop</option><option value="2054">Time Manipulation</option><option value="887">Time Paradox</option><option value="360">Time Skip</option><option value="92">Time Travel</option><option value="5085">Timid Protagonist</option><option value="267">Tomboyish Female Lead</option><option value="355">Torture</option><option value="14264">Toys</option><option value="148">Tragic Past</option><option value="16178">Transformation Ability</option><option value="3046">Transmigration</option><option value="4323">Transplanted Memories</option><option value="7663">Transported into a Game World</option><option value="6559">Transported Modern Structure</option><option value="15008">Transported to Another World</option><option value="1279">Trap</option><option value="5825">Tribal Society</option><option value="4856">Trickster</option><option value="795">Tsundere</option><option value="518">Twins</option><option value="10488">Twisted Personality</option><option value="12155">Ugly Protagonist</option><option value="4851">Ugly to Beautiful</option><option value="1595">Unconditional Love</option><option value="12907">Underestimated Protagonist</option><option value="3718">Unique Cultivation Technique</option><option value="13875">Unique Weapon User</option><option value="13874">Unique Weapons</option><option value="38534">Unlimited Flow</option><option value="2314">Unlucky Protagonist</option><option value="4697">Unreliable Narrator</option><option value="1268">Unrequited Love</option><option value="17588">Valkyries</option><option value="149">Vampires</option><option value="11404">Villainess Noble Girls</option><option value="109">Virtual Reality</option><option value="1634">Vocaloid</option><option value="2887">Voice Actors</option><option value="3256">Voyeurism</option><option value="1216">Waiters</option><option value="5128">War Records</option><option value="101">Wars</option><option value="12813">Weak Protagonist</option><option value="71">Weak to Strong</option><option value="13334">Wealthy Characters</option><option value="17533">Werebeasts</option><option value="1338">Wishes</option><option value="1829">Witches</option><option value="634">Wizards</option><option value="21767">World Hopping</option><option value="364">World Travel</option><option value="13198">World Tree</option><option value="548">Writers</option><option value="291">Yandere</option><option value="926">Youkai</option><option value="11677">Younger Brothers</option><option value="16201">Younger Love Interests</option><option value="11678">Younger Sisters</option><option value="350">Zombies</option></select> 
		</div> 
	</div> </div>
</div>      
<div class="extopsf"></div>   
<span class="sftitle sfpad new">Reading Lists <select name="releases_mm" id="rli_mm" class="select sf" style="width:70px;">
						<option value="exclude" selected>Exclude</option>
						<option value="include">Include</option>
					</select>]<i onclick="sf_getnovelinfo('readinglists', 'Reading Lists');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 

<div class="extrasf"></div>     

You must be logged in to use this option.  
<div class="extopsf"></div>      

 
<span class="sftitle sfpad new">Story Status</span> 
<br />

<div>
 <select name="storystatus" id="storystatus" class="storystatus" onchange="">
  <option value="1" selected="">All</option>
  <option value="2">Completed</option>
  <option value="3">Ongoing</option>
  <option value="4">Hiatus</option>
  </select>
</div>
 <div class="extopsf"></div>
    
<span class="sftitle sfpad new">Sort Results By... <i onclick="sf_getnovelinfo('sorting', 'Sort Results By...');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 

<div class="extrasf"></div>
<div class="g-cols wpb_row offset_default">
<div class="one-half">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			
 <select name="sortmyresults" id="sortresults" class="sortresults" style="width: 40%;" onchange="">
  <option value="srel">Chapters</option>
  <option value="sfrel">Frequency</option>
  <option value="srank">Rank</option>
  <option value="srate">Rating</option>
  <option value="sread">Readers</option>
  <option value="sreview">Reviews</option>
  <option value="abc">Title</option>
  <option value="sdate" selected="">Last Updated</option>
  </select> 
 
		</div> 
	</div> </div>
<div class="one-half">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			 
<select name="sortmyorder" id="sortorder" class="sortorder" style="width: 40%;" onchange="">
  <option value="asc">Ascending</option>
  <option value="desc" selected="">Descending</option>
  </select> 
 
		</div> 
	</div> </div>
</div>  
<div class="extopsf"></div>
<a class="crfr finder" href="//www.novelupdates.com/series-finder/">Clear Filters</a>
</div>  <input type="submit" id="rk_filter" class="btnrev review" style="float:none; margin-top: 15px;" onclick="applyfilter('//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc');" value="Apply Filters"> 
</div>
</form>


   	                 
                
            <div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_46672.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (3.4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid46672" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/another-world-awakening-transcendental-create-skill/">Another World Awakening Transcendental Create Skill</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 100 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 1.1 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 1039 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 2 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a></div>A bus accident during a school trip sends the entire class to another world. After being transported to another world, they are attacked by monsters and bandits, and while their classmates awaken sword and magic skills one after another, Takaya Nagami’s abilities never awaken, and he ends up being kicked out of the class party. Finding himself all alone in a world where he can’t rely on anyone, Takaya despairs and impulsively tries to commit suicide, but he survived and is rescued by local adventurers while in a miserable state.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
In the course of his adventures with these adventurers, he gradually gains strength, confidence, and trust (as well as cute girls!) by using his production and processing skills. An exhilarating and thrilling tale of growth in a different world that follows the dullest boy in his class as he becomes an unstoppable adventurer!
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_30279.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (4.1)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid30279" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/ill-become-a-villainess-that-will-go-down-in-history/">I&#8217;ll Become a Villainess That Will Go Down in History</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 225 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 1.6 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 7819 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 49 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="9" href="https://www.novelupdates.com/genre/drama/" title="View All Drama Related Series" >Drama</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="157" href="https://www.novelupdates.com/genre/shoujo/" title="View All Shoujo Related Series" >Shoujo</a></div>I’ve always dreamed of becoming a villainess, but I never thought that I would actually become one….!<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
This is the story of a young girl who aspires to become a villainous noble girl who’s capable of growing stronger through every confrontation that she is faced with.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_26659.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.1)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid26659" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/i-help-the-richest-man-spend-money-to-prevent-disasters/">I Help the Richest Man Spend Money to Prevent Disasters</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 372 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 1.7 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 5617 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 83 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a></div>Ye Zhi not only inherited the run-down house, but also the engagement with a man. The man who met his true love disliked his fiancée in all ways and asked for divorce.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
In order to express their guilt, the man’s family decided to make up for it with Ye Zhi. But she did not want their ‘compensation’ and simply left without turning her back.<p style="margin-top:-5px;"></p>
The next second, Ye Zhi was picked up by a Rolls Royce.<p style="margin-top:-5px;"></p>
Gu Ren, the son of the richest man, had a tough life and couldn’t live past 30 years old. Only Ye Zhi, who was born on the cloudy day of the lunar year, was the most suitable match that could make him survive.<p style="margin-top:-5px;"></p>
Ye Zhi and Gu Ren got married. Gu Ren gave her a bank card. All she has to do is—<p style="margin-top:-5px;"></p>
Spend money for him! Squander his fortune!<p style="margin-top:-5px;"></p>
Ye Zhi bought a Lamborghini for Gu Ren, but he had an accident as soon as he drove the car.<p style="margin-top:-5px;"></p>
As soon as she bought some Hermes platinum bags for herself, Gu Ren signed a more than one billion worth contract.<p style="margin-top:-5px;"></p>
It turned out that the money could only be spent on herself.  The more she spent, the more prosperous she became.<p style="margin-top:-5px;"></p>
Since then her life has changed completely. She became a rich woman who owns luxury items, limited edition bags and a villa overlooking the Forbidden City.<p style="margin-top:-5px;"></p>
The man thought that Ye Zhi, who he had broken up with, must have suffered a lot. Until one day, he saw his former fiancée walking in the street, followed by a striking Lamborghini.<p style="margin-top:-5px;"></p>
Wait! Who is the man next to her?<p style="margin-top:-5px;"></p>
This man is a super-rich man—Hundreds of times richer than his own family!
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_45277.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (3.8)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid45277" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/i-just-want-to-be-a-salted-fish-quietly/">I Just Want To Be a Salted Fish Quietly</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 71 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 10.2 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 649 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 5 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="9" href="https://www.novelupdates.com/genre/drama/" title="View All Drama Related Series" >Drama</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="245" href="https://www.novelupdates.com/genre/mystery/" title="View All Mystery Related Series" >Mystery</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="6" href="https://www.novelupdates.com/genre/school-life/" title="View All School Life Related Series" >School Life</a> <a class="gennew search" gid="7" href="https://www.novelupdates.com/genre/slice-of-life/" title="View All Slice of Life Related Series" >Slice of Life</a> <a class="gennew search" gid="16" href="https://www.novelupdates.com/genre/supernatural/" title="View All Supernatural Related Series" >Supernatural</a></div>Lin Xian said bitterly: I really just want to be a salted fish quietly.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
System: No, you can’t!<p style="margin-top:-5px;"></p>
System: Congratulations to the host for obtaining the VIP privilege gift package and gaining the skill, 【Driving a Ferrari with one hand】…<p style="margin-top:-5px;"></p>
Congratulations to the host for obtaining props, 【Absolute focus】…<p style="margin-top:-5px;"></p>
Congratulations to the host for obtaining props, 【Consumption Rebate Critical Card】…<p style="margin-top:-5px;"></p>
…<p style="margin-top:-5px;"></p>
With the help of the system, Lin Xian found that he seemed to be on the path to becoming an Almighty Male God, getting farther and farther…<p style="margin-top:-5px;"></p>
(Salted Fish – a term used to refer to a person who has no intention of doing anything or person who wants to live a carefree life)<p style="margin-top:-5px;"></p>
(Male God – Prince Charming / Mister Right)
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_39787.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.1)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid39787" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/tai-sui/">Tai Sui</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 251 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 1.1 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 907 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 5 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="330" href="https://www.novelupdates.com/genre/historical/" title="View All Historical Related Series" >Historical</a> <a class="gennew search" gid="480" href="https://www.novelupdates.com/genre/xianxia/" title="View All Xianxia Related Series" >Xianxia</a></div>If I had a choice, I would only want to be a little insect in the mundane dust, born in confusion, dying in mediocrity, never seeing the light of day beneath the fog of Jinping City.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
Better than taking this wrong road to heaven.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_27879.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (3.7)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid27879" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/hellmode-a-hardcore-gamer-becomes-peerless-in-another-world-with-retro-game-settings/">Hellmode ~A Hardcore Gamer Becomes Peerless in Another World with Retro Game Settings~</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 264 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 1.3 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 4730 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 23 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="6" href="https://www.novelupdates.com/genre/school-life/" title="View All School Life Related Series" >School Life</a> <a class="gennew search" gid="16" href="https://www.novelupdates.com/genre/supernatural/" title="View All Supernatural Related Series" >Supernatural</a></div>Kenichi Yamada was a 35-year-old salaryman. As a hardcore game enthusiast, he was saddened by the modern trend towards casual games.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
So, when a site claimed to be “for people who like to do things the hard way,” he just couldn’t resist. Thus, he was reincarnated into another world as Allen, playing on Hell Mode.<p style="margin-top:-5px;"></p>
This is fantasy light novel about Allen’s journey as a summoner.<p style="margin-top:-5px;"></p>
Reincarnated as a serf, he starts from nothing.<p style="margin-top:-5px;"></p>
His journey begins with absolutely no knowledge of the world around him, just like those games he played 10, 20 years ago.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_33838.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.2)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid33838" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/little-tyrant-doesnt-want-to-meet-with-a-bad-end/">Little Tyrant Doesn&#8217;t Want to Meet with a Bad End</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 562 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 1.7 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 10072 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 244 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a></div>The moment the little tyrant of the nobles, Roel Ascart, saw his stepsister, he recalled his memories. He realized that he was in the world of a gal game he played in his previous life. To make things worse, he was the greatest villain in the common route of the game!<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
“I’ll be killed by the main character and the four capture targets ten years from now. Is there still any hope for me?”<p style="margin-top:-5px;"></p>
Just thinking about the fearsome glints of those sharp swords those beautiful capture targets held in their hands, Roel couldn’t help but tremble in fear.<p style="margin-top:-5px;"></p>
Till a voice finally sounded in his head.<p style="margin-top:-5px;"></p>
【Welcome to the House Resurgence System】
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_22254.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid22254" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/star-odyssey/">Star Odyssey</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 1264 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.5 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 2122 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 17 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="9" href="https://www.novelupdates.com/genre/drama/" title="View All Drama Related Series" >Drama</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="11" href="https://www.novelupdates.com/genre/sci-fi/" title="View All Sci-fi Related Series" >Sci-fi</a> <a class="gennew search" gid="12" href="https://www.novelupdates.com/genre/shounen/" title="View All Shounen Related Series" >Shounen</a> <a class="gennew search" gid="3954" href="https://www.novelupdates.com/genre/xuanhuan/" title="View All Xuanhuan Related Series" >Xuanhuan</a></div>Join Lu Yin on an epic journey across the Universe, pursuing the truth and tragedy of his past. This is a world of science fantasy<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"> where the older generations step back and allow the young to take charge of affairs. Heart-wrenching separations, terrifying situations, all with comic relief that will leave you coming back for more. This is a world where the other characters actually matter, and are revisited frequently as their own lives unfold. Dotting Lu Yin’s path are monumental feats of kingdom-building and treacherous political situations where he must tread carefully if he wants to get to the truth of his history.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_52799.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgkr" style="font-weight: bold;">KR</span> (3.9)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid52799" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/how-to-live-as-the-vampire-lord/">How to Live As the Vampire Lord</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 67 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.9 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 619 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 3 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a></div>Vampire Eugene -- a sacrificial lamb slaughtered after half a year of running to fulfil the desires of a templar for fame.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
He was given a second chance at life after concluding his life with regrets.<p style="margin-top:-5px;"></p>
"I will never again die the same way. If I really did return to the past, no matter what it takes... I will never regret again."
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_52696.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.3)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid52696" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/the-reincarnated-villain-makes-the-heroines-tearfully-beg-for-forgiveness/">The Reincarnated Villain Makes The Heroines Tearfully Beg for Forgiveness</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 10 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 3.9 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 441 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 6 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="280" href="https://www.novelupdates.com/genre/adult/" title="View All Adult Related Series" >Adult</a> <a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="4" href="https://www.novelupdates.com/genre/mature/" title="View All Mature Related Series" >Mature</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="3954" href="https://www.novelupdates.com/genre/xuanhuan/" title="View All Xuanhuan Related Series" >Xuanhuan</a></div>Lin Yan had transmigrated into the world of Immortals and had become a Heavenly Emperor who was already defeated by the Male Lead along with the ten Female Leads at the beginning of the plot.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
But as Lin Yan was about to die, he awakened a power that he could use to rewrite his past doings.<p style="margin-top:-5px;"></p>
So, Lin Yan has begun to rewrite his past.<p style="margin-top:-5px;"></p>
When Lin Yan was still a child, he struggled to feed Xiao Yanran, one of the female leads, who was then younger than him. He had to strip his Supreme Bone and give it to her, protect her from the obstacles. He made a vow, “As long as I am here, no one will be allowed to bully you.”<p style="margin-top:-5px;"></p>
Later when Lin Yan’s cultivation rose , he had run into Xiao Mei, yet another heroine, and the two of them spent time together and eventually got married. When the Devils were about to invade their world, Lin Yan said staring right into his wife’s eyes, “I will protect you with my life…”<p style="margin-top:-5px;"></p>
And soon he rose to become the Heavenly Emperor.<p style="margin-top:-5px;"></p>
Everything he did was to shelter his loved ones and to protect the world.<p style="margin-top:-5px;"></p>
However, his actions were misunderstood by the world.<p style="margin-top:-5px;"></p>
…<p style="margin-top:-5px;"></p>
Let’s see what all Lin Yan had done in the past…
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_23882.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (3.3)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid23882" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/even-if-im-reborn-as-a-cute-dragon-girl-i-will-still-make-a-harem/">Even If I&#8217;m Reborn as a Cute Dragon Girl, I Will Still Make a Harem</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 296 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 3.3 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 3234 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 28 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="168" href="https://www.novelupdates.com/genre/gender-bender/" title="View All Gender Bender Related Series" >Gender Bender</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="6" href="https://www.novelupdates.com/genre/school-life/" title="View All School Life Related Series" >School Life</a> <a class="gennew search" gid="851" href="https://www.novelupdates.com/genre/shoujo-ai/" title="View All Shoujo Ai Related Series" >Shoujo Ai</a></div>We, Aristides Castro G Morris Brooklyn Washington Napoleon George I, are the lord of this empire. Our demise art the result of saving our people from impending disaster. In order to award us for our glorious achievement, the goddess hast granted us the chance at rebirth. However… why did we turn into a girl, and the princess of dragons at that?!<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
This is the story of a man who suffered from Eight-Grade Syndrome (chuunibyou) and has been reincarnated as the dragons’ princess. However, because of his chuunibyou behavior, one of the goddesses was annoyed at him and cast a curse on him. One that summoned lightning upon him whenever his symptom acts up.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_49371.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (3)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid49371" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/i-reincarnated-but-will-try-to-live-without-using-my-cheat-ability/">I Reincarnated But Will Try To Live Without Using My Cheat Ability</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 294 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.6 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 939 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 8 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="12" href="https://www.novelupdates.com/genre/shounen/" title="View All Shounen Related Series" >Shounen</a></div>When someone says, “I’ll give you a cheat ability, so just use it”, don’t you feel like not using it?<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
There is nothing more expensive than free stuff.<p style="margin-top:-5px;"></p>
It’s not cool to get excited about something you got for free.<p style="margin-top:-5px;"></p>
In a sense, the protagonist is a reincarnation target, but he perseveres while messing with God’s plans.<p style="margin-top:-5px;"></p>
God makes a mistake in choosing his reincarnation target, and later learns of it.<p style="margin-top:-5px;"></p>
He later realizes that it’s not like everyone wants to be strong and not everyone wants to make a harem.<p style="margin-top:-5px;"></p>
And that he has chosen someone who is more attentive than others and is a real pain in the ass…
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_50621.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (4.1)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid50621" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/the-story-of-a-manga-artist-who-was-imprisoned-by-a-strange-high-school-girl/">The Story of a Manga Artist Who Was Imprisoned by a Strange High School Girl</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 18 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 6.5 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 251 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 0 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="9" href="https://www.novelupdates.com/genre/drama/" title="View All Drama Related Series" >Drama</a> <a class="gennew search" gid="292" href="https://www.novelupdates.com/genre/ecchi/" title="View All Ecchi Related Series" >Ecchi</a> <a class="gennew search" gid="343" href="https://www.novelupdates.com/genre/horror/" title="View All Horror Related Series" >Horror</a> <a class="gennew search" gid="245" href="https://www.novelupdates.com/genre/mystery/" title="View All Mystery Related Series" >Mystery</a> <a class="gennew search" gid="486" href="https://www.novelupdates.com/genre/psychological/" title="View All Psychological Related Series" >Psychological</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="18" href="https://www.novelupdates.com/genre/seinen/" title="View All Seinen Related Series" >Seinen</a> <a class="gennew search" gid="7" href="https://www.novelupdates.com/genre/slice-of-life/" title="View All Slice of Life Related Series" >Slice of Life</a></div>"Where am I?" When I awoke, I noticed a strange ceiling, no lights around me, and a――collar and chain around my neck!? The room's light<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"> suddenly turns on. There was a high school girl holding a kitchen knife. A JK and a manga artist, their life of confinement is about to begin!
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_490.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (3.8)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid490" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/manowa/">Manowa</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 442 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 9.1 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 4682 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 12 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a></div>When she came to, she was in another world. A world just like the stage of a popular game, the girl “Kazane” collects and fights with monster skills in a Hyper Learning Story.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
Before she realizes it, she’s saved a town, exposed a scandal, met a ghost, saved a princess, went into a hot spring, and spent her time going to the hot spring daily with her party. In the fight against the approaching devil, Kazane confronts her own dark side during her destined reunion with the Big Cat. Did I mention she also spends time at a hot spring?
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/noimagemid.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (3.7)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid53841" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/an-s-class-adventurer-banished-as-the-tank-of-the-party-uses-his-slave-release-skill-to-build-the-strongest-country-in-history/">An S-class Adventurer, Banished As The Tank of the Party, uses his &#8220;S*ave Release&#8221; skill to Build the Strongest Country in History!</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 5 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.4 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 106 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 0 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a></div>I became an S-class adventurer with my childhood friends, using the strength that I’m proud of to attract enemies and have my friends defeat them in the meantime, continuing to play the so-called “tank” role.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
I was a paladin, a job specializing in protecting my companions, but perhaps because everyone was protected by me, they did not learn any defense or evasion skills, but only improved their attack skills.<p style="margin-top:-5px;"></p>
In such a situation, they forgot that they are protected by my skills, and they fired me …… Wait, I don’t care what happens, okay!? Even though your natural defensive power are as good as paper.<p style="margin-top:-5px;"></p>
But since I had been through a terrible time on my way up to become an S-class adventurer, I decided to take it easy for once and went to the Adventurers’ Guild to accept a request for a relaxing day alone.<p style="margin-top:-5px;"></p>
I was told that they had tried to cultivate the land in the past, but the monsters that appeared were so strong that all the farmers had run away, so the vast land would be given as a reward to anyone who succeeds. they were rewarded with a large tract of land.<p style="margin-top:-5px;"></p>
When I went there after being told that “Alex, an S-class adventurer, would be fine!” …… I found that there was not even a village or a house around it, and it’s beyond of what you call the countryside, because it was a land inhabited by monsters.<p style="margin-top:-5px;"></p>
The land was deserted, with farming tools left unattended, and there was a monster that I had never seen before, and when I managed to defeat it, I acquired the “s*ave Release” skill.<p style="margin-top:-5px;"></p>
I was told that if I used this skill, it would summon one person of the opposite sex, who is ens*aved somewhere in the world, to me.<p style="margin-top:-5px;"></p>
I thought it would be better for the world if I used this skill, and I activated it immediately, and a beautiful elf girl appeared.<p style="margin-top:-5px;"></p>
After that, a girl with architecture skill, a princess from a foreign country who was a prisoner …… and I kept freeing people who were ens*aved, and before I knew it, a country was formed and I was treated like a king.<p style="margin-top:-5px;"></p>
We received a special award in the Otherworldly Fantasy category at Kakuyomukon 6!<p style="margin-top:-5px;"></p>
≪With the decision to publish a book, the old title “No need for a tank role! I was banished from the S-class party even though I helped my paper defense friends, but I was given the «s*ave Release» skill in the undeveloped land, and while being loved by the beautiful girls who helped me … the strongest country in history was made. The title has been changed.≫<p style="margin-top:-5px;"></p>
*Episode XX: Protagonist’s Point of View<p style="margin-top:-5px;"></p>
Episode X: The point of view of the character written in the title.<p style="margin-top:-5px;"></p>
The following is an example.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/noimagemid.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid53749" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/i-the-clock-user-was-kicked-out-of-the-craftsmens-guild-for-being-incompetent-but-at-the-dungeons-depths-i-awakened-my-true-power-even-though-i-was-told-to-return-because-they-couldnt/">I, the &#8220;Clock User&#8221;, was kicked out of the craftsmen&#8217;s guild for being incompetent, but at the dungeon&#8217;s depths, I awakened my true power &#8211; even though I was told to return because they couldn&#8217;t handle the work, it&#8217;s too late now, and I will live freely as an SSS-class adventurer.</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 14 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.4 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 89 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 0 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a></div>Cyclo Owen. As a genius boy engineer, he was a person of great promise in the circles of magical tool production. However, at the age of 16, he was given the job skill of "Clock User" in a skill-giving ceremony. In this world, it was forbidden for those with job skills to work in any other occupation. Thus, Cyclo's path as a genius engineer was cut off.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
However, Cyclo did not despair and worked diligently as a watchmaker and contributed to the craftsmen's guild by using his job skill as a Clock User. One day, however, the guildmaster of the craftsmen's guild banished him from the guild, saying that the guild did not need a watchmaker, a profession with little usefulness. Cyclo, who has lost his job, faces further despair. He happens to hear that a police officer is planning to r*pe Miranda Lilibel, the neighborhood apothecary's lady. Since the police is involved, he cannot carelessly report the crime. So Cyclo decides to stop the plan that night. He tries to restrain the police officer who is about to attack Miranda's house, but he is detained by the officer instead. He is then accused of trying to r*pe Miranda, and when Miranda hears the commotion, she tells Cyclo she has misjudged him.<p style="margin-top:-5px;"></p>
Cyclo is thus made a criminal for attempted r*pe, and is sent to a remote area to explore the worst dungeon in the world. In the frontier, imprisoned adventurers are considered to be at the bottom of the heap, so Cyclo is looked down upon. Furthermore, due to rumors spread by his younger sister, Alice Owen, a well-known genius alchemist among adventurers, he is treated as incompetent here as well.<p style="margin-top:-5px;"></p>
He is then mistreated by adventurers who makes uses him as a baggage carrier, and is pushed down to the bottom of a valley in the dungeon.<p style="margin-top:-5px;"></p>
At the very bottom of the dungeon, Cyclo wished he could use his Clock User abilities to stop his internal clock. Then his hunger stopped. This led Cyclo to understand that his skill coul manipulate anything that could be counted as a clock.<p style="margin-top:-5px;"></p>
Thus awakening to his true power, Cyclo escaped from the deepest part of the worst dungeon. In recognition of his achievements and abilities, he is recognized as an SSS-class adventurer.<p style="margin-top:-5px;"></p>
On the other hand, those who had driven Cyclo to misfortune are now faced with misfortune in Cyclo's absence.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_12586.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (3.5)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid12586" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/mechanical-god-emperor/">Mechanical God Emperor</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 1233 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.9 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 3899 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 37 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="10" href="https://www.novelupdates.com/genre/mecha/" title="View All Mecha Related Series" >Mecha</a> <a class="gennew search" gid="11" href="https://www.novelupdates.com/genre/sci-fi/" title="View All Sci-fi Related Series" >Sci-fi</a> <a class="gennew search" gid="3954" href="https://www.novelupdates.com/genre/xuanhuan/" title="View All Xuanhuan Related Series" >Xuanhuan</a></div>Yang Feng somehow transmigrated into a different world and received a legacy of an ‘ancient high tech’ family, which does not directly raise his power, but gives him the technology to build things which are way more advanced than the seemingly medieval world.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
But to build something you need resources and energy. To receive resources you need strength. To gain strength you need knowledge. To gain knowledge you… need strength? or a background? Or maybe a fully armed army of high tech robots who aren’t afraid of death?<p style="margin-top:-5px;"></p>
But is this legacy really for him to keep?
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_47256.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (4.6)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid47256" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/strange-memories-with-you-my-dear/">Strange Memories With You, My Dear</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 16 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 13.3 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 261 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 0 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="245" href="https://www.novelupdates.com/genre/mystery/" title="View All Mystery Related Series" >Mystery</a> <a class="gennew search" gid="486" href="https://www.novelupdates.com/genre/psychological/" title="View All Psychological Related Series" >Psychological</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="16" href="https://www.novelupdates.com/genre/supernatural/" title="View All Supernatural Related Series" >Supernatural</a> <a class="gennew search" gid="132" href="https://www.novelupdates.com/genre/tragedy/" title="View All Tragedy Related Series" >Tragedy</a></div>I’ll come to your aid. I’m not going to give up no matter what―<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
In his dreams, Yukinari has been interacting and spending time with a mysterious girl. He experiences a sensation of déjà vu one day when he visits a place for the first time, and there he meets Yuuko, the girl from his dream.And, to her astonishment, she had met Yukinari in a dream as well. Then they start to go out together to confirm the phenomenon, and eventually fall in love. But one day, when everything appeared to be going well, Yuuko fainted and went into a coma….. A story about two people who have the same dream and are looking for a miracle――
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_41062.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (3.8)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid41062" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/youre-my-belated-happiness/">You&#8217;re My Belated Happiness</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 111 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 3.3 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 835 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 6 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="7" href="https://www.novelupdates.com/genre/slice-of-life/" title="View All Slice of Life Related Series" >Slice of Life</a></div>-- All the reunions with the person you had a crush on in the past are long-cherished schemes.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
Ruan Yu’s web novel on Jinjiang, “Really Want to Nibble on Your Ear”, has been accused of plagiarism.<p style="margin-top:-5px;"></p>
She posted on Weibo: God knows, this story about a crush is from my own personal experience back when I was still a student.<p style="margin-top:-5px;"></p>
The other author called someone on the phone after seeing it: “Brother, I seem…… to have found the person you have a crush on.”<p style="margin-top:-5px;"></p>
A few days later, Ruan Yu, who had been “exposed,” looked at the prototype of the male lead in her novel and kept shaking her hand: “I don’t know him, don’t know him……” <p style="margin-top:-5px;"></p>
Xu Huaisong smiled with gritted teeth: You’ve already nibbled my ear. It’s too late to play dumb.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_53052.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid53052" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/cannon-fodder-fake-master-was-stunned-after-being-reborn/">Cannon Fodder Fake Master Was Stunned After Being Reborn</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 24 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 2.2 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 1143 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 10 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="9" href="https://www.novelupdates.com/genre/drama/" title="View All Drama Related Series" >Drama</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="7" href="https://www.novelupdates.com/genre/slice-of-life/" title="View All Slice of Life Related Series" >Slice of Life</a> <a class="gennew search" gid="16" href="https://www.novelupdates.com/genre/supernatural/" title="View All Supernatural Related Series" >Supernatural</a> <a class="gennew search" gid="560" href="https://www.novelupdates.com/genre/yaoi/" title="View All Yaoi Related Series" >Yaoi</a></div>It was only after Jian Xingsui’s death that he realized that he was the cannon fodder in a book with a dog-blooded plot. He was a dumb and s*upid fake young master in it.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
​In the book, he was reluctant to return to his poor ”biological” parents. He was greedy of the Jian Family’s wealth and also wanted to suppress the protagonist shou of the book, the real young master. With the tricks that he had personally thought out of his pig-headed brain, he finally succeeded in making everyone in the family annoy the hell out of him:<p style="margin-top:-5px;"></p>
His former CEO elder brother: “I don’t have a younger brother like you.”<p style="margin-top:-5px;"></p>
His former star second brother: “Never mention to anyone in the entertainment industry that I used to be your brother!!!”<p style="margin-top:-5px;"></p>
His former father: “Get out, I’m so disappointed in you.”<p style="margin-top:-5px;"></p>
​​After that, he was kicked out of the house and had a terrible ending.<p style="margin-top:-5px;"></p>
After the rebirth, just when it was the day when the protagonist shou came up to his door, instead of waiting for his family to talk to him, he took the initiative: “You guys won’t have to say anything, I am willing to leave the house!”
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_48680.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.1)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid48680" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/dressed-as-the-school-grass-ex-boyfriend/">Dressed as The School Grass&#8217; Ex-boyfriend</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 77 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 2 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 1091 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 5 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="6" href="https://www.novelupdates.com/genre/school-life/" title="View All School Life Related Series" >School Life</a> <a class="gennew search" gid="560" href="https://www.novelupdates.com/genre/yaoi/" title="View All Yaoi Related Series" >Yaoi</a></div>Jing Ci: Math problems are not easy to do, and aren’t exams not fun? So why not fall in love?<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
No interest, impossible, time-consuming. <p style="margin-top:-5px;"></p>
The school bully’s arrogant words—-<p style="margin-top:-5px;"></p>
“Go bother someone else if you want to talk about love. I’m not interested.”<p style="margin-top:-5px;"></p>
“Just kidding, what kind of thing is Jing Ci? You think dad will look at him more?” <p style="margin-top:-5px;"></p>
Later–<p style="margin-top:-5px;"></p>
“You see the one that scored first place in the exam? That’s my boyfriend.” <p style="margin-top:-5px;"></p>
“Come on, Jing Ci, which one do you choose? Mathematics or me.”<p style="margin-top:-5px;"></p>
Later–<p style="margin-top:-5px;"></p>
The school bully pressed Jing Ci against the wall: “Good, say something nice, and I will let you go.”
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/noimagemid.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid29237" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/heros-redo-a-hero-that-once-saved-the-world-reborn-as-a-girl/">Hero&#8217;s Redo ~ A Hero That Once Saved the World Reborn as a Girl ~</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 145 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 1.2 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 2169 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 11 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="168" href="https://www.novelupdates.com/genre/gender-bender/" title="View All Gender Bender Related Series" >Gender Bender</a> <a class="gennew search" gid="7" href="https://www.novelupdates.com/genre/slice-of-life/" title="View All Slice of Life Related Series" >Slice of Life</a></div>The Hero was traumatized by politics and people afraid of him for being too powerful, so the mage in his party decided to turn him<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"> into a girl. But he was so traumatized to interact with humans that he just hid inside the mountains until two girls who came to the mountain were caught in danger. After saving them, they decided to bring her to their town to pay her back for helping them.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_25273.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (3.8)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid25273" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/shangrila-frontier-shitty-games-hunter-challenges-godly-game/">ShangriLa Frontier ~ Shitty Games Hunter Challenges Godly Game ~</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 390 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 2 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 1367 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 5 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="11" href="https://www.novelupdates.com/genre/sci-fi/" title="View All Sci-fi Related Series" >Sci-fi</a> <a class="gennew search" gid="12" href="https://www.novelupdates.com/genre/shounen/" title="View All Shounen Related Series" >Shounen</a></div>For every God out there, there are about thousand shitty games in this world.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
Bugs, Errors, Texture Collapses, Inconsistent Scenarios…… Those are the things that fill the players around the world with grief and remorse.<p style="margin-top:-5px;"></p>
A certain boy who loves such shitty games decides to challenge a Godly game recognized by the general public for a change.<p style="margin-top:-5px;"></p>
As a result, both the gaming world and the real world in which the boy is revolving begin to change. Still, the specs of the Godly game still fills the boy’s heart with dread.<p style="margin-top:-5px;"></p>
「Are Encounter Rates like that really a common sense in Godly Games……?」
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_53511.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid53511" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/after-becoming-the-alpha-protagonist-i-snatched-the-cannon-fodder-omega/">After Becoming the Alpha Protagonist, I Snatched the Cannon Fodder Omega</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 6 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 2.2 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 546 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 3 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="10" href="https://www.novelupdates.com/genre/mecha/" title="View All Mecha Related Series" >Mecha</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="6" href="https://www.novelupdates.com/genre/school-life/" title="View All School Life Related Series" >School Life</a> <a class="gennew search" gid="11" href="https://www.novelupdates.com/genre/sci-fi/" title="View All Sci-fi Related Series" >Sci-fi</a> <a class="gennew search" gid="1692" href="https://www.novelupdates.com/genre/shounen-ai/" title="View All Shounen Ai Related Series" >Shounen Ai</a> <a class="gennew search" gid="560" href="https://www.novelupdates.com/genre/yaoi/" title="View All Yaoi Related Series" >Yaoi</a></div>The thick-skinned Qiu Zhenyang just transmigrated, and became the crazy, cool alpha protagonist gong in an interstellar ABO campus story.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
He silently looked at the young man glaring at him with misty eyes. This young man was the cannon fodder omega the original owner had recently broken the marriage contract with. The pheromone overflowing all over his body was indicating he was about to enter estrus.<p style="margin-top:-5px;"></p>
Do or not do?<p style="margin-top:-5px;"></p>
This is a question worth thinking about...<p style="margin-top:-5px;"></p>
---------<p style="margin-top:-5px;"></p>
Transmigrating into the young master of the Federation's Chief Legion, the First Consortium, and a Pharmacy Family, Qiu Zhenyang has whatever he wants. Whoever is not pleasing to the eye, he can deal with.<p style="margin-top:-5px;"></p>
On his first day in the Interstellar Academy, facing the confession of the omega school flower: Sorry, you aren't my type.<p style="margin-top:-5px;"></p>
Inadvertently, hero saves the beauty. Facing the protagonist shou in the original text: Go away, I don't drink green tea.<p style="margin-top:-5px;"></p>
Besieged at home, facing the pursuit of the coquettish childhood sweetheart: Why are you wearing such showy clothes?<p style="margin-top:-5px;"></p>
Qiu Zhenyang's pheromone tastes like sunshine, and his personality is like a little sun. But only when facing Ling Mu, the cannon fodder covered in thorns, did heart soften again and again. His affection grew stronger, and he loved him uncontrollably.<p style="margin-top:-5px;"></p>
---------<p style="margin-top:-5px;"></p>
After the marriage contract with the Qiu family's young master was called off, Ling Mu's difficult life became even worse.<p style="margin-top:-5px;"></p>
But, just over a month later, he had an affair with the young master, and was even temporarily marked!<p style="margin-top:-5px;"></p>
That bastard!<p style="margin-top:-5px;"></p>
Qiu Zhenyang, who temporarily marked him: "I will be responsible for you!"<p style="margin-top:-5px;"></p>
Ling Mu: "Shut up, go away!"<p style="margin-top:-5px;"></p>
Qiu Zhenyang, who officially introduced Ling Mu to his friends: "This is my wife!"<p style="margin-top:-5px;"></p>
Ling Mu: "What wife, go away!"<p style="margin-top:-5px;"></p>
Qiu Zhenyang, who had just rejected the protagonist's confession: "Oh, are you jealous?"<p style="margin-top:-5px;"></p>
Ling Mu: "F*ck, go away!"<p style="margin-top:-5px;"></p>
Qiu Zhenyang: "Are you serious? Then... Xiao Mumu, let's do something harmonious?"<p style="margin-top:-5px;"></p>
Ling Mu: "..."<p style="margin-top:-5px;"></p>
After a few seconds, Qiu Zhenyang: "Ow! Wife, your fists are so powerful!"<p style="margin-top:-5px;"></p>
---------<p style="margin-top:-5px;"></p>
My pheromone is sunshine, your pheromone is sunflower, we are born to fit together.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_19687.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (3.4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid19687" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/the-sovereigns-ascension/">The Sovereign&#8217;s Ascension</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 1083 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.5 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 1351 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 14 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="3954" href="https://www.novelupdates.com/genre/xuanhuan/" title="View All Xuanhuan Related Series" >Xuanhuan</a></div>As a lawyer back on earth, Lin Yun had never lost a case. He owed his success to three things: two blessings he received at birth (a photographic memory and the ability to comprehend anything he studies) and an indomitable will he forged himself.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
While on a trip in Shandong province, he decided he would pay a visit to Mt. Tai. Just as he was cresting the peak, he felt a sharp pain in his chest and his vision went black.<p style="margin-top:-5px;"></p>
Upon waking up, he found himself in the world of Profound Amber occupying the body of a sword s*ave who had shared his name. The last thing he remembered before dying was the image of a sword piercing his chest.<p style="margin-top:-5px;"></p>
Through the memories his body retained the sword s*ave’s life, Lin Yun came to understand the brutality of this world. If he sought respect, he would have to earn it through strength. The weak found no compassion here.<p style="margin-top:-5px;"></p>
Refusing to leave his fate in the hands of others, Lin Yun set out to become a sovereign. No man or beast would stop him from achieving his destiny.<p style="margin-top:-5px;"></p>
With his sword in hand, he would overcome any obstacle.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div>               								</div>
							</div>
                                
			<div class="digg_pagination" style=""><em class="current">1</em><a href="//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc&pg=2">2</a><a href="//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc&pg=3">3</a><span class="my_popup_a_open">…</span><a href="//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc&pg=496">496</a><a href="//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc&pg=2" rel="next" class="next_page"> →</a></div><div id="my_popup_b">
  	<div class="page_digg_bg">
  	<div class="page_digg_title">Go to Page...</div>
    <input type="text" name="p_num_b" id="p_num_b" onkeydown="if (event.keyCode == 13) { document.getElementById('mypu_b_button').click(); }">
    <button onclick="gotoPage('//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc&pg=xxxxx','1');" class="my_popup_b_button" id="mypu_b_button">Go</button>
    </div>
  </div>
  
    <div id="my_popup_a">
  	<div class="page_digg_bg">
  	<div class="page_digg_title">Go to Page...</div>
    <input type="text" name="p_num_a" id="p_num_a" onkeydown="if (event.keyCode == 13) { document.getElementById('mypu_a_button').click(); }">
    <button onclick="gotoPage('//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc&pg=xxxxx','2');" class="my_popup_a_button" id="mypu_a_button">Go</button>
    </div>
  </div>        

<script>
    $(document).ready(function() {
     // Initialize the plugin
$('.fa.fa-info-circle.sfinder').on({
    click: function(event) {
        $('#my_popupsfinder').popup({
            tooltipanchor: event.target,
			vertical: 'top',
            autoopen: true,
            type: 'tooltip'
        });
    }
});
    });
</script>
  
  <div id="my_popupsfinder" style="display:none;">
  <div class="pop-title">Title</div>
  <div class="pop-content">...popup content...</div>
  <div class="pop-footer"><button class="my_popupsfinder_close">OK</button></div>
  <div class="arrow_down_sfinder"></div>
  </div>
  
  <span class="popcontent_noveltype" style="display:none;">The type of novel.</span>
  <span class="popcontent_language" style="display:none;">The language of the novel.</span>
  
  <span class="popcontent_releases" style="display:none;">This is the number of releases (chapters).</br></br>
  <ul>
  <li><b>Min</b> - Minimum amount of chapters.</li>
  <li><b>Max</b>- Maximum amount of chapters.</li>
 </ul></span>
 
  <span class="popcontent_frequency" style="display:none;">The release frequency of a novel. Higher frequency means the novel is updated more often.</br></br>
  <ul>
  <li><b>Min</b> - Minimum release frequency. Lower frequency = slower updates.</li>
  <li><b>Max</b>- Maximum release frequency. Higher frequency = quicker updates.</li>
 </ul></span>
 
   <span class="popcontent_rating" style="display:none;">Novel rating on a scale of 1 to 5.</br></br>
  <ul>
  <li><b>Min</b> - Minimum rating (1 to 5)</li>
  <li><b>Max</b> - Maximum rating (1 to 5)</li>
 </ul></span>
 
   <span class="popcontent_numberratings" style="display:none;">The amount of ratings for a novel.</br></br>
  <ul>
  <li><b>Min</b> - Minimum number of ratings.</li>
  <li><b>Max</b> - Maximum number of ratings.</li>
 </ul></span>
 
    <span class="popcontent_readers" style="display:none;">The number of readers a novel has.</br></br>
  <ul>
  <li><b>Min</b> - Minimum number of readers.</li>
  <li><b>Max</b> - Maximum number of readers.</li>
 </ul></span>
 
   <span class="popcontent_lastdate" style="display:none;">The last release date.</span>
   
   <span class="popcontent_genre" style="display:none;">The genres of the novel.</br></br>
     <ul>
  <li><b>Include</b> - Click a genre once to include it in your search.</li>
  <li><b>Exclude</b> - Click a genre twice to exclude it in your search.</li>
   <li><b>AND</b> - Matches ALL genres selected.</li>
  <li><b>OR</b> - Matches ANY of the genres selected.</li>
 </ul></span>
 
   <span class="popcontent_tags" style="display:none;">The tags of the novel.</br></br>   <ul>
   <li><b>AND</b> - Matches ALL tags selected.</li>
  <li><b>OR</b> - Matches ANY of the tags selected.</li>
 </ul></span>
 
   <span class="popcontent_readinglists" style="display:none;">Your reading lists are listed here.</br></br><ul>
   <li><b>Include</b> - Searches only your reading lists.</li>
  <li><b>Exclude</b> - Exclude series that matches your reading lists.</li>
 </ul> </span>
 
   <span class="popcontent_complete" style="display:none;">Completely translated option.</span>
   <span class="popcontent_sorting" style="display:none;">You can sort results by selecting the options from the drop down list.</span>
<script>
$(document).ready(function() 
    { 
		$('#sortorder option[value="desc"]').prop('selected', true);
$('#sortresults option[value="sdate"]').prop('selected', true);
		$(".sortresults").chosen({disable_search_threshold: 20,width: "90%"});
		$(".sortorder").chosen({disable_search_threshold: 20,width: "90%"});
		$(".storystatus").chosen({disable_search_threshold: 20,width: "40%"});
		$(".ctranslate").chosen({width: "40%"});
		$("#ardate").datepicker();
		$("#ardate_first").datepicker();
		
		$(".tablesorter").tablesorter(); 
		
		$(".chzn-select").chosen({width: "90%", search_contains:true});$("#hdlist").chosen({width: "95%"});	} 
); 
</script>
<script>

function getgrouptags()
{
		var strType = '';
		
		$('#tags_include').empty();
		$('#tags_exclude').empty();
		
		$("#tags_include").attr('disabled', true).trigger("chosen:updated")
		$("#tags_exclude").attr('disabled', true).trigger("chosen:updated")
		
		var chkIcon = $('.tag_icon').attr('current');
		
		if ( chkIcon == 'org' )
		{
			$('.fa.fa-tag').addClass('fa-spin');
			strType = 'tag';
		}
		else
		{
		   	$('.fa.fa-undo').addClass('fa-spin');
			strType = 'undo';
		}
		
	
        $.ajax({
            type:"POST",
            url: "https://www.novelupdates.com/wp-admin/admin-ajax.php",
            data: {
                action:'nd_gettags',
				strType : strType
            },
            success:function(response){
				response = response.slice(0,-1);


				$("#tags_include").append(response);		
				$('#tags_include').trigger("chosen:updated");
				$("#tags_exclude").append(response);
				$('#tags_exclude').trigger("chosen:updated");
			
				$("#tags_include").attr('disabled', false).trigger("chosen:updated")
				$("#tags_exclude").attr('disabled', false).trigger("chosen:updated")
			
				if ( chkIcon == 'org' )
				{
					$('.fa.fa-tag').removeClass('fa-spin');
					$('.tag_icon').replaceWith('<span class="tag_icon" current="extra"><i style="cursor:pointer;" onclick="getgrouptags();" class="fa fa-undo" aria-hidden="true"></i></span>');
				}
				else
				{
					$('.fa.fa-undo').removeClass('fa-spin');
					$('.tag_icon').replaceWith('<span class="tag_icon" current="org"><i style="cursor:pointer;" onclick="getgrouptags();" class="fa fa-tag" aria-hidden="true"></i></span>');
				}
				
            }}); 
}

rl_open_start();
function showtext(n){$(n).parent().find("span").show(),$(n).parent().find(".dots").hide(),$(n).hide()}function hidetext(n){$(n).parent().parent().find(".testhide").hide(),$(n).parent().parent().find(".dots").show(),$(n).parent().parent().find("span.morelink").show()}

</script>	
			</div>

			<div class="l-sidebar at_right">
				

<!-- begin generated sidebar -->
<div id="text-2" class="widget widget_text"><h3 class="widgettitle">Novel Updates</h3>			<div class="textwidget"><ul>
<li><a href="https://forum.novelupdates.com/">Forum</a></li>
<li><a href="https://www.novelupdates.com/random-novel/">Random Novel</a></li>
<li><a href="https://www.novelupdates.com/series-finder/">Series Finder</a></li>
<li><a href="https://www.novelupdates.com/novelslisting/">Series Listing</a></li>
<li><a href="https://www.novelupdates.com/series-ranking/">Series Ranking</a></li>
<li><a href="https://www.novelupdates.com/latest-series/">Latest Series</a></li>
<li><a href="https://www.novelupdates.com/recommendation-lists/">Rec Lists</a></li>
</ul></div>
		</div><div id="text-3" class="widget widget_text"><h3 class="widgettitle">User Tools</h3>			<div class="textwidget"><ul>
<li><a href="https://www.novelupdates.com/reading-list/">Reading List</a></li>
<li><a href="https://www.novelupdates.com/release-filtering/">Release Filtering</a></li>
<li><a href="https://www.novelupdates.com/series-filtering/">Series Filtering</a></li>
</ul></div>
		</div><style>.search_img { float: left;}.search_img img { height: 71px; width: 50px; }</style><div style="margin-bottom: 25px;"><div class="search_main_box"><div class="search_img">
			<a href="https://www.scribblehub.com/series/471295/a-dragon-idols-reincarnation-tale/"><img src="https://cdn.scribblehub.com/seriesimg/mid/23/mid_471295.jpg"></a>
			</div><div class="search_body" style="line-height: 20px; margin-left: 55px;"><div style="font-size:14px; position: relative; top: -3px; max-height: 40px; overflow: hidden;"><a href="https://www.scribblehub.com/series/471295/a-dragon-idols-reincarnation-tale/">A Dragon Idol's Reincarnation Tale</a></div><div><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i></div></div></div><div style="clear: both;"></div><div class="search_main_box"><div class="search_img">
			<a href="https://www.scribblehub.com/series/278470/retribution-engine-arc-/"><img src="https://cdn.scribblehub.com/seriesimg/mid/13/mid_278470.jpg"></a>
			</div><div class="search_body" style="line-height: 20px; margin-left: 55px;"><div style="font-size:14px; position: relative; top: -3px; max-height: 40px; overflow: hidden;"><a href="https://www.scribblehub.com/series/278470/retribution-engine-arc-/">Retribution Engine ARC 2</a></div><div><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div></div></div><div style="clear: both;"></div><div class="search_main_box"><div class="search_img">
			<a href="https://www.scribblehub.com/series/366931/the-hero-has-many-flaws-a-journey-in-another-world/"><img src="https://cdn.scribblehub.com/seriesimg/mid/18/mid_366931.jpg"></a>
			</div><div class="search_body" style="line-height: 20px; margin-left: 55px;"><div style="font-size:14px; position: relative; top: -3px; max-height: 40px; overflow: hidden;"><a href="https://www.scribblehub.com/series/366931/the-hero-has-many-flaws-a-journey-in-another-world/">The Hero has many flaws, a...</a></div><div><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div></div></div><div style="clear: both;"></div></div><h3 class="widgettitle_nuf" style="margin-bottom: 5px;">Latest Discussions</h3><div class="nuf_bd"><ul id="nuf_ul"><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="Regressor Instruction Manual" class="nuftrend_link" href="http://forum.novelupdates.com/threads/120915/unread">Regressor Instruction Manual</a></div><div class="additionalRow muted">4h, 14m ago</div></div></li><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="The Villain Wants to Live" class="nuftrend_link" href="http://forum.novelupdates.com/threads/134041/unread">The Villain Wants...</a></div><div class="additionalRow muted">4h, 23m ago</div></div></li><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="Clearing an Isekai with the Zero-Believers Goddess – The Weakest Mage among the Classmates (WN)" class="nuftrend_link" href="http://forum.novelupdates.com/threads/126095/unread">Clearing an Isekai...</a></div><div class="additionalRow muted">10h, 7m ago</div></div></li><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="Absolute Resonance" class="nuftrend_link" href="http://forum.novelupdates.com/threads/133372/unread">Absolute Resonance</a></div><div class="additionalRow muted">13h, 46m ago</div></div></li><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="The Protagonists Are Murdered by Me" class="nuftrend_link" href="http://forum.novelupdates.com/threads/110159/unread">The Protagonists Are...</a></div><div class="additionalRow muted">16h, 0m ago</div></div></li><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="The Death Mage Who Doesn&#8217;t Want a Fourth Time" class="nuftrend_link" href="http://forum.novelupdates.com/threads/91000/unread">The Death Mage...</a></div><div class="additionalRow muted">23h, 44m ago</div></div></li><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="Possessing Nothing" class="nuftrend_link" href="http://forum.novelupdates.com/threads/40682/unread">Possessing Nothing</a></div><div class="additionalRow muted">1d, 19h, 33m ago</div></div></li></ul></div><!-- TAGNAME: 300x600_Parent_SmartSlot --><div id='div-gpt-ad-novelupdatescom40265'></div>
<!-- end generated sidebar -->

			</div>
		</div>
	</div>



</div>
<!-- /MAIN -->

</div>
<!-- /CANVAS -->

<!-- FOOTER -->
<div class="l-footer">
			<!-- subfooter: bottom -->
	<div class="l-subfooter at_bottom">
		<div class="l-subfooter-h i-cf">

									
			

			<div class="w-copyright"><script>
(function waitGEO() {
    var readyGEO;
    if (window['UnicI'] && window['UnicI'].geo && window['UnicI'].geo !== '-' ) {
        readyGEO = true;
        if (window['UnicI'].geo === 'EU') {
            if(document.getElementById("unic-gdpr")) {
              document.getElementById("unic-gdpr").style.display = 'inline-block';
            }
        }
        if (window['UnicI'].geo === 'CA') {
            if(document.getElementById("unic-ccpa")) {
              document.getElementById("unic-ccpa").style.display = 'inline-block';
            }
        }
    }
    if (!readyGEO) {
        setTimeout(waitGEO, 200);
    }
})();
</script>

<!-- Consent links -->

<a id='unic-gdpr' onclick='__tcfapi("openunic");return false;' style='display:none;cursor:pointer;'>Change Ad Consent</a>
<a id='unic-ccpa' onclick="window.__uspapi('openunic')" style='display:none;cursor:pointer;'>Do not sell my data</a> | <a href="https://www.novelupdates.com/privacy-policy/">Privacy Policy</a> | <a href="https://www.novelupdates.com/terms-of-service/">Terms of Service</a> | <a href="https://www.novelupdates.com/contact-us/">Contact Us</a></div>
					</div>
	</div>
	
</div>
<!-- /FOOTER -->
<a class="w-toplink" href="#"><i class="fa fa-angle-up"></i></a>
<script type="text/javascript">
	if (window.$us === undefined) window.$us = {};
	$us.canvasOptions = ($us.canvasOptions || {});
	$us.canvasOptions.headerDisableStickyHeaderWidth = parseInt('1000');
	$us.canvasOptions.headerDisableAnimationWidth = parseInt('900');
	$us.canvasOptions.headerMainHeight = parseInt('60');
	$us.canvasOptions.headerMainShrinkedHeight = parseInt('60');
	$us.canvasOptions.headerExtraHeight = parseInt('36');
	$us.canvasOptions.mobileNavWidth = parseInt('1000');

	$us.navOptions = ($us.navOptions || {});
	$us.navOptions.togglable = 1;

</script>
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-66522525-1', 'auto');
  ga('send', 'pageview');

</script><script type='text/javascript' src='https://www.novelupdates.com/wp-content/plugins/user-meta-display/assets/js/scripts-user_meta_display.js?ver=1.2.2'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/plugins/yet-another-stars-rating/js/jquery.rateit.min.js?ver=1.0.22'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/plugins/yet-another-stars-rating/js/yasr-front.js?ver=1.0.1'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates/js/jquery.easing.min.js?ver=5.2.7'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates/js/us.core.js?ver=4.4.8'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates/js/us.widgets.js?ver=4.4.8'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-includes/js/comment-reply.min.js?ver=5.2.7'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-includes/js/wp-embed.min.js?ver=5.2.7'></script>
</body>
</html>
`)

var TestFailureRequester = DummyRequester(`
<!DOCTYPE HTML>
<html class="" lang="en-US">
<head>
	<meta charset="UTF-8">
	<title>Series Finder - Novel Updates</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
	<meta name="SKYPE_TOOLBAR" content="SKYPE_TOOLBAR_PARSER_COMPATIBLE" />
    <meta name="referrer" content="unsafe-url">

	    <link rel="apple-touch-icon" href="/appicon.png"/>
   
<!-- UniConsent Cookies Consent Notice start -->
<script src="https://cmp.uniconsent.com/v2/stub.min.js"></script>
<script async src='https://cmp.uniconsent.com/v2/69a34e0934/cmp.js'></script>
<!-- UniConsent Cookies Consent Notice end -->
	
	<link rel='dns-prefetch' href='//www.novelupdates.com' />
<link rel='dns-prefetch' href='//www.google.com' />
<link rel='dns-prefetch' href='//cdn.novelupdates.com' />
<link rel='dns-prefetch' href='//ajax.googleapis.com' />
<link rel='dns-prefetch' href='//fonts.googleapis.com' />
<link rel='dns-prefetch' href='//s.w.org' />
<link rel="alternate" type="application/rss+xml" title="Novel Updates &raquo; Feed" href="https://www.novelupdates.com/feed/" />
<link rel="alternate" type="application/rss+xml" title="Novel Updates &raquo; Comments Feed" href="https://www.novelupdates.com/comments/feed/" />
		<script type="text/javascript">
			window._wpemojiSettings = {"baseUrl":"https:\/\/s.w.org\/images\/core\/emoji\/12.0.0-1\/72x72\/","ext":".png","svgUrl":"https:\/\/s.w.org\/images\/core\/emoji\/12.0.0-1\/svg\/","svgExt":".svg","source":{"concatemoji":"https:\/\/www.novelupdates.com\/wp-includes\/js\/wp-emoji-release.min.js?ver=5.2.7"}};
			!function(a,b,c){function d(a,b){var c=String.fromCharCode;l.clearRect(0,0,k.width,k.height),l.fillText(c.apply(this,a),0,0);var d=k.toDataURL();l.clearRect(0,0,k.width,k.height),l.fillText(c.apply(this,b),0,0);var e=k.toDataURL();return d===e}function e(a){var b;if(!l||!l.fillText)return!1;switch(l.textBaseline="top",l.font="600 32px Arial",a){case"flag":return!(b=d([55356,56826,55356,56819],[55356,56826,8203,55356,56819]))&&(b=d([55356,57332,56128,56423,56128,56418,56128,56421,56128,56430,56128,56423,56128,56447],[55356,57332,8203,56128,56423,8203,56128,56418,8203,56128,56421,8203,56128,56430,8203,56128,56423,8203,56128,56447]),!b);case"emoji":return b=d([55357,56424,55356,57342,8205,55358,56605,8205,55357,56424,55356,57340],[55357,56424,55356,57342,8203,55358,56605,8203,55357,56424,55356,57340]),!b}return!1}function f(a){var c=b.createElement("script");c.src=a,c.defer=c.type="text/javascript",b.getElementsByTagName("head")[0].appendChild(c)}var g,h,i,j,k=b.createElement("canvas"),l=k.getContext&&k.getContext("2d");for(j=Array("flag","emoji"),c.supports={everything:!0,everythingExceptFlag:!0},i=0;i<j.length;i++)c.supports[j[i]]=e(j[i]),c.supports.everything=c.supports.everything&&c.supports[j[i]],"flag"!==j[i]&&(c.supports.everythingExceptFlag=c.supports.everythingExceptFlag&&c.supports[j[i]]);c.supports.everythingExceptFlag=c.supports.everythingExceptFlag&&!c.supports.flag,c.DOMReady=!1,c.readyCallback=function(){c.DOMReady=!0},c.supports.everything||(h=function(){c.readyCallback()},b.addEventListener?(b.addEventListener("DOMContentLoaded",h,!1),a.addEventListener("load",h,!1)):(a.attachEvent("onload",h),b.attachEvent("onreadystatechange",function(){"complete"===b.readyState&&c.readyCallback()})),g=c.source||{},g.concatemoji?f(g.concatemoji):g.wpemoji&&g.twemoji&&(f(g.twemoji),f(g.wpemoji)))}(window,document,window._wpemojiSettings);
		</script>
		<style type="text/css">
img.wp-smiley,
img.emoji {
	display: inline !important;
	border: none !important;
	box-shadow: none !important;
	height: 1em !important;
	width: 1em !important;
	margin: 0 .07em !important;
	vertical-align: -0.1em !important;
	background: none !important;
	padding: 0 !important;
}
</style>
	<link rel='stylesheet' id='dashicons-css'  href='https://www.novelupdates.com/wp-includes/css/dashicons.min.css?ver=5.2.7' type='text/css' media='all' />
<link rel='stylesheet' id='theme-my-login-css'  href='https://www.novelupdates.com/wp-content/plugins/theme-my-login/theme-my-login.css?ver=6.4.10' type='text/css' media='all' />
<link rel='stylesheet' id='wp-block-library-css'  href='https://www.novelupdates.com/wp-includes/css/dist/block-library/style.min.css?ver=5.2.7' type='text/css' media='all' />
<link rel='stylesheet' id='bbspoiler-css'  href='https://www.novelupdates.com/wp-content/plugins/bbspoiler/inc/bbspoiler.css?ver=5.2.7' type='text/css' media='all' />
<link rel='stylesheet' id='cstmsrch_stylesheet-css'  href='https://www.novelupdates.com/wp-content/plugins/custom-search-plugin/css/style.css?ver=5.2.7' type='text/css' media='all' />
<link rel='stylesheet' id='yasrcss-css'  href='https://www.novelupdates.com/wp-content/plugins/yet-another-stars-rating/css/yasr.css' type='text/css' media='all' />
<link rel='stylesheet' id='yasrcsslightscheme-css'  href='https://www.novelupdates.com/wp-content/plugins/yet-another-stars-rating/css/yasr-table-light.css' type='text/css' media='all' />
<link rel='stylesheet' id='breadcrumbs-nu-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/breadcrumbs.css?ver=1.2.5' type='text/css' media='' />
<link rel='stylesheet' id='global-extras-css'  href='//cdn.novelupdates.com/js/global_extras.css?ver=1.0.6' type='text/css' media='' />
<link rel='stylesheet' id='ajax-search-posts-css-css'  href='//cdn.novelupdates.com/js/ajax_search.post.css?ver=1.3.6' type='text/css' media='' />
<link rel='stylesheet' id='alert-nu-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/alert_nu.css?ver=1.4.1' type='text/css' media='' />
<link rel='stylesheet' id='select2-css-css'  href='//cdn.novelupdates.com/js/select2.min.css?ver=1.0.1' type='text/css' media='all' />
<link rel='stylesheet' id='select2-bootstrap-css-css'  href='//cdn.novelupdates.com/js/select2-bootstrap.min.css?ver=1.0.1' type='text/css' media='all' />
<link rel='stylesheet' id='pagination-nu-css'  href='//cdn.novelupdates.com/js/pg/pagination.css?ver=1.1.1' type='text/css' media='' />
<link rel='stylesheet' id='us-font-1-css'  href='https://fonts.googleapis.com/css?family=Noto+Sans%3A300%2C400%2C700&#038;subset=latin&#038;ver=5.2.7' type='text/css' media='all' />
<link rel='stylesheet' id='us-font-2-css'  href='https://fonts.googleapis.com/css?family=Open+Sans%3A400%2C400italic%2C700%2C700italic%2C300&#038;subset=latin&#038;ver=5.2.7' type='text/css' media='all' />
<link rel='stylesheet' id='us-motioncss-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates/css/motioncss.css?ver=4.4.8' type='text/css' media='all' />
<link rel='stylesheet' id='us-motioncss-responsive-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates/css/motioncss-responsive.css?ver=4.4.8' type='text/css' media='all' />
<link rel='stylesheet' id='us-font-awesome-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates/css/font-awesome.css?ver=4.7.1' type='text/css' media='all' />
<link rel='stylesheet' id='us-style-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates/css/style.css?ver=4.4.8' type='text/css' media='all' />
<link rel='stylesheet' id='us-responsive-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates/css/responsive.css?ver=4.4.8' type='text/css' media='all' />
<link rel='stylesheet' id='ndupdates-style-css'  href='//www.novelupdates.com/wp-content/themes/ndupdates-child/style.css?ver=4.4.8' type='text/css' media='all' />
<script type='text/javascript' src='//ajax.googleapis.com/ajax/libs/jquery/1.8.3/jquery.min.js'></script>
<script type='text/javascript'>
/* <![CDATA[ */
var title = {"unfolded":"Expand","folded":"Collapse"};
/* ]]> */
</script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/plugins/bbspoiler/inc/bbspoiler.js?ver=5.2.7'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-includes/js/jquery/ui/core.min.js?ver=1.11.4'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-includes/js/jquery/ui/widget.min.js?ver=1.11.4'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-includes/js/jquery/ui/mouse.min.js?ver=1.11.4'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-includes/js/jquery/ui/sortable.min.js?ver=1.11.4'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/plugins/custom-search-plugin/js/script.js?ver=5.2.7'></script>
<script type='text/javascript' src='https://www.google.com/recaptcha/api.js?hl=en-US&#038;ver=5.2.7'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/plugins/theme-my-login/modules/themed-profiles/themed-profiles.js?ver=5.2.7'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/nu_extras_js.js?ver=1.0.4'></script>
<script type='text/javascript' src='//cdn.novelupdates.com/js/ajax_search_post.js?ver=1.3.6'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/alert_nu.js?ver=1.4.1'></script>
<script type='text/javascript' src='//cdn.novelupdates.com/js/select2.min.js?ver=1.0.1'></script>
<script type='text/javascript' src='//cdn.novelupdates.com/js/popoverlay/jquery.popupoverlay.js?ver=1.0.1'></script>
<script type='text/javascript' src='//cdn.novelupdates.com/js/pg/pagination.js?ver=1.0.2'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/seriesfinder.js?ver=1.0.8'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/chosen.jquery.min.js?ver=1.0.2'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/jquery.tablesorter.min.js?ver=1.0.1'></script>
<link rel='https://api.w.org/' href='https://www.novelupdates.com/wp-json/' />
<link rel="EditURI" type="application/rsd+xml" title="RSD" href="https://www.novelupdates.com/xmlrpc.php?rsd" />
<link rel="wlwmanifest" type="application/wlwmanifest+xml" href="https://www.novelupdates.com/wp-includes/wlwmanifest.xml" /> 
<meta name="generator" content="WordPress 5.2.7" />
<link rel="canonical" href="https://www.novelupdates.com/series-finder/" />
<link rel='shortlink' href='https://www.novelupdates.com/?p=3370' />
<link rel="alternate" type="application/json+oembed" href="https://www.novelupdates.com/wp-json/oembed/1.0/embed?url=https%3A%2F%2Fwww.novelupdates.com%2Fseries-finder%2F" />
<link rel="alternate" type="text/xml+oembed" href="https://www.novelupdates.com/wp-json/oembed/1.0/embed?url=https%3A%2F%2Fwww.novelupdates.com%2Fseries-finder%2F&#038;format=xml" />
 
   
    

	<style id="us_fonts_inline">
body {
		font-family: 'Open Sans';
		font-size: 15px;
	line-height: 26px;
	font-weight: 400;
	}
.page-template-page-blank-php .l-main {
	font-size: 15px;
	}
	
.l-header .menu-item-language,
.l-header .w-nav-item {
		font-family: 'Open Sans';
		font-weight: 300;
	}
.type_desktop .menu-item-language > a,
.l-header .type_desktop .w-nav-anchor.level_1,
.type_desktop [class*="columns"] .menu-item-has-children .w-nav-anchor.level_2 {
	font-size: 16px;
	}
.type_desktop .submenu-languages .menu-item-language > a,
.l-header .type_desktop .w-nav-anchor.level_2,
.l-header .type_desktop .w-nav-anchor.level_3,
.l-header .type_desktop .w-nav-anchor.level_4 {
	font-size: 15px;
	}
.type_mobile .menu-item-language > a,
.l-header .type_mobile .w-nav-anchor.level_1 {
	font-size: 16px;
	}
.l-header .type_mobile .w-nav-anchor.level_2,
.l-header .type_mobile .w-nav-anchor.level_3,
.l-header .type_mobile .w-nav-anchor.level_4 {
	font-size: 15px;
	}

h1, h2, h3, h4, h5, h6,
.w-counter-number,
.w-logo-title,
.w-pricing-item-title,
.w-pricing-item-price,
.w-shortblog-entry-date-day,
.ult_price_figure,
.ult_countdown-amount,
.ultb3-box .ultb3-title,
.stats-block .stats-desc .stats-number {
		font-family: 'Noto Sans';
		font-weight: 300;
	}
h1 {
	font-size: 38px;
	}
h2 {
	font-size: 32px;
	}
h3 {
	font-size: 26px;
	}
h4,
.widgettitle,
.comment-reply-title,
.ultb3-box .ultb3-title,
.flip-box-wrap .flip-box .ifb-face h3,
.aio-icon-box .aio-icon-header h3.aio-icon-title {
	font-size: 20px;
	}
h5,
.w-portfolio-item-title {
	font-size: 20px;
	}
h6 {
	font-size: 18px;
	}
@media only screen and (max-width: 767px) {
body {
	font-size: 13px;
	line-height: 23px;
	}
h1 {
	font-size: 30px;
	}
h2 {
	font-size: 26px;
	}
h3 {
	font-size: 22px;
	}
h4,
.widgettitle,
.comment-reply-title,
.ultb3-box .ultb3-title,
.flip-box-wrap .flip-box .ifb-face h3,
.aio-icon-box .aio-icon-header h3.aio-icon-title {
	font-size: 20px;
	}
h5,
.w-portfolio-item-title {
	font-size: 18px;
	}
h6 {
	font-size: 16px;
	}
}

.l-body,
.headerpos_fixed .l-header {
	min-width: 1240px;
	}
.l-canvas.type_boxed,
.l-canvas.type_boxed .l-subheader,
.l-canvas.type_boxed ~ .l-footer .l-subfooter {
	max-width: 1240px;
	}
.l-subheader-h,
.l-submain-h,
.l-subfooter-h {
	max-width: 1140px;
	}
.l-sidebar {
	width: 25%;
	}
.l-content {
	width: 70%;
	}
@media only screen and (max-width: 767px) {
.g-cols.type_boxed,
.g-cols.type_boxed > div {
	display: block;
	}
.g-cols > div {
	width: 100% !important;
	margin-left: 0 !important;
	margin-bottom: 30px;
	}
.g-cols.offset_none > div,
.g-cols > div:last-child {
	margin-bottom: 0 !important;
	}
}

@media only screen and (min-width: 900px) {
.w-logo-img {
	height: 30px;
	}
.w-logo.with_transparent .w-logo-img > img.for_default {
	margin-bottom: -30px;
	}
.l-header.sticky .w-logo-img {
	height: 30px;
	}
.l-header.sticky .w-logo.with_transparent .w-logo-img > img.for_default {
	margin-bottom: -30px;
	}
}
@media only screen and (min-width: 600px) and (max-width: 899px) {
.w-logo-img {
	height: 30px;
	}
.w-logo.with_transparent .w-logo-img > img.for_default {
	margin-bottom: -30px;
	}
}
@media only screen and (max-width: 599px) {
.w-logo-img {
	height: 30px;
	}
.w-logo.with_transparent .w-logo-img > img.for_default {
	margin-bottom: -30px;
	}
}
</style>
<style id="us_colors_inline">
/*************************** BODY ***************************/

/* Body Background Color */
.l-body {
	background-color: #2c3e50;
	}

/*************************** HEADER ***************************/

/* Header Background Color */
.l-subheader.at_middle,
.l-subheader.at_middle .w-lang-list,
.l-subheader.at_middle .type_mobile .w-nav-list.level_1 {
	background-color: #2c3e50;
	}

/* Header Text Color */
.l-subheader.at_middle,
.transparent .l-subheader.at_middle .type_mobile .w-nav-list.level_1 {
	color: #edf0f2;
	}

/* Header Text Hover Color */
.no-touch .w-logo-link:hover,
.no-touch .l-subheader.at_middle .w-contacts-item-value a:hover,
.no-touch .l-subheader.at_middle .w-lang-item:hover,
.no-touch .transparent .l-subheader.at_middle .w-lang.active .w-lang-item:hover,
.no-touch .l-subheader.at_middle .w-cart:hover .w-cart-link,
.no-touch .l-subheader.at_middle .w-search-show:hover,
.l-subheader.at_middle .w-cart-quantity {
	color: #fc4349;
	}

/* Extended Header Background Color */
.l-subheader.at_top,
.l-subheader.at_top .w-lang-list,
.l-subheader.at_bottom,
.l-subheader.at_bottom .type_mobile .w-nav-list.level_1 {
	background-color: #384b5f;
	}

/* Extended Header Text Color */
.l-subheader.at_top,
.l-subheader.at_bottom,
.transparent .l-subheader.at_bottom .type_mobile .w-nav-list.level_1,
.w-lang.active .w-lang-item {
	color: #d3d8db;
	}

/* Extended Header Text Hover Color */
.no-touch .l-subheader.at_top .w-contacts-item-value a:hover,
.no-touch .l-subheader.at_top .w-lang-item:hover,
.no-touch .transparent .l-subheader.at_top .w-lang.active .w-lang-item:hover,
.no-touch .l-subheader.at_bottom .w-cart:hover .w-cart-link,
.no-touch .l-subheader.at_bottom .w-search-show:hover,
.l-subheader.at_bottom .w-cart-quantity {
	color: #fff;
	}
	
/* Transparent Header Text Color */
.l-header.transparent .l-subheader {
	color: #fff;
	}
	
/* Transparent Header Text Hover Color */
.no-touch .l-header.transparent .type_desktop .menu-item-language > a:hover,
.no-touch .l-header.transparent .type_desktop .menu-item-language:hover > a,
.no-touch .l-header.transparent .w-logo-link:hover,
.no-touch .l-header.transparent .l-subheader .w-contacts-item-value a:hover,
.no-touch .l-header.transparent .l-subheader .w-lang-item:hover,
.no-touch .l-header.transparent .l-subheader .w-search-show:hover,
.no-touch .l-header.transparent .l-subheader .w-cart-link:hover,
.l-header.transparent .l-subheader .w-cart-quantity,
.no-touch .l-header.transparent .type_desktop .w-nav-item.level_1:hover .w-nav-anchor.level_1 {
	color: #fc4349;
	}
.l-header.transparent .w-nav-title:after {
	background-color: #fc4349;
	}
	
/* Search Screen Background Color */
.l-subheader .w-search-form:before {
	background-color: #6dbcdb;
	}

/* Search Screen Text Color */
.l-subheader .w-search-form {
	color: #fff;
	}

/*************************** MAIN MENU ***************************/

/* Menu Hover Background Color */
.no-touch .l-header .menu-item-language > a:hover,
.no-touch .type_desktop .menu-item-language:hover > a,
.no-touch .l-header .w-nav-item.level_1:hover .w-nav-anchor.level_1 {
	background-color: #2c3e50;
	}

/* Menu Hover Text Color */
.no-touch .l-header .menu-item-language > a:hover,
.no-touch .type_desktop .menu-item-language:hover > a,
.no-touch .l-header .w-nav-item.level_1:hover .w-nav-anchor.level_1 {
	color: #fc4349;
	}
.w-nav-title:after {
	background-color: #fc4349;
	}

/* Menu Active Background Color */
.l-header .w-nav-item.level_1.current-menu-item .w-nav-anchor.level_1,
.l-header .w-nav-item.level_1.current-menu-ancestor .w-nav-anchor.level_1 {
	background-color: #2c3e50;
	}

/* Menu Active Text Color */
.l-header .w-nav-item.level_1.current-menu-item .w-nav-anchor.level_1,
.l-header .w-nav-item.level_1.current-menu-ancestor .w-nav-anchor.level_1 {
	color: #6dbcdb;
	}
	
/* Transparent Menu Active Text Color */
.l-header.transparent .type_desktop .w-nav-item.level_1.active .w-nav-anchor.level_1,
.l-header.transparent .type_desktop .w-nav-item.level_1.current-menu-item .w-nav-anchor.level_1,
.l-header.transparent .type_desktop .w-nav-item.level_1.current-menu-ancestor .w-nav-anchor.level_1 {
	color: #6dbcdb;
	}

/* Dropdown Background Color */
.type_desktop .submenu-languages,
.l-header .w-nav-list.level_2,
.l-header .w-nav-list.level_3,
.l-header .w-nav-list.level_4 {
	background-color: #2c3e50;
	}

/* Dropdown Text Color */
.type_desktop .submenu-languages,
.l-header .w-nav-anchor.level_2,
.l-header .w-nav-anchor.level_3,
.l-header .w-nav-anchor.level_4,
.type_desktop [class*="columns"] .w-nav-item.menu-item-has-children.current-menu-item .w-nav-anchor.level_2,
.type_desktop [class*="columns"] .w-nav-item.menu-item-has-children.current-menu-ancestor .w-nav-anchor.level_2,
.no-touch .type_desktop [class*="columns"] .w-nav-item.menu-item-has-children:hover .w-nav-anchor.level_2 {
	color: #edf0f2;
	}

/* Dropdown Hover Background Color */
.no-touch .type_desktop .submenu-languages .menu-item-language:hover > a,
.no-touch .l-header .w-nav-item.level_2:hover .w-nav-anchor.level_2,
.no-touch .l-header .w-nav-item.level_3:hover .w-nav-anchor.level_3,
.no-touch .l-header .w-nav-item.level_4:hover .w-nav-anchor.level_4 {
	background-color: #2c3e50;
	}

/* Dropdown Hover Text Color */
.no-touch .type_desktop .submenu-languages .menu-item-language:hover > a,
.no-touch .l-header .w-nav-item.level_2:hover .w-nav-anchor.level_2,
.no-touch .l-header .w-nav-item.level_3:hover .w-nav-anchor.level_3,
.no-touch .l-header .w-nav-item.level_4:hover .w-nav-anchor.level_4 {
	color: #fc4349;
	}

/* Dropdown Active Background Color */
.l-header .w-nav-item.level_2.current-menu-item .w-nav-anchor.level_2,
.l-header .w-nav-item.level_2.current-menu-ancestor .w-nav-anchor.level_2,
.l-header .w-nav-item.level_3.current-menu-item .w-nav-anchor.level_3,
.l-header .w-nav-item.level_3.current-menu-ancestor .w-nav-anchor.level_3,
.l-header .w-nav-item.level_4.current-menu-item .w-nav-anchor.level_4,
.l-header .w-nav-item.level_4.current-menu-ancestor .w-nav-anchor.level_4 {
	background-color: #2c3e50;
	}

/* Dropdown Active Text Color */
.l-header .w-nav-item.level_2.current-menu-item .w-nav-anchor.level_2,
.l-header .w-nav-item.level_2.current-menu-ancestor .w-nav-anchor.level_2,
.l-header .w-nav-item.level_3.current-menu-item .w-nav-anchor.level_3,
.l-header .w-nav-item.level_3.current-menu-ancestor .w-nav-anchor.level_3,
.l-header .w-nav-item.level_4.current-menu-item .w-nav-anchor.level_4,
.l-header .w-nav-item.level_4.current-menu-ancestor .w-nav-anchor.level_4 {
	color: #6dbcdb;
	}

/* Menu Button Background Color */
.btn.w-nav-item .w-nav-anchor.level_1,
.headerpos_fixed .transparent .btn.w-nav-item .w-nav-anchor.level_1 {
	background-color: #6dbcdb !important;
	}

/* Menu Button Text Color */
.btn.w-nav-item .w-nav-anchor.level_1 {
	color: #fff !important;
	}

/* Menu Button Hover Background Color */
.no-touch .btn.w-nav-item .w-nav-anchor.level_1:before {
	background-color: #fc4349 !important;
	}

/* Menu Button Hover Text Color */
.no-touch .btn.w-nav-item .w-nav-anchor.level_1:hover {
	color: #fff !important;
	}

/*************************** MAIN CONTENT ***************************/

/* Background Color */
.l-canvas,
.w-blog.type_masonry .w-blog-entry-preview:before,
.w-cart-dropdown,
.w-filters-item.active,
.no-touch .w-filters-item.active:hover,
.w-portfolio-item-anchor,
.w-tabs-item.active,
.no-touch .w-tabs-item.active:hover,
.w-timeline-item,
.w-timeline-section-title,
.w-timeline.type_vertical .w-timeline-section-content,
#lang_sel ul ul,
#lang_sel_click ul ul,
#lang_sel_footer,
.woocommerce div.product .woocommerce-tabs .tabs li.active,
.no-touch .woocommerce div.product .woocommerce-tabs .tabs li.active:hover,
.woocommerce .stars span:after,
.woocommerce .stars span a:after,
#bbp-user-navigation li.current,
.gform_wrapper .chosen-container .chosen-drop,
.gform_wrapper .chosen-container-multi .chosen-choices li.search-choice {
	background-color: #fff;
	}
.g-btn.color_contrast,
.no-touch .g-btn.color_contrast:hover,
.no-touch .g-btn.color_contrast.outlined:hover,
.w-icon.color_border.type_solid .w-icon-link,
.w-iconbox.color_contrast.type_solid .w-iconbox-icon {
	color: #fff;
	}

/* Alternate Background Color */
input[type="text"],
input[type="password"],
input[type="email"],
input[type="url"],
input[type="tel"],
input[type="number"],
input[type="date"],
input[type="search"],
textarea,
select,
.w-actionbox.color_alternate,
.w-author,
.w-blog.imgpos_atleft .w-blog-entry-preview-icon,
.w-filters,
.w-icon.color_text.type_solid .w-icon-link,
.w-icon.color_fade.type_solid .w-icon-link,
.w-pricing-item-title,
.w-pricing-item-price,
.w-progbar-bar,
.w-tabs-list,
.no-touch .widget_nav_menu .menu-item a:hover,
.no-touch #lang_sel a:hover,
.no-touch #lang_sel_click a:hover,
.woocommerce .quantity .plus,
.woocommerce .quantity .minus,
.woocommerce div.product .woocommerce-tabs .tabs,
.woocommerce table.shop_table .actions .coupon .input-text,
.woocommerce #payment .payment_box,
#bbp-user-navigation,
#subscription-toggle,
#favorite-toggle,
.bbp-row-actions #favorite-toggle a,
.bbp-row-actions #subscription-toggle a,
.gform_wrapper .chosen-container-single .chosen-single,
.gform_wrapper .chosen-container-multi .chosen-choices,
.select2-container a.select2-choice,
.smile-icon-timeline-wrap .timeline-wrapper .timeline-block,
.smile-icon-timeline-wrap .timeline-feature-item.feat-item {
	background-color: #f2f4f5;
	}
.woocommerce #payment .payment_box:after,
.timeline-wrapper .timeline-post-right .ult-timeline-arrow l,
.timeline-wrapper .timeline-post-left .ult-timeline-arrow l,
.timeline-feature-item.feat-item .ult-timeline-arrow l {
	border-color: #f2f4f5;
	}

/* Border Color */
.l-submain,
.g-cols > div,
.w-blog-entry,
.w-bloglist,
.w-blognav,
.w-comments,
.w-comments-item,
.w-pricing-item-h,
.w-profile,
.w-sharing.type_simple .w-sharing-item,
.w-tabs.layout_accordion,
.w-tabs.layout_accordion .w-tabs-section,
.w-timeline.type_vertical .w-timeline-section-content,
#wp-calendar thead th,
#wp-calendar tbody td,
#wp-calendar tfoot td,
.widget_nav_menu .menu-item a,
#lang_sel a,
#lang_sel_click a,
.woocommerce table th,
.woocommerce table td,
.woocommerce .login,
.woocommerce .checkout_coupon,
.woocommerce .register,
.woocommerce .cart.variations_form,
.woocommerce .commentlist,
.woocommerce .commentlist li,
.woocommerce .comment-respond,
.woocommerce .related,
.woocommerce .upsells,
.woocommerce .cross-sells,
.woocommerce .checkout #order_review,
.woocommerce ul.order_details li,
.woocommerce .shop_table.my_account_orders,
.widget_price_filter .ui-slider-handle,
.widget_layered_nav ul,
.widget_layered_nav ul li,
#bbpress-forums .bbp-body ul.forum,
#bbpress-forums .bbp-forums li.bbp-header,
#bbpress-forums .bbp-body ul.topic,
#bbpress-forums .bbp-topics li.bbp-header,
.bbp-replies .bbp-body,
div.bbp-forum-header,
div.bbp-topic-header,
div.bbp-reply-header,
.bbp-pagination-links a,
.bbp-pagination-links span.current,
span.bbp-topic-pagination a.page-numbers,
.bbp-logged-in,
fieldset,
.form_saved_message,
.gform_wrapper .gsection,
.gform_wrapper .gf_page_steps,
.gform_wrapper li.gfield_creditcard_warning,
.smile-icon-timeline-wrap .timeline-line,
.ult_pricing_table_wrap.ult_design_6 .ult_pricing_table {
	border-color: #e4e8eb;
	}
.g-hr-h i,
.w-icon.color_border .w-icon-link,
.w-iconbox.color_light .w-iconbox-icon {
	color: #e4e8eb;
	}
.g-hr-h:before,
.g-hr-h:after,
.g-btn.color_default,
.g-btn.color_default.outlined:before,
.w-icon.color_border.type_solid .w-icon-link,
.w-iconbox.color_light.type_solid .w-iconbox-icon,
.w-timeline-list:before,
.woocommerce .button,
.no-touch .woocommerce .quantity .plus:hover,
.no-touch .woocommerce .quantity .minus:hover,
.widget_price_filter .ui-slider,
.gform_wrapper .gform_page_footer .gform_previous_button {
	background-color: #e4e8eb;
	}
.g-btn.color_default.outlined,
.pagination .page-numbers,
.w-iconbox.color_light.type_outlined .w-iconbox-icon,
.w-socials-item-link,
.w-tags-item-link,
.w-team-links-item,
.w-testimonial,
.woocommerce-pagination a,
.woocommerce-pagination span {
	box-shadow: 0 0 0 2px #e4e8eb inset;
	}

/* Heading Color */
h1, h2, h3, h4, h5, h6,
input[type="text"],
input[type="password"],
input[type="email"],
input[type="url"],
input[type="tel"],
input[type="number"],
input[type="date"],
input[type="search"],
textarea,
select,
.w-form-field > i,
.no-touch .g-btn.color_default:hover,
.no-touch .g-btn.color_default.outlined:hover,
.g-btn.color_contrast.outlined,
.w-blog-entry-title,
.w-counter-number,
.w-iconbox.color_contrast .w-iconbox-icon,
.w-portfolio-item-anchor,
.no-touch a.w-portfolio-item-anchor:hover,
.l-submain.color_primary a.w-portfolio-item-anchor,
.l-submain.color_secondary a.w-portfolio-item-anchor,
.w-pricing-item-title,
.w-pricing-item-price,
.woocommerce .product .price {
	color: #292e33;
	}
.g-btn.color_contrast,
.g-btn.color_contrast.outlined:before,
.w-iconbox.color_contrast.type_solid .w-iconbox-icon,
.w-progbar.color_contrast .w-progbar-bar-h {
	background-color: #292e33;
	}
.g-btn.color_contrast.outlined,
.w-iconbox.color_contrast.type_outlined .w-iconbox-icon {
	box-shadow: 0 0 0 2px #292e33 inset;
	}

/* Text Color */
.l-canvas,
.g-btn.color_default,
.g-btn.color_default.outlined,
.w-cart-dropdown,
.w-icon.color_text .w-icon-link,
.w-iconbox.color_light.type_solid .w-iconbox-icon,
.color_primary .w-icon.color_text.type_solid .w-icon-link,
.woocommerce .button,
.l-subfooter.at_top .woocommerce .button {
	color: #5c6166;
	}

/* Primary Color */
a,
.g-html .highlight,
.g-btn.color_primary.outlined,
.no-touch .w-blog-entry-link:hover .w-blog-entry-title-h,
.no-touch .w-blog-entry-link:hover,
.l-main .w-contacts-item i,
.w-counter.color_primary .w-counter-number,
.w-filters-item.active,
.no-touch .w-filters-item.active:hover,
.w-form-field > input:focus + i,
.w-form-field > textarea:focus + i,
.w-icon.color_primary .w-icon-link,
.w-iconbox.color_primary .w-iconbox-icon,
.no-touch .w-iconbox-link:hover .w-iconbox-title,
.no-touch .w-pagehead-nav-item:hover,
.w-tabs-item.active,
.no-touch .w-tabs-item.active:hover,
.w-tabs.layout_accordion .w-tabs-section.active .w-tabs-section-header,
.no-touch .w-tags-item-link:hover,
.w-team-link .w-team-name,
.no-touch .w-clients .slick-prev:hover,
.no-touch .w-clients .slick-next:hover,
.woocommerce .products .product .button,
.no-touch .woocommerce .products .product .button.loading:hover,
.woocommerce .star-rating span:before,
.woocommerce-breadcrumb a,
.woocommerce div.product .woocommerce-tabs .tabs li.active,
.no-touch .woocommerce div.product .woocommerce-tabs .tabs li.active:hover,
.woocommerce .stars span a:after,
#subscription-toggle span.is-subscribed:before,
#favorite-toggle span.is-favorite:before {
	color: #6dbcdb;
	}
.l-submain.color_primary,
.w-actionbox.color_primary,
.g-btn.color_primary,
.g-btn.color_primary.outlined:before,
button,
input[type="submit"],
.no-touch .pagination .page-numbers:before,
.pagination .page-numbers.current,
.no-touch .w-iconbox.type_outlined .w-iconbox-icon:before,
.no-touch .w-iconbox.type_solid .w-iconbox-icon:before,
.w-iconbox.color_primary.type_solid .w-iconbox-icon,
.no-touch .w-filters-item:hover,
.w-icon.color_primary.type_solid .w-icon-link,
.w-pricing-item.type_featured .w-pricing-item-title,
.w-pricing-item.type_featured .w-pricing-item-price,
.w-progbar.color_primary .w-progbar-bar-h,
.no-touch .w-team-links,
.w-timeline-item:before,
.w-timeline.type_vertical .w-timeline-section:before,
.w-timeline-section-title:before,
.w-timeline-section.active .w-timeline-section-title:before,
.no-touch .w-toplink.active:hover,
.no-touch .tp-leftarrow.custom:before,
.no-touch .tp-rightarrow.custom:before,
.widget_nav_menu .menu-item.current-menu-item > a,
.no-touch .widget_nav_menu .menu-item.current-menu-item > a:hover,
p.demo_store,
.woocommerce .button.alt,
.woocommerce .button.checkout,
.no-touch .woocommerce .products .product .button:hover,
.no-touch .woocommerce-pagination a:hover,
.woocommerce-pagination span.current,
.woocommerce .onsale,
.widget_price_filter .ui-slider-range,
.widget_layered_nav ul li.chosen,
.widget_layered_nav_filters ul li a,
.no-touch .bbp-pagination-links a:hover,
.bbp-pagination-links span.current,
.no-touch span.bbp-topic-pagination a.page-numbers:hover,
.gform_wrapper .gf_progressbar_percentage,
.gform_wrapper .gform_page_footer .gform_next_button,
.gform_wrapper .chosen-container .chosen-results li.highlighted,
.smile-icon-timeline-wrap .timeline-separator-text .sep-text,
.smile-icon-timeline-wrap .timeline-wrapper .timeline-dot,
.smile-icon-timeline-wrap .timeline-feature-item .timeline-dot {
	background-color: #6dbcdb;
	}
.g-html blockquote,
.w-blog-entry.sticky,
.no-touch .w-clients-item-h:hover,
.w-filters-item.active,
.w-tabs-item.active,
.no-touch .w-tabs-item.active:hover,
.widget_nav_menu .menu-item.current-menu-item > a,
.fotorama__thumb-border,
.woocommerce div.product .woocommerce-tabs .tabs li.active,
.widget_layered_nav ul li.chosen,
.no-touch .bbp-pagination-links a:hover,
.bbp-pagination-links span.current,
.no-touch span.bbp-topic-pagination a.page-numbers:hover,
#bbp-user-navigation li.current {
	border-color: #6dbcdb;
	}
input[type="text"]:focus,
input[type="password"]:focus,
input[type="email"]:focus,
input[type="url"]:focus,
input[type="tel"]:focus,
input[type="number"]:focus,
input[type="date"]:focus,
input[type="search"]:focus,
textarea:focus,
select:focus,
#bbpress-forums div.bbp-the-content-wrapper textarea:focus {
	box-shadow: 0 0 0 2px #6dbcdb;
	}
.g-btn.color_primary.outlined,
.l-main .w-contacts-item i,
.w-iconbox.color_primary.type_outlined .w-iconbox-icon,
.no-touch .w-pagehead-nav-item:hover,
.no-touch .w-tags-item-link:hover,
.no-touch .w-testimonial:hover,
.w-timeline-item,
.w-timeline-section-title,
.no-touch .w-clients .slick-prev:hover,
.no-touch .w-clients .slick-next:hover,
.woocommerce .products .product .button,
.no-touch .woocommerce .products .product .button.loading:hover {
	box-shadow: 0 0 0 2px #6dbcdb inset;
	}

/* Secondary Color */
.no-touch a:hover,
.g-btn.color_secondary.outlined,
.no-touch .w-blog.type_masonry .w-blog-meta a:hover,
.no-touch .w-blognav-prev:hover .w-blognav-title,
.no-touch .w-blognav-next:hover .w-blognav-title,
.w-counter.color_secondary .w-counter-number,
.w-icon.color_secondary .w-icon-link,
.w-iconbox.color_secondary .w-iconbox-icon,
.no-touch .w-team-link:hover .w-team-name,
.no-touch .widget_archive ul li a:hover,
.no-touch .widget_categories ul li a:hover,
.no-touch .widget_tag_cloud .tagcloud a:hover,
.no-touch .woocommerce-breadcrumb a:hover,
.no-touch .widget_product_tag_cloud .tagcloud a:hover,
.no-touch .bbp_widget_login a.button.logout-link:hover {
	color: #fc4349;
	}
.l-submain.color_secondary,
.w-actionbox.color_secondary,
.g-btn.color_secondary,
.g-btn.color_secondary.outlined:before,
.w-icon.color_secondary.type_solid .w-icon-link,
.w-iconbox.color_secondary.type_solid .w-iconbox-icon,
.w-progbar.color_secondary .w-progbar-bar-h,
.no-touch .button:hover,
.no-touch input[type="submit"]:hover,
.no-touch .woocommerce .button:hover,
.no-touch .woocommerce .button.alt:hover,
.no-touch .woocommerce .button.checkout:hover,
.no-touch .woocommerce .shop_table.cart .product-remove a:hover,
.no-touch .widget_layered_nav_filters ul li a:hover,
.no-touch .bbp-row-actions #favorite-toggle a:hover,
.no-touch .bbp-row-actions #subscription-toggle a:hover {
	background-color: #fc4349;
	}
.g-btn.color_secondary.outlined,
.w-iconbox.color_secondary.type_outlined .w-iconbox-icon {
	box-shadow: 0 0 0 2px #fc4349 inset;
	}

/* Fade Elements Color */
.w-blog-meta,
.w-blog-meta a,
.w-icon.color_fade .w-icon-link,
.no-touch .w-icon.color_fade.type_solid .w-icon-link:hover,
.w-profile-link.for_logout,
.widget_tag_cloud .tagcloud a,
.woocommerce-breadcrumb,
.woocommerce .star-rating:before,
.woocommerce .stars span:after,
.woocommerce table.shop_table .product-remove a.remove,
.widget_product_tag_cloud .tagcloud a,
p.bbp-topic-meta,
.bbp_widget_login a.button.logout-link {
	color: #a4abb3;
	}
.w-shortblog-entry-date {
	box-shadow: 0 0 0 2px #a4abb3 inset;
	}

/*************************** ALTERNATE CONTENT ***************************/

/* Background Color */
.l-submain.color_alternate,
.color_alternate .w-blog.type_masonry .w-blog-entry-preview:before,
.color_alternate .w-filters-item.active,
.no-touch .color_alternate .w-filters-item.active:hover,
.color_alternate .w-tabs-item.active,
.no-touch .color_alternate .w-tabs-item.active:hover,
.color_alternate .w-timeline-item,
.color_alternate .w-timeline-section-title,
.color_alternate .w-timeline.type_vertical .w-timeline-section-content {
	background-color: #f2f4f5;
	}
.color_alternate .g-btn.color_contrast,
.no-touch .color_alternate .g-btn.color_contrast:hover,
.no-touch .color_alternate .g-btn.color_contrast.outlined:hover,
.color_alternate .w-icon.color_border.type_solid .w-icon-link {
	color: #f2f4f5;
	}

/* Alternate Background Color */
.color_alternate input[type="text"],
.color_alternate input[type="password"],
.color_alternate input[type="email"],
.color_alternate input[type="url"],
.color_alternate input[type="tel"],
.color_alternate input[type="number"],
.color_alternate input[type="date"],
.color_alternate input[type="search"],
.color_alternate textarea,
.color_alternate select,
.color_alternate .w-blog.imgpos_atleft .w-blog-entry-preview-icon,
.color_alternate .w-filters,
.color_alternate .w-icon.color_text.type_solid .w-icon-link,
.color_alternate .w-icon.color_fade.type_solid .w-icon-link,
.color_alternate .w-pricing-item-title,
.color_alternate .w-pricing-item-price,
.color_alternate .w-tabs-list {
	background-color: #fff;
	}

/* Border Color */
.color_alternate .w-blog-entry,
.color_alternate .w-bloglist,
.color_alternate .w-comments-item,
.color_alternate .w-pricing-item-h,
.color_alternate .w-tabs.layout_accordion,
.color_alternate .w-tabs.layout_accordion .w-tabs-section,
.color_alternate .w-timeline.type_vertical .w-timeline-section-content {
	border-color: #e3e6e8;
	}
.color_alternate .g-hr-h i,
.color_alternate .page-404 i,
.color_alternate .w-icon.color_border .w-icon-link {
	color: #e3e6e8;
	}
.color_alternate .g-hr-h:before,
.color_alternate .g-hr-h:after,
.color_alternate .g-btn.color_default,
.color_alternate .g-btn.color_default.outlined:before,
.color_alternate .w-icon.color_border.type_solid .w-icon-link,
.color_alternate .w-timeline-list:before {
	background-color: #e3e6e8;
	}
.color_alternate .g-btn.color_default.outlined,
.color_alternate .pagination .page-numbers,
.color_alternate .w-socials-item-link,
.color_alternate .w-tags-item-link,
.color_alternate .w-team-links-item,
.color_alternate .w-testimonial {
	box-shadow: 0 0 0 2px #e3e6e8 inset;
	}

/* Heading Color */
.color_alternate h1,
.color_alternate h2,
.color_alternate h3,
.color_alternate h4,
.color_alternate h5,
.color_alternate h6,
.color_alternate input[type="text"],
.color_alternate input[type="password"],
.color_alternate input[type="email"],
.color_alternate input[type="url"],
.color_alternate input[type="tel"],
.color_alternate input[type="number"],
.color_alternate input[type="date"],
.color_alternate textarea,
.color_alternate select,
.color_alternate .w-form-field > i,
.no-touch .color_alternate .g-btn.color_default:hover,
.no-touch .color_alternate .g-btn.color_default.outlined:hover,
.color_alternate .g-btn.color_contrast.outlined,
.color_alternate .w-blog-entry-title,
.color_alternate .w-counter-number,
.color_alternate .w-pricing-item-title,
.color_alternate .w-pricing-item-price {
	color: #292e33;
	}
.color_alternate .g-btn.color_contrast,
.color_alternate .g-btn.color_contrast.outlined:before {
	background-color: #292e33;
	}
.color_alternate .g-btn.color_contrast.outlined {
	box-shadow: 0 0 0 2px #292e33 inset;
	}

/* Text Color */
.l-submain.color_alternate,
.color_alternate .g-btn.color_default,
.color_alternate .g-btn.color_default.outlined,
.color_alternate .w-icon.color_text .w-icon-link {
	color: #5c6166;
	}

/* Primary Color */
.color_alternate a,
.color_alternate .g-btn.color_primary.outlined,
.no-touch .color_alternate .w-blog-entry-link:hover .w-blog-entry-title-h,
.no-touch .color_alternate .w-blog-entry-link:hover,
.color_alternate .l-main .w-contacts-item i,
.color_alternate .w-counter.color_primary .w-counter-number,
.color_alternate .w-filters-item.active,
.no-touch .color_alternate .w-filters-item.active:hover,
.color_alternate .w-icon.color_primary .w-icon-link,
.color_alternate .w-iconbox-icon,
.no-touch .color_alternate .w-iconbox-link:hover .w-iconbox-title,
.no-touch .color_alternate .w-pagehead-nav-item:hover,
.color_alternate .w-tabs-item.active,
.no-touch .color_alternate .w-tabs-item.active:hover,
.color_alternate .w-tabs.layout_accordion .w-tabs-section.active .w-tabs-section-header,
.no-touch .color_alternate .w-tags-item-link:hover,
.color_alternate .w-team-link .w-team-name {
	color: #6dbcdb;
	}
.color_alternate .g-btn.color_primary,
.color_alternate .g-btn.color_primary.outlined:before,
.color_alternate input[type="submit"],
.no-touch .color_alternate .pagination .page-numbers:before,
.color_alternate .pagination .page-numbers.current,
.no-touch .color_alternate .w-filters-item:hover,
.color_alternate .w-icon.color_primary.type_solid .w-icon-link,
.color_alternate .w-pricing-item.type_featured .w-pricing-item-title,
.color_alternate .w-pricing-item.type_featured .w-pricing-item-price,
.no-touch .color_alternate .w-team-links,
.color_alternate .w-timeline-item:before,
.color_alternate .w-timeline.type_vertical .w-timeline-section:before,
.color_alternate .w-timeline-section-title:before,
.color_alternate .w-timeline-section.active .w-timeline-section-title:before {
	background-color: #6dbcdb;
	}
.color_alternate .g-html blockquote,
.color_alternate .w-blog-entry.sticky,
.color_alternate .w-filters-item.active,
.color_alternate .w-tabs-item.active,
.no-touch .color_alternate .w-tabs-item.active:hover {
	border-color: #6dbcdb;
	}
.color_alternate input[type="text"]:focus,
.color_alternate input[type="password"]:focus,
.color_alternate input[type="email"]:focus,
.color_alternate input[type="url"]:focus,
.color_alternate input[type="tel"]:focus,
.color_alternate input[type="number"]:focus,
.color_alternate input[type="date"]:focus,
.color_alternate input[type="search"]:focus,
.color_alternate textarea:focus,
.color_alternate select:focus {
	box-shadow: 0 0 0 2px #6dbcdb;
	}
.color_alternate .g-btn.color_primary.outlined,
.color_alternate .l-main .w-contacts-item i,
.no-touch .color_alternate .w-pagehead-nav-item:hover,
.no-touch .color_alternate .w-tags-item-link:hover,
.no-touch .color_alternate .w-testimonial:hover,
.color_alternate .w-timeline-item:before,
.color_alternate .w-timeline-section-title:before {
	box-shadow: 0 0 0 2px #6dbcdb inset;
	}

/* Secondary Color */
.no-touch .color_alternate a:hover,
.no-touch .color_alternate a:active,
.color_alternate .g-btn.color_secondary.outlined,
.no-touch .color_alternate .w-blog.type_masonry .w-blog-meta a:hover,
.color_alternate .w-counter.color_secondary .w-counter-number,
.color_alternate .w-icon.color_secondary .w-icon-link,
.no-touch .color_alternate .w-team-link:hover .w-team-name {
	color: #fc4349;
	}
.color_alternate .g-btn.color_secondary,
.color_alternate .g-btn.color_secondary.outlined:before,
.color_alternate .w-icon.color_secondary.type_solid .w-icon-link {
	background-color: #fc4349;
	}
.color_alternate .g-btn.color_secondary.outlined {
	box-shadow: 0 0 0 2px #fc4349 inset;
	}

/* Fade Elements Color */
.color_alternate .w-blog-meta,
.color_alternate .w-blog-meta a,
.color_alternate .w-bloglist-entry-date,
.color_alternate .w-bloglist-entry-author,
.color_alternate .w-icon.color_fade .w-icon-link {
	color: #a4abb3;
	}
.color_alternate .w-shortblog-entry-date {
	box-shadow: 0 0 0 2px #a4abb3 inset;
	}

/*************************** SUBFOOTER ***************************/

/* Background Color */
.l-subfooter.at_top,
.l-subfooter.at_top #lang_sel ul ul,
.l-subfooter.at_top #lang_sel_click ul ul {
	background-color: #2c3e50;
	}
	
/* Alternate Background Color */
.l-subfooter.at_top input[type="text"],
.l-subfooter.at_top input[type="password"],
.l-subfooter.at_top input[type="email"],
.l-subfooter.at_top input[type="url"],
.l-subfooter.at_top input[type="tel"],
.l-subfooter.at_top input[type="number"],
.l-subfooter.at_top input[type="date"],
.l-subfooter.at_top input[type="search"],
.l-subfooter.at_top textarea,
.l-subfooter.at_top select,
.no-touch .l-subfooter.at_top #lang_sel a:hover,
.no-touch .l-subfooter.at_top #lang_sel_click a:hover {
	background-color: #384b5f;
	}

/* Border Color */
.l-subfooter.at_top,
.l-subfooter.at_top #wp-calendar thead th,
.l-subfooter.at_top #wp-calendar tbody td,
.l-subfooter.at_top #wp-calendar tfoot td,
.l-subfooter.at_top #lang_sel a,
.l-subfooter.at_top #lang_sel_click a,
.l-subfooter.at_top .widget_nav_menu .menu-item a {
	border-color: #384b5f;
	}
.l-subfooter.at_top .w-socials-item-link {
	box-shadow: 0 0 0 2px #384b5f inset;
	}

/* Text Color */
.l-subfooter.at_top,
.l-subfooter.at_top .w-socials-item-link {
	color: #939da3;
	}

/* Heading Color */
.l-subfooter.at_top h1,
.l-subfooter.at_top h2,
.l-subfooter.at_top h3,
.l-subfooter.at_top h4,
.l-subfooter.at_top h5,
.l-subfooter.at_top h6,
.l-subfooter.at_top .w-form-field > i,
.l-subfooter.at_top input[type="text"],
.l-subfooter.at_top input[type="password"],
.l-subfooter.at_top input[type="email"],
.l-subfooter.at_top input[type="url"],
.l-subfooter.at_top input[type="tel"],
.l-subfooter.at_top input[type="number"],
.l-subfooter.at_top input[type="date"],
.l-subfooter.at_top input[type="search"],
.l-subfooter.at_top textarea,
.l-subfooter.at_top select {
	color: #d3d8db;
	}

/* Link Color */
.l-subfooter.at_top a,
.l-subfooter.at_top .widget_tag_cloud .tagcloud a,
.l-subfooter.at_top .widget_product_tag_cloud .tagcloud a {
	color: #6dbcdb;
	}

/* Link Hover Color */
.no-touch .l-subfooter.at_top a:hover,
.no-touch .l-subfooter.at_top .widget_tag_cloud .tagcloud a:hover,
.no-touch .l-subfooter.at_top .widget_product_tag_cloud .tagcloud a:hover {
	color: #fc4349;
	}

/*************************** FOOTER ***************************/

/* Background Color */
.l-subfooter.at_bottom {
	background-color: #ffffff;
	}

/* Text Color */
.l-subfooter.at_bottom {
	color: #939da3;
	}

/* Link Color */
.l-subfooter.at_bottom a {
	color: #d3d8db;
	}

/* Link Hover Color */
.no-touch .l-subfooter.at_bottom a:hover {
	color: #fc4349;
	}
</style>
</head>


<link rel='stylesheet' id='mesh-nightmode-css'  href='https://www.novelupdates.com/wp-content/themes/ndupdates-child/mystyles/Default.css?ver=2.1.4' type='text/css' media='all' /><body class="page-template page-template-pg-seriesfinder page-template-pg-seriesfinder-php page page-id-3370 l-body us-theme_nd_4-4-8">


<!-- CANVAS -->

<div class="l-canvas  col_contside">

	<!-- HEADER -->
	<div class="l-header">

		<div class="l-subheader at_top" style="line-height: 36px; ">
			<div class="l-subheader-h i-cf">
			</div>
		</div>
				<div class="l-subheader at_middle"  style="line-height: 60px;">
			<div class="l-subheader-h i-widgets i-cf">

				<div class="w-logo  with_title">
					<a class="w-logo-link" href="https://www.novelupdates.com/">
												<span class="w-logo-title">Novel Updates</span>
					</a>
				</div>

				                
                   
			    <script>$( document ).ready(function() {$.get( "https://forum.novelupdates.com/nufonline_guest.php" );});</script>                
                <span class="menu_username_right" id="menu_username_right"><span class="menu_right_icons notloggedin"><a class="menu_right_links" href="/register/">Register</a></span><span class="menu_right_icons notloggedin"><a class="menu_right_links" href="/login/">Login</a></span><span id="gs_menu" class="menu_right_icons search guest" onclick="toggleUserMenu();"><i id="gs_menu" class="fa fa-user-circle menu" aria-hidden="true" dp="yes"></i><i id="gs_menu" class="fa fa-angle-down menu" aria-hidden="true" dp="yes"></i></span><div title="Guest Menu" class="lo_main_themain guest" style="display: inline; position:relative;">
						<div class="lo_main">
						<div class="arrow-up_toc_username"></div>
						<div class="lo_box">
						<span><span class="user_menu_links atpmn" id="user_menu_links">Theme <select class="wi_themes" id="wi_themes" dp="yes" onchange="toggleTheme(this.value);"><option value="Default" dp="yes">Default</option><option value="Dark" dp="yes">Dark</option><option value="Dark - Crisp" dp="yes">Dark Crisp</option></select></span></span>
						</div></div>
						</div><span class="menu_right_icons search" title="Search" onclick="show_search_bar(this);"><i class="fa fa-search menu"></i></span><span class="menu_right_icons mobile" style="display:none;" onclick="toggleMobileMenu();"><i class="fa fa-bars menu" aria-hidden="true"></i></span><span class="nu_menu_search"><div class="ab_title" dp="yes">Search</div><form id="search_nu_novel" action="https://www.novelupdates.com/"><span class="w-search-input-h">
						<div class="srh_menu">
						<input id="searchme-rl newmenu" class="nusearchnovel" autocomplete="off" onkeyup="showSearch(this.value)" type="text" value="" name="s" placeholder="Search...">
						</div>
						<input type="hidden" name="post_type" value="seriesplans">
							
						<div class="w-search-submit">
						<span class="search_menu_right"><i class="fa fa-search btn" onclick="search_btn_click();"></i></span>
						
						<div class ="search_type">			
						<span style="padding-right: 35px;">
						<span onclick="g_checkmark(this);" id="rl_checkbox_main" class="s_series islbl" value="1"><i dp="yes" class="fa fa-check" value="1" aria-hidden="true"></i></span>
						<span onclick="g_check_lbl(this);" style="position: relative; top: -1px ;padding-left: 2px; cursor:pointer;">Series</span></span>
					
						<span>
						<span onclick="g_checkmark(this);" id="rl_checkbox_main" class="s_users islbl" value="0">
						<i class="fa fa-check icon-invisible" aria-hidden="true" value="0"></i></span>
						<span onclick="g_check_lbl(this);" style="position: relative; top: -1px ;padding-left: 2px; cursor:pointer;">Group</span></span>
						</div>
						
						<div style="font-size: 13px; padding-right:5px; position: relative; top: 15px;"><a href="https://www.novelupdates.com/series-finder/">Series Finder</a></div>
						</div>
							
						</span><div class="livesearchresult" style="display: none; z-index: 2;"></div></form></span></span>                
            
				<!-- NAV -->
				<nav class="w-nav layout_hor animation_">
					<ul class="w-nav-list level_1 ">

                        <div class="menu_items_display" style="display:none;"><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item menu-item-73"><a class="w-nav-anchor level_1" href="/login/"><span class="w-nav-title">Login</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item menu-item-73"><a class="w-nav-anchor level_1" href="/register/"><span class="w-nav-title">Register</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="http://forum.novelupdates.com/"><span class="w-nav-title">Forum</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="//www.novelupdates.com/random-novel/"><span class="w-nav-title">Random Novel</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="//www.novelupdates.com/series-finder/"><span class="w-nav-title">Series Finder</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="//www.novelupdates.com/novelslisting/"><span class="w-nav-title">Series Listing</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="//www.novelupdates.com/series-ranking/"><span class="w-nav-title">Series Ranking</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="//www.novelupdates.com/latest-series/"><span class="w-nav-title">Latest Series</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><a class="w-nav-anchor level_1" href="https://www.novelupdates.com/recommendation-lists/"><span class="w-nav-title">Rec Lists</span><span class="w-nav-arrow"></span></a></li><li id="menu-item-73" class="menu-item menu-item-type-post_type menu-item-object-page w-nav-item level_1 menu-item-73"><span class="w-an_cus">Theme</span> <span class="c_theme"><select class="wi_themes" id="wi_themes" dp="yes" onchange="toggleTheme(this.value);"><option value="Default" dp="yes">Default</option><option value="Dark" dp="yes">Dark</option><option value="Dark - Crisp" dp="yes">Dark Crisp</option></select></span></li></div>                    </ul>
				</nav><!-- /NAV -->
				
							</div>
		</div>


	</div>
	<!-- /HEADER -->
    <div class="ol_account_main" style="position:relative; display: none;">
			<div class="ol_account"><span class="ol_title">Your changes has been saved.</span></div></div>	<!-- MAIN -->
	<div class="l-main">




<link rel="stylesheet" href="https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/gadient/style.css"><link rel="stylesheet" href="https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/genre_search.css"><link rel="stylesheet" href="https://www.novelupdates.com/wp-content/themes/ndupdates-child/js/chosen.min.css"><link rel="stylesheet" href="//code.jquery.com/ui/1.11.4/themes/smoothness/jquery-ui.css"><script src="//ajax.googleapis.com/ajax/libs/jqueryui/1.11.1/jquery-ui.min.js"></script>

<style>
.select2
{
	width:100% !important;
}
input.select2-search__field
{
	box-shadow:none !important;
}
.select2-container--bootstrap .select2-selection--multiple .select2-selection__choice
{
	margin: 10px 0 0 6px !important;
}
.select2-container--bootstrap .select2-selection--multiple .select2-search--inline .select2-search__field
{
	margin-top: 3px !important;
}
</style> 

	<div class="l-submain">
		<div class="l-submain-h g-html i-cf other">
			<div class="l-content">
				<div class="w-blog post-3370 page type-page status-publish hentry">
              

 	 	 
					<div class="w-blog-content other">

     
<h3 class="mypage bc top">Series Finder</h3>
<div class="breadcrumb_nu">
		<div id="Breadcrumb_Top" style="max-width:1420px; margin:0px auto;">
				<dl>
					<dd><a href="//www.novelupdates.com/" title="Home">Home</a>&nbsp;&gt;&nbsp;</dd>
          			<dd>Series Finder</dd>
				</dl>
	
		</div>	
</div>
<div class="ricon sf bc" style="margin-top: -72px;""> <i class="fa fa-cogs sfinder" style="margin-right: -2px;" aria-hidden="true"></i> <span class="ftext sf"> [ <a href="#" class="rlhidesme" datahide="no" onclick="hiderk(this)">Filters</a> ] </span></div>


<form style="padding-bottom: 20px;" id="rankfilter" name="rankfilter" action="/series-finder/?sf=1&sort=sdate&order=desc" method="post">
<div class="sf_helpinfo">
<ul style="margin-bottom: 5px;">
<li><b>Chapters</b> - This is the number of releases (chapters).</li>
	<ul style="margin-bottom: 5px;">
	<li>Min - Minimum amount of chapters.</li>
	<li>Max - Maximum amount of chapters.</li>
	</ul>
<li><b>Release Frequency</b> - The release frequency of a novel. Higher frequency means the novel is updated more often.</li>
	<ul style="margin-bottom: 5px;">
	<li>Min - Minimum release frequency. Lower frequency = slower updates. </li>
	<li>Max - Maximum release frequency. Higher frequency = quicker updates.</li>
	</ul>
<li><b>Rating</b> - Novel rating on a scale of 1 to 5.</li>
	<ul style="margin-bottom: 5px;">
	<li>Min - Minimum rating (1 to 5)</li>
	<li>Max - Maximum rating (1 to 5)</li>
	</ul>
<li><b>Number of Ratings</b> - The amount of ratings for a novel.</li>
	<ul style="margin-bottom: 5px;">
	<li>Min - Minimum number of ratings.</li>
	<li>Max - Maximum number of ratings.</li>
	</ul>
<li><b>Readers</b> - The number of readers a novel has.</li>
	<ul style="margin-bottom: 5px;">
	<li>Min - Minimum readers.</li>
	<li>Max - Maximum readers.</li>
	</ul>
<li><b>Last Release Date</b> - The last release date of the latest novel.</li> 
	<ul style="margin-bottom: 5px;">
	<li>Min - Minimum last release date.</li>
	<li>Max - Maximum last release date.</li>
	</ul>
<li><b>Genre</b> - The genres of the novel.</li>
	<ul style="margin-bottom: 5px;">
	<li>Include - Click a genre once to include it in your search.</li>
	<li>Exclude - Click a genre twice to exclude it in your search.</li>
	</ul>
<li><b>Tags</b> - The tags of a novel.</li>
	<ul style="margin-bottom: 5px;">
	<li>OR - Uses the 'OR' condition when searching for tags.</li>
	<li>AND - Uses the 'AND' condition when searching for tags.</li>
	</ul>
<li><b>Hide Series from... (Members Only)</b> - Prevents novels from showing up from the selected list. This is for members using the reading list feature.</li> 
</ul>
</div>

<div class="rankfl">
<span class="rktext sf bc" style="padding-top: 15px;">Current Filters - <span class="clist list"><b>Order By:</b></span> Last Updated, Descending</span>
<span class="sftitle sfpad new">Novel Type <i onclick="sf_getnovelinfo('noveltype', 'Novel Type');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span>   

<div class="g-cols wpb_row offset_default"><div class="one-third" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			
<ul class="genrematch rank"><li><a class="typerank" href="javascript:void(0);" genreid="2443" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Light Novel</a></li></ul></div> 
	</div> </div><div class="one-third" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			

<ul class="genrematch rank"><li></i><a class="typerank" href="javascript:void(0);" genreid="26874" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Published Novel</a></li></ul></div> 
	</div> </div><div class="one-third" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			

<ul class="genrematch rank"><li></i><a class="typerank" href="javascript:void(0);" genreid="2444" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Web Novel</a></li></ul></div> 
	</div> </div></div>
                    
                  
<span class="sftitle sfpad new">Language <i onclick="sf_getnovelinfo('language', 'Language');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span>   
 

<div class="g-cols wpb_row offset_default"><div class="one-third" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			
<ul class="genrematch rank"><li><a class="langrank" href="javascript:void(0);" genreid="495" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Chinese</a></li><li><a class="langrank" href="javascript:void(0);" genreid="9181" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Filipino</a></li><li><a class="langrank" href="javascript:void(0);" genreid="9179" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Indonesian</a></li></ul></div> 
	</div> </div><div class="one-third" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			

<ul class="genrematch rank"><li></i><a class="langrank" href="javascript:void(0);" genreid="496" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Japanese</a></li><li></i><a class="langrank" href="javascript:void(0);" genreid="18657" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Khmer</a></li><li></i><a class="langrank" href="javascript:void(0);" genreid="497" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Korean</a></li></ul></div> 
	</div> </div><div class="one-third" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			

<ul class="genrematch rank"><li></i><a class="langrank" href="javascript:void(0);" genreid="9183" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Malaysian</a></li><li></i><a class="langrank" href="javascript:void(0);" genreid="9954" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Thai</a></li><li></i><a class="langrank" href="javascript:void(0);" genreid="9177" chkme="off" onclick="cOptionRank(this);"><i class="fa fa-square-o"></i>Vietnamese</a></li></ul></div> 
	</div> <div class="g-hr type_invisible no_icon">
						<span class="g-hr-h">
							<i class="fa fa-"></i>
						</span>
					</div></div></div>
                    
 


<span class="sftitle sfpad new">Chapters [<select name="releases_mm" id="r_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>] <i onclick="sf_getnovelinfo('releases', 'Chapters');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
<input autocomplete="off" class="inputrank finder" type="text" name="rk_releases" id="rk_releases" value="">

<span class="sftitle sfpad new">Release Frequency [<select name="releases_mm" id="rf_mm" class="select sf">
						<option value="min">min</option>
						<option value="max" selected>max</option>
					</select>]<i onclick="sf_getnovelinfo('frequency', 'Release Frequency');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
<input autocomplete="off" class="inputrank finder" type="text" name="rk_rfreq" id="rk_rfreq" value=""> 

<span class="sftitle sfpad new">Reviews [<select name="releases_mm" id="rvc_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>]</span> 
<input autocomplete="off" class="inputrank finder" type="text" name="rk_rreviews" id="rk_rreviews" value=""> 

<span class="sftitle sfpad new">Rating  [<select name="releases_mm" id="rt_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>]<i onclick="sf_getnovelinfo('rating', 'Ratings');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
<input autocomplete="off" class="inputrank finder" type="text" name="rk_sr" id="rk_sr" value=""> 
<span class="sftitle sfpad new">Number of Ratings [<select name="releases_mm" id="rtc_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>] <i onclick="sf_getnovelinfo('numberratings', 'Number of Ratings');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
<input autocomplete="off" class="inputrank finder" type="text" name="rk_rcount" id="rk_rcount" value="">  
<span class="sftitle sfpad new">Readers [<select name="releases_mm" id="rct_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>] <i onclick="sf_getnovelinfo('readers', 'Readers');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
<input autocomplete="off" class="inputrank finder" type="text" name="rk_sread" id="rk_sread" value="">  


<span class="sftitle sfpad new">First Release Date  [<select name="releases_mm" id="dtf_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>]</span> 
 <p>
 <input autocomplete="off" class="inputrank sf" type="text" name="ardate_first" id="ardate_first" value="">
</p>

<span class="sftitle sfpad new">Last Release Date  [<select name="releases_mm" id="dt_mm" class="select sf">
						<option value="min" selected>min</option>
						<option value="max">max</option>
					</select>] <i onclick="sf_getnovelinfo('lastdate', 'Last Release Date');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
 <p>
 <input autocomplete="off" class="inputrank sf" type="text" name="ardate" id="ardate" value="">
</p>


<span class="sftitle sfpad new">Genre [<select name="releases_mm" id="gi_mm" class="select sf">
						<option value="and" selected>AND</option>
						<option value="or">OR</option>
					</select>]<i onclick="sf_getnovelinfo('genre', 'Genre');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 
<div style="font-size:12px;">Click once to include genre. Click the same genre twice to exclude.</div>
<div class="g-cols wpb_row offset_default"><div class="one-quarter" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			
<ul class="genrematch rank"><li><a class="genreme" href="javascript:void(0);" genreid="8" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Action</a></li><li><a class="genreme" href="javascript:void(0);" genreid="280" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Adult</a></li><li><a class="genreme" href="javascript:void(0);" genreid="13" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Adventure</a></li><li><a class="genreme" href="javascript:void(0);" genreid="17" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Comedy</a></li><li><a class="genreme" href="javascript:void(0);" genreid="9" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Drama</a></li><li><a class="genreme" href="javascript:void(0);" genreid="292" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Ecchi</a></li><li><a class="genreme" href="javascript:void(0);" genreid="5" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Fantasy</a></li><li><a class="genreme" href="javascript:void(0);" genreid="168" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Gender Bender</a></li><li><a class="genreme" href="javascript:void(0);" genreid="3" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Harem</a></li></ul></div> 
	</div> </div><div class="one-quarter" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			

<ul class="genrematch rank"><li></i><a class="genreme" href="javascript:void(0);" genreid="330" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Historical</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="343" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Horror</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="324" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Josei</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="14" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Martial Arts</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="4" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Mature</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="10" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Mecha</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="245" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Mystery</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="486" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Psychological</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="15" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Romance</a></li></ul></div> 
	</div> </div><div class="one-quarter" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			
<ul class="genrematch rank"><li></i><a class="genreme" href="javascript:void(0);" genreid="6" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>School Life</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="11" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Sci-fi</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="18" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Seinen</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="157" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Shoujo</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="851" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Shoujo Ai</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="12" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Shounen</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="1692" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Shounen Ai</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="7" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Slice of Life</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="281" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Smut</a></li></ul></div> 
	</div> 
    
  </div>
            
            
       <div class="one-quarter" style="margin-bottom: 0px;">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			
<ul class="genrematch rank"><li></i><a class="genreme" href="javascript:void(0);" genreid="1357" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Sports</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="16" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Supernatural</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="132" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Tragedy</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="479" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Wuxia</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="480" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Xianxia</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="3954" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Xuanhuan</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="560" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Yaoi</a></li><li></i><a class="genreme" href="javascript:void(0);" genreid="922" chkme="off" onclick="cRKGenre(this);"><i class="fa fa-square-o"></i>Yuri</a></li></ul></div> 
	</div> </div>     
 
                    </div>
				<div class="tagfix" style="top:0px;">           
<span class="sftitle sfpad new"> Tags [<select name="releases_mm" id="tgi_mm" class="select sf">
						<option value="and">AND</option>
						<option value="or" selected>OR</option>
					</select>]<i onclick="sf_getnovelinfo('tags', 'Tags');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 

<div class="extrasf"></div>  
<div class="g-cols wpb_row offset_default">
<div class="one-half">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			 <select multiple id="tags_include" class="chzn-select" data-placeholder="Include..." name="tagsinclude&#091;&#093;"><option value="16185">Abandoned Children</option><option value="4859">Ability Steal</option><option value="1248">Absent Parents</option><option value="11325">Abusive Characters</option><option value="4885">Academy</option><option value="475">Accelerated Growth</option><option value="13403">Acting</option><option value="4976">Adapted from Manga</option><option value="6280">Adapted from Manhua</option><option value="269">Adapted to Anime</option><option value="2999">Adapted to Drama</option><option value="928">Adapted to Drama CD</option><option value="891">Adapted to Game</option><option value="270">Adapted to Manga</option><option value="182">Adapted to Manhua</option><option value="2721">Adapted to Manhwa</option><option value="1133">Adapted to Movie</option><option value="1334">Adapted to Visual Novel</option><option value="858">Adopted Children</option><option value="603">Adopted Protagonist</option><option value="4974">Adultery</option><option value="3108">Adventurers</option><option value="3802">Affair</option><option value="357">Age Progression</option><option value="216">Age Regression</option><option value="13720">Aggressive Characters</option><option value="234">Alchemy</option><option value="621">Aliens</option><option value="598">All-Girls School</option><option value="7909">Alternate World</option><option value="625">Amnesia</option><option value="774">Amusement Park</option><option value="4594">Anal</option><option value="476">Ancient China</option><option value="2662">Ancient Times</option><option value="13263">Androgynous Characters</option><option value="892">Androids</option><option value="1327">Angels</option><option value="255">Animal Characteristics</option><option value="7323">Animal Rearing</option><option value="724">Anti-Magic</option><option value="1732">Anti-social Protagonist</option><option value="4297">Antihero Protagonist</option><option value="203">Antique Shop</option><option value="513">Apartment Life</option><option value="12913">Apathetic Protagonist</option><option value="217">Apocalypse</option><option value="721">Appearance Changes</option><option value="430">Appearance Different from Actual Age</option><option value="697">Archery</option><option value="6442">Aristocracy</option><option value="9591">Arms Dealers</option><option value="306">Army</option><option value="9522">Army Building</option><option value="2676">Arranged Marriage</option><option value="13427">Arrogant Characters</option><option value="9851">Artifact Crafting</option><option value="4699">Artifacts</option><option value="141">Artificial Intelligence</option><option value="271">Artists</option><option value="81">Assassins</option><option value="8965">Astrologers</option><option value="9776">Autism</option><option value="11491">Automatons</option><option value="12917">Average-looking Protagonist</option><option value="1134">Award-winning Work</option><option value="13016">Awkward Protagonist</option><option value="4975">Bands</option><option value="1346">Based on a Movie</option><option value="1632">Based on a Song</option><option value="6572">Based on a TV Show</option><option value="1442">Based on a Video Game</option><option value="1447">Based on a Visual Novel</option><option value="1230">Based on an Anime</option><option value="825">Battle Academy</option><option value="6308">Battle Competition</option><option value="10668">BDSM</option><option value="5363">Beast Companions</option><option value="5406">Beastkin</option><option value="116">Beasts</option><option value="111">Beautiful Female Lead</option><option value="15625">Bestiality</option><option value="256">Betrayal</option><option value="2573">Bickering Couple</option><option value="13371">Biochip</option><option value="6359">Bisexual Protagonist</option><option value="2943">Black Belly</option><option value="916">Blackmail</option><option value="2734">Blacksmith</option><option value="9280">Blind Dates</option><option value="2347">Blind Protagonist</option><option value="9492">Blood Manipulation</option><option value="5824">Bloodlines</option><option value="2642">Body Swap</option><option value="9907">Body Tempering</option><option value="568">Body-double</option><option value="224">Bodyguards</option><option value="1288">Books</option><option value="13805">Bookworm</option><option value="325">Boss-Subordinate Relationship</option><option value="1710">Brainwashing</option><option value="5607">Breast Fetish</option><option value="9410">Broken Engagement</option><option value="604">Brother Complex</option><option value="8185">Brotherhood</option><option value="3825">Buddhism</option><option value="123">Bullying</option><option value="782">Business Management</option><option value="11786">Businessmen</option><option value="1597">Butlers</option><option value="12910">Calm Protagonist</option><option value="4248">Cannibalism</option><option value="2774">Card Games</option><option value="11345">Carefree Protagonist</option><option value="12951">Caring Protagonist</option><option value="12914">Cautious Protagonist</option><option value="2304">Celebrities</option><option value="257">Character Growth</option><option value="14493">Charismatic Protagonist</option><option value="14985">Charming Protagonist</option><option value="1942">Chat Rooms</option><option value="5708">Cheats</option><option value="1832">Chefs</option><option value="1368">Child Abuse</option><option value="3470">Child Protagonist</option><option value="1391">Childcare</option><option value="85">Childhood Friends</option><option value="917">Childhood Love</option><option value="691">Childhood Promise</option><option value="14986">Childish Protagonist</option><option value="169">Chuunibyou</option><option value="4423">Clan Building</option><option value="7777">Classic</option><option value="14587">Clever Protagonist</option><option value="2356">Clingy Lover</option><option value="2274">Clones</option><option value="611">Clubs</option><option value="14989">Clumsy Love Interests</option><option value="7951">Co-Workers</option><option value="710">Cohabitation</option><option value="14606">Cold Love Interests</option><option value="14605">Cold Protagonist</option><option value="8656">Collection of Short Stories</option><option value="7776">College/University</option><option value="1675">Coma</option><option value="124">Comedic Undertone</option><option value="282">Coming of Age</option><option value="4822">Complex Family Relationships</option><option value="7117">Conditional Power</option><option value="14673">Confident Protagonist</option><option value="1404">Confinement</option><option value="3241">Conflicting Loyalties</option><option value="859">Contracts</option><option value="2506">Cooking</option><option value="352">Corruption</option><option value="9926">Cosmic Wars</option><option value="7844">Cosplay</option><option value="592">Couple Growth</option><option value="5792">Court Official</option><option value="836">Cousins</option><option value="3752">Cowardly Protagonist</option><option value="5944">Crafting</option><option value="2969">Crime</option><option value="443">Criminals</option><option value="769">Cross-dressing</option><option value="1795">Crossover</option><option value="12956">Cruel Characters</option><option value="19556">Cryostasis</option><option value="117">Cultivation</option><option value="4596">Cunnilingus</option><option value="14593">Cunning Protagonist</option><option value="14990">Curious Protagonist</option><option value="579">Curses</option><option value="16457">Cute Children</option><option value="14987">Cute Protagonist</option><option value="5525">Cute Story</option><option value="8248">Dancers</option><option value="10343">Dao Companion</option><option value="6242">Dao Comprehension</option><option value="3824">Daoism</option><option value="4599">Dark</option><option value="2215">Dead Protagonist</option><option value="2142">Death</option><option value="1615">Death of Loved Ones</option><option value="5333">Debts</option><option value="749">Delinquents</option><option value="1362">Delusions</option><option value="5526">Demi-Humans</option><option value="2386">Demon Lord</option><option value="21268">Demonic Cultivation Technique</option><option value="86">Demons</option><option value="14541">Dense Protagonist</option><option value="12964">Depictions of Cruelty</option><option value="1946">Depression</option><option value="3135">Destiny</option><option value="1200">Detectives</option><option value="12959">Determined Protagonist</option><option value="15456">Devoted Love Interests</option><option value="2243">Different Social Status</option><option value="19241">Disabilities</option><option value="258">Discrimination</option><option value="13407">Disfigurement</option><option value="2343">Dishonest Protagonist</option><option value="7949">Distrustful Protagonist</option><option value="2019">Divination</option><option value="4706">Divine Protection</option><option value="2305">Divorce</option><option value="2876">Doctors</option><option value="16412">Dolls/Puppets</option><option value="2663">Domestic Affairs</option><option value="15026">Doting Love Interests</option><option value="15025">Doting Older Siblings</option><option value="14674">Doting Parents</option><option value="509">Dragon Riders</option><option value="897">Dragon Slayers</option><option value="72">Dragons</option><option value="1406">Dreams</option><option value="311">Drugs</option><option value="8126">Druids</option><option value="174">Dungeon Master</option><option value="175">Dungeons</option><option value="3569">Dwarfs</option><option value="1283">Dystopia</option><option value="15108">e-Sports</option><option value="7711">Early Romance</option><option value="1661">Earth Invasion</option><option value="5431">Easy Going Life</option><option value="784">Economics</option><option value="1878">Editors</option><option value="3796">Eidetic Memory</option><option value="13177">Elderly Protagonist</option><option value="9855">Elemental Magic</option><option value="165">Elves</option><option value="13305">Emotionally Weak Protagonist</option><option value="3570">Empires</option><option value="840">Enemies Become Allies</option><option value="113">Enemies Become Lovers</option><option value="263">Engagement</option><option value="2735">Engineer</option><option value="4480">Enlightenment</option><option value="1798">Episodic</option><option value="10593">Eunuch</option><option value="7778">European Ambience</option><option value="10172">Evil Gods</option><option value="5411">Evil Organizations</option><option value="12874">Evil Protagonist</option><option value="7087">Evil Religions</option><option value="2049">Evolution</option><option value="6365">Exhibitionism</option><option value="711">Exorcism</option><option value="200">Eye Powers</option><option value="525">Fairies</option><option value="134">Fallen Angels</option><option value="734">Fallen Nobility</option><option value="12009">Familial Love</option><option value="2152">Familiars</option><option value="641">Family</option><option value="2013">Family Business</option><option value="8664">Family Conflict</option><option value="1833">Famous Parents</option><option value="14555">Famous Protagonist</option><option value="5399">Fanaticism</option><option value="5691">Fanfiction</option><option value="771">Fantasy Creatures</option><option value="99">Fantasy World</option><option value="3068">Farming</option><option value="5068">Fast Cultivation</option><option value="7599">Fast Learner</option><option value="13721">Fat Protagonist</option><option value="15974">Fat to Fit</option><option value="2574">Fated Lovers</option><option value="14675">Fearless Protagonist</option><option value="3953">Fellatio</option><option value="5262">Female Master</option><option value="2879">Female Protagonist</option><option value="4178">Female to Male</option><option value="10919">Feng Shui</option><option value="9340">Firearms</option><option value="1569">First Love</option><option value="4880">First-time Intercourse</option><option value="6061">Flashbacks</option><option value="17020">Fleet Battles</option><option value="923">Folklore</option><option value="1353">Forced into a Relationship</option><option value="606">Forced Living Arrangements</option><option value="3052">Forced Marriage</option><option value="2367">Forgetful Protagonist</option><option value="7405">Former Hero</option><option value="6222">Fox Spirits</option><option value="1576">Friends Become Enemies</option><option value="1476">Friendship</option><option value="797">Fujoshi</option><option value="3468">Futanari</option><option value="2616">Futuristic Setting</option><option value="7643">Galge</option><option value="9425">Gambling</option><option value="93">Game Elements</option><option value="16272">Game Ranking System</option><option value="225">Gamers</option><option value="313">Gangs</option><option value="5382">Gate to Another World</option><option value="9295">Genderless Protagonist</option><option value="13888">Generals</option><option value="17629">Genetic Modifications</option><option value="2014">Genies</option><option value="14581">Genius Protagonist</option><option value="515">Ghosts</option><option value="2044">Gladiators</option><option value="15547">Glasses-wearing Love Interests</option><option value="15548">Glasses-wearing Protagonist</option><option value="11494">Goblins</option><option value="6683">God Protagonist</option><option value="1354">God-human Relationship</option><option value="1355">Goddesses</option><option value="73">Godly Powers</option><option value="177">Gods</option><option value="4437">Golems</option><option value="455">Gore</option><option value="1736">Grave Keepers</option><option value="807">Grinding</option><option value="2005">Guardian Relationship</option><option value="438">Guilds</option><option value="792">Gunfighters</option><option value="1753">Hackers</option><option value="9445">Half-human Protagonist</option><option value="9873">Handjob</option><option value="192">Handsome Male Lead</option><option value="14431">Hard-Working Protagonist</option><option value="14992">Harem-seeking Protagonist</option><option value="8149">Harsh Training</option><option value="2307">Hated Protagonist</option><option value="12141">Healers</option><option value="5267">Heartwarming</option><option value="242">Heaven</option><option value="4770">Heavenly Tribulation</option><option value="4720">Hell</option><option value="15545">Helpful Protagonist</option><option value="808">Herbalist</option><option value="100">Heroes</option><option value="2111">Heterochromia</option><option value="7227">Hidden Abilities</option><option value="6384">Hiding True Abilities</option><option value="10134">Hiding True Identity</option><option value="765">Hikikomori</option><option value="2741">Homunculus</option><option value="1772">Honest Protagonist</option><option value="1387">Hospital</option><option value="13502">Hot-blooded Protagonist</option><option value="12338">Human Experimentation</option><option value="1687">Human Weapon</option><option value="259">Human-Nonhuman Relationship</option><option value="17218">Humanoid Protagonist</option><option value="494">Hunters</option><option value="1711">Hypnotism</option><option value="10075">Identity Crisis</option><option value="3555">Imaginary Friend</option><option value="233">Immortals</option><option value="9033">Imperial Harem</option><option value="626">Incest</option><option value="3398">Incubus</option><option value="2319">Indecisive Protagonist</option><option value="12135">Industrialization</option><option value="506">Inferiority Complex</option><option value="1400">Inheritance</option><option value="5132">Inscriptions</option><option value="550">Insects</option><option value="4539">Interconnected Storylines</option><option value="11229">Interdimensional Travel</option><option value="14993">Introverted Protagonist</option><option value="9064">Investigations</option><option value="3978">Invisibility</option><option value="3687">Jack of All Trades</option><option value="1864">Jealousy</option><option value="6665">Jiangshi</option><option value="7086">Jobless Class</option><option value="5300">JSDF</option><option value="1250">Kidnappings</option><option value="13901">Kind Love Interests</option><option value="3869">Kingdom Building</option><option value="1904">Kingdoms</option><option value="137">Knights</option><option value="571">Kuudere</option><option value="5086">Lack of Common Sense</option><option value="5426">Language Barrier</option><option value="5920">Late Romance</option><option value="4551">Lawyers</option><option value="338">Lazy Protagonist</option><option value="6706">Leadership</option><option value="3430">Legends</option><option value="7443">Level System</option><option value="1402">Library</option><option value="2222">Limited Lifespan</option><option value="3154">Living Abroad</option><option value="1437">Living Alone</option><option value="3715">Loli</option><option value="865">Loneliness</option><option value="651">Loner Protagonist</option><option value="2891">Long Separations</option><option value="1859">Long-distance Relationship</option><option value="1696">Lost Civilizations</option><option value="7151">Lottery</option><option value="1561">Love at First Sight</option><option value="19605">Love Interest Falls in Love First</option><option value="627">Love Rivals</option><option value="607">Love Triangles</option><option value="1616">Lovers Reunited</option><option value="9161">Low-key Protagonist</option><option value="12035">Loyal Subordinates</option><option value="6526">Lucky Protagonist</option><option value="60">Magic</option><option value="6095">Magic Beasts</option><option value="3038">Magic Formations</option><option value="516">Magical Girls</option><option value="3460">Magical Space</option><option value="14900">Magical Technology</option><option value="336">Maids</option><option value="3257">Male Protagonist</option><option value="171">Male to Female</option><option value="557">Male Yandere</option><option value="2604">Management</option><option value="2781">Mangaka</option><option value="12963">Manipulative Characters</option><option value="2084">Manly Gay Couple</option><option value="642">Marriage</option><option value="706">Marriage of Convenience</option><option value="10145">Martial Spirits</option><option value="19604">Masochistic Characters</option><option value="667">Master-Disciple Relationship</option><option value="260">Master-Servant Relationship</option><option value="2501">Masturbation</option><option value="6437">Matriarchy</option><option value="279">Mature Protagonist</option><option value="9021">Medical Knowledge</option><option value="4498">Medieval</option><option value="9573">Mercenaries</option><option value="14516">Merchants</option><option value="589">Military</option><option value="6073">Mind Break</option><option value="456">Mind Control</option><option value="11503">Misandry</option><option value="1516">Mismatched Couple</option><option value="172">Misunderstandings</option><option value="105">MMORPG</option><option value="12487">Mob Protagonist</option><option value="14781">Models</option><option value="2606">Modern Day</option><option value="2666">Modern Knowledge</option><option value="3754">Money Grubber</option><option value="510">Monster Girls</option><option value="8988">Monster Society</option><option value="253">Monster Tamer</option><option value="261">Monsters</option><option value="345">Movies</option><option value="20139">Mpreg</option><option value="7163">Multiple Identities</option><option value="8266">Multiple Personalities</option><option value="14639">Multiple POV</option><option value="441">Multiple Protagonists</option><option value="3706">Multiple Realms</option><option value="5149">Multiple Reincarnated Individuals</option><option value="7802">Multiple Timelines</option><option value="16138">Multiple Transported Individuals</option><option value="885">Murders</option><option value="1231">Music</option><option value="5036">Mutated Creatures</option><option value="68">Mutations</option><option value="1141">Mute Character</option><option value="8440">Mysterious Family Background</option><option value="2224">Mysterious Illness</option><option value="4783">Mysterious Past</option><option value="3537">Mystery Solving</option><option value="8995">Mythical Beasts</option><option value="1474">Mythology</option><option value="14644">Naive Protagonist</option><option value="14994">Narcissistic Protagonist</option><option value="6204">Nationalism</option><option value="130">Near-Death Experience</option><option value="186">Necromancer</option><option value="1728">Neet</option><option value="4126">Netorare</option><option value="11561">Netorase</option><option value="3862">Netori</option><option value="6164">Nightmares</option><option value="1725">Ninjas</option><option value="265">Nobles</option><option value="17199">Non-humanoid Protagonist</option><option value="2514">Non-linear Storytelling</option><option value="1524">Nudity</option><option value="7551">Nurses</option><option value="1477">Obsessive Love</option><option value="9331">Office Romance</option><option value="16202">Older Love Interests</option><option value="14574">Omegaverse</option><option value="10980">Oneshot</option><option value="5122">Online Romance</option><option value="2060">Onmyouji</option><option value="6211">Orcs</option><option value="2645">Organized Crime</option><option value="11994">Orgy</option><option value="125">Orphans</option><option value="277">Otaku</option><option value="466">Otome Game</option><option value="726">Outcasts</option><option value="1222">Outdoor Intercourse</option><option value="875">Outer Space</option><option value="4598">Overpowered Protagonist</option><option value="145">Overprotective Siblings</option><option value="8832">Pacifist Protagonist</option><option value="6049">Paizuri</option><option value="318">Parallel Worlds</option><option value="1300">Parasites</option><option value="4741">Parent Complex</option><option value="2350">Parody</option><option value="628">Part-Time Job</option><option value="220">Past Plays a Big Role</option><option value="6457">Past Trauma</option><option value="15546">Persistent Love Interests</option><option value="488">Personality Changes</option><option value="14643">Perverted Protagonist</option><option value="2710">Pets</option><option value="5623">Pharmacist</option><option value="1799">Philosophical</option><option value="1817">Phobias</option><option value="8790">Phoenixes</option><option value="1371">Photography</option><option value="3019">Pill Based Cultivation</option><option value="9071">Pill Concocting</option><option value="1215">Pilots</option><option value="3025">Pirates</option><option value="1311">Playboys</option><option value="13369">Playful Protagonist</option><option value="7813">Poetry</option><option value="2674">Poisons</option><option value="83">Police</option><option value="14996">Polite Protagonist</option><option value="298">Politics</option><option value="11890">Polyandry</option><option value="2684">Polygamy</option><option value="12909">Poor Protagonist</option><option value="8801">Poor to Rich</option><option value="13481">Popular Love Interests</option><option value="94">Possession</option><option value="12966">Possessive Characters</option><option value="1301">Post-apocalyptic</option><option value="2551">Power Couple</option><option value="673">Power Struggle</option><option value="564">Pragmatic Protagonist</option><option value="2020">Precognition</option><option value="3330">Pregnancy</option><option value="1763">Pretend Lovers</option><option value="1124">Previous Life Talent</option><option value="3534">Priestesses</option><option value="2341">Priests</option><option value="1426">Prison</option><option value="701">Proactive Protagonist</option><option value="6302">Programmer</option><option value="13215">Prophecies</option><option value="1892">Prostitutes</option><option value="19606">Protagonist Falls in Love First</option><option value="167">Protagonist Strong from the Start</option><option value="18652">Protagonist with Multiple Bodies</option><option value="13480">Psychic Powers</option><option value="846">Psychopaths</option><option value="9950">Puppeteers</option><option value="13370">Quiet Characters</option><option value="931">Quirky Characters</option><option value="2738">R-15</option><option value="4074">R-18</option><option value="2903">Race Change</option><option value="3314">Racism</option><option value="431">Rape</option><option value="11714">Rape Victim Becomes Lover</option><option value="5574">Rebellion</option><option value="447">Reincarnated as a Monster</option><option value="9480">Reincarnated as an Object</option><option value="7297">Reincarnated in a Game World</option><option value="6304">Reincarnated in Another World</option><option value="120">Reincarnation</option><option value="15178">Religions</option><option value="179">Reluctant Protagonist</option><option value="1684">Reporters</option><option value="1835">Restaurant</option><option value="1209">Resurrection</option><option value="13303">Returning from Another World</option><option value="121">Revenge</option><option value="558">Reverse Harem</option><option value="4500">Reverse Rape</option><option value="28883">Reversible Couple</option><option value="11448">Rich to Poor</option><option value="7780">Righteous Protagonist</option><option value="614">Rivalry</option><option value="334">Romantic Subplot</option><option value="106">Roommates</option><option value="335">Royalty</option><option value="12916">Ruthless Protagonist</option><option value="13496">Sadistic Characters</option><option value="7288">Saints</option><option value="1189">Salaryman</option><option value="826">Samurai</option><option value="788">Saving the World</option><option value="24959">Schemes And Conspiracies</option><option value="2288">Schizophrenia</option><option value="843">Scientists</option><option value="426">Sculptors</option><option value="732">Sealed Power</option><option value="2571">Second Chance</option><option value="1475">Secret Crush</option><option value="214">Secret Identity</option><option value="827">Secret Organizations</option><option value="1190">Secret Relationship</option><option value="14997">Secretive Protagonist</option><option value="1623">Secrets</option><option value="7536">Sect Development</option><option value="2595">Seduction</option><option value="201">Seeing Things Other Humans Can't</option><option value="1463">Selfish Protagonist</option><option value="1372">Selfless Protagonist</option><option value="14319">Seme Protagonist</option><option value="629">Senpai-Kouhai Relationship</option><option value="13520">Sentient Objects</option><option value="13498">Sentimental Protagonist</option><option value="1328">Serial Killers</option><option value="254">Servants</option><option value="3863">Seven Deadly Sins</option><option value="13993">Seven Virtues</option><option value="4432">Sex Friends</option><option value="656">Sex Slaves</option><option value="3252">Sexual Abuse</option><option value="6245">Sexual Cultivation Technique</option><option value="13802">Shameless Protagonist</option><option value="10602">Shapeshifters</option><option value="7874">Sharing A Body</option><option value="13996">Sharp-tongued Characters</option><option value="17746">Shield User</option><option value="2280">Shikigami</option><option value="3358">Short Story</option><option value="5580">Shota</option><option value="7147">Shoujo-Ai Subplot</option><option value="7146">Shounen-Ai Subplot</option><option value="565">Showbiz</option><option value="13499">Shy Characters</option><option value="1971">Sibling Rivalry</option><option value="8186">Sibling's Care</option><option value="1411">Siblings</option><option value="65">Siblings Not Related by Blood</option><option value="13199">Sickly Characters</option><option value="8631">Sign Language</option><option value="209">Singers</option><option value="1985">Single Parent</option><option value="668">Sister Complex</option><option value="3820">Skill Assimilation</option><option value="7032">Skill Books</option><option value="6753">Skill Creation</option><option value="3936">Slave Harem</option><option value="5420">Slave Protagonist</option><option value="180">Slaves</option><option value="3014">Sleeping</option><option value="3185">Slow Growth at Start</option><option value="76">Slow Romance</option><option value="831">Smart Couple</option><option value="652">Social Outcasts</option><option value="674">Soldiers</option><option value="7828">Soul Power</option><option value="4627">Souls</option><option value="8777">Spatial Manipulation</option><option value="4842">Spear Wielder</option><option value="323">Special Abilities</option><option value="2069">Spies</option><option value="2586">Spirit Advisor</option><option value="11422">Spirit Users</option><option value="202">Spirits</option><option value="1713">Stalkers</option><option value="1373">Stockholm Syndrome</option><option value="13500">Stoic Characters</option><option value="13633">Store Owner</option><option value="1191">Straight Seme</option><option value="1594">Straight Uke</option><option value="675">Strategic Battles</option><option value="4777">Strategist</option><option value="14788">Strength-based Social Hierarchy</option><option value="12881">Strong Love Interests</option><option value="4971">Strong to Stronger</option><option value="1643">Stubborn Protagonist</option><option value="608">Student Council</option><option value="1224">Student-Teacher Relationship</option><option value="645">Succubus</option><option value="898">Sudden Strength Gain</option><option value="9574">Sudden Wealth</option><option value="1743">Suicides</option><option value="2990">Summoned Hero</option><option value="4127">Summoning Magic</option><option value="347">Survival</option><option value="348">Survival Game</option><option value="4302">Sword And Magic</option><option value="18792">Sword Wielder</option><option value="7357">System Administrator</option><option value="1749">Teachers</option><option value="1847">Teamwork</option><option value="16962">Technological Gap</option><option value="8760">Tentacles</option><option value="2225">Terminal Illness</option><option value="2196">Terrorists</option><option value="1360">Thieves</option><option value="1420">Threesome</option><option value="2970">Thriller</option><option value="886">Time Loop</option><option value="2054">Time Manipulation</option><option value="887">Time Paradox</option><option value="360">Time Skip</option><option value="92">Time Travel</option><option value="5085">Timid Protagonist</option><option value="267">Tomboyish Female Lead</option><option value="355">Torture</option><option value="14264">Toys</option><option value="148">Tragic Past</option><option value="16178">Transformation Ability</option><option value="3046">Transmigration</option><option value="4323">Transplanted Memories</option><option value="7663">Transported into a Game World</option><option value="6559">Transported Modern Structure</option><option value="15008">Transported to Another World</option><option value="1279">Trap</option><option value="5825">Tribal Society</option><option value="4856">Trickster</option><option value="795">Tsundere</option><option value="518">Twins</option><option value="10488">Twisted Personality</option><option value="12155">Ugly Protagonist</option><option value="4851">Ugly to Beautiful</option><option value="1595">Unconditional Love</option><option value="12907">Underestimated Protagonist</option><option value="3718">Unique Cultivation Technique</option><option value="13875">Unique Weapon User</option><option value="13874">Unique Weapons</option><option value="38534">Unlimited Flow</option><option value="2314">Unlucky Protagonist</option><option value="4697">Unreliable Narrator</option><option value="1268">Unrequited Love</option><option value="17588">Valkyries</option><option value="149">Vampires</option><option value="11404">Villainess Noble Girls</option><option value="109">Virtual Reality</option><option value="1634">Vocaloid</option><option value="2887">Voice Actors</option><option value="3256">Voyeurism</option><option value="1216">Waiters</option><option value="5128">War Records</option><option value="101">Wars</option><option value="12813">Weak Protagonist</option><option value="71">Weak to Strong</option><option value="13334">Wealthy Characters</option><option value="17533">Werebeasts</option><option value="1338">Wishes</option><option value="1829">Witches</option><option value="634">Wizards</option><option value="21767">World Hopping</option><option value="364">World Travel</option><option value="13198">World Tree</option><option value="548">Writers</option><option value="291">Yandere</option><option value="926">Youkai</option><option value="11677">Younger Brothers</option><option value="16201">Younger Love Interests</option><option value="11678">Younger Sisters</option><option value="350">Zombies</option></select> 
		</div> 
	</div> </div>
<div class="one-half">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			 <select multiple id="tags_exclude" class="chzn-select" data-placeholder="Exclude..." name="tagsexclude&#091;&#093;"><option value="16185">Abandoned Children</option><option value="4859">Ability Steal</option><option value="1248">Absent Parents</option><option value="11325">Abusive Characters</option><option value="4885">Academy</option><option value="475">Accelerated Growth</option><option value="13403">Acting</option><option value="4976">Adapted from Manga</option><option value="6280">Adapted from Manhua</option><option value="269">Adapted to Anime</option><option value="2999">Adapted to Drama</option><option value="928">Adapted to Drama CD</option><option value="891">Adapted to Game</option><option value="270">Adapted to Manga</option><option value="182">Adapted to Manhua</option><option value="2721">Adapted to Manhwa</option><option value="1133">Adapted to Movie</option><option value="1334">Adapted to Visual Novel</option><option value="858">Adopted Children</option><option value="603">Adopted Protagonist</option><option value="4974">Adultery</option><option value="3108">Adventurers</option><option value="3802">Affair</option><option value="357">Age Progression</option><option value="216">Age Regression</option><option value="13720">Aggressive Characters</option><option value="234">Alchemy</option><option value="621">Aliens</option><option value="598">All-Girls School</option><option value="7909">Alternate World</option><option value="625">Amnesia</option><option value="774">Amusement Park</option><option value="4594">Anal</option><option value="476">Ancient China</option><option value="2662">Ancient Times</option><option value="13263">Androgynous Characters</option><option value="892">Androids</option><option value="1327">Angels</option><option value="255">Animal Characteristics</option><option value="7323">Animal Rearing</option><option value="724">Anti-Magic</option><option value="1732">Anti-social Protagonist</option><option value="4297">Antihero Protagonist</option><option value="203">Antique Shop</option><option value="513">Apartment Life</option><option value="12913">Apathetic Protagonist</option><option value="217">Apocalypse</option><option value="721">Appearance Changes</option><option value="430">Appearance Different from Actual Age</option><option value="697">Archery</option><option value="6442">Aristocracy</option><option value="9591">Arms Dealers</option><option value="306">Army</option><option value="9522">Army Building</option><option value="2676">Arranged Marriage</option><option value="13427">Arrogant Characters</option><option value="9851">Artifact Crafting</option><option value="4699">Artifacts</option><option value="141">Artificial Intelligence</option><option value="271">Artists</option><option value="81">Assassins</option><option value="8965">Astrologers</option><option value="9776">Autism</option><option value="11491">Automatons</option><option value="12917">Average-looking Protagonist</option><option value="1134">Award-winning Work</option><option value="13016">Awkward Protagonist</option><option value="4975">Bands</option><option value="1346">Based on a Movie</option><option value="1632">Based on a Song</option><option value="6572">Based on a TV Show</option><option value="1442">Based on a Video Game</option><option value="1447">Based on a Visual Novel</option><option value="1230">Based on an Anime</option><option value="825">Battle Academy</option><option value="6308">Battle Competition</option><option value="10668">BDSM</option><option value="5363">Beast Companions</option><option value="5406">Beastkin</option><option value="116">Beasts</option><option value="111">Beautiful Female Lead</option><option value="15625">Bestiality</option><option value="256">Betrayal</option><option value="2573">Bickering Couple</option><option value="13371">Biochip</option><option value="6359">Bisexual Protagonist</option><option value="2943">Black Belly</option><option value="916">Blackmail</option><option value="2734">Blacksmith</option><option value="9280">Blind Dates</option><option value="2347">Blind Protagonist</option><option value="9492">Blood Manipulation</option><option value="5824">Bloodlines</option><option value="2642">Body Swap</option><option value="9907">Body Tempering</option><option value="568">Body-double</option><option value="224">Bodyguards</option><option value="1288">Books</option><option value="13805">Bookworm</option><option value="325">Boss-Subordinate Relationship</option><option value="1710">Brainwashing</option><option value="5607">Breast Fetish</option><option value="9410">Broken Engagement</option><option value="604">Brother Complex</option><option value="8185">Brotherhood</option><option value="3825">Buddhism</option><option value="123">Bullying</option><option value="782">Business Management</option><option value="11786">Businessmen</option><option value="1597">Butlers</option><option value="12910">Calm Protagonist</option><option value="4248">Cannibalism</option><option value="2774">Card Games</option><option value="11345">Carefree Protagonist</option><option value="12951">Caring Protagonist</option><option value="12914">Cautious Protagonist</option><option value="2304">Celebrities</option><option value="257">Character Growth</option><option value="14493">Charismatic Protagonist</option><option value="14985">Charming Protagonist</option><option value="1942">Chat Rooms</option><option value="5708">Cheats</option><option value="1832">Chefs</option><option value="1368">Child Abuse</option><option value="3470">Child Protagonist</option><option value="1391">Childcare</option><option value="85">Childhood Friends</option><option value="917">Childhood Love</option><option value="691">Childhood Promise</option><option value="14986">Childish Protagonist</option><option value="169">Chuunibyou</option><option value="4423">Clan Building</option><option value="7777">Classic</option><option value="14587">Clever Protagonist</option><option value="2356">Clingy Lover</option><option value="2274">Clones</option><option value="611">Clubs</option><option value="14989">Clumsy Love Interests</option><option value="7951">Co-Workers</option><option value="710">Cohabitation</option><option value="14606">Cold Love Interests</option><option value="14605">Cold Protagonist</option><option value="8656">Collection of Short Stories</option><option value="7776">College/University</option><option value="1675">Coma</option><option value="124">Comedic Undertone</option><option value="282">Coming of Age</option><option value="4822">Complex Family Relationships</option><option value="7117">Conditional Power</option><option value="14673">Confident Protagonist</option><option value="1404">Confinement</option><option value="3241">Conflicting Loyalties</option><option value="859">Contracts</option><option value="2506">Cooking</option><option value="352">Corruption</option><option value="9926">Cosmic Wars</option><option value="7844">Cosplay</option><option value="592">Couple Growth</option><option value="5792">Court Official</option><option value="836">Cousins</option><option value="3752">Cowardly Protagonist</option><option value="5944">Crafting</option><option value="2969">Crime</option><option value="443">Criminals</option><option value="769">Cross-dressing</option><option value="1795">Crossover</option><option value="12956">Cruel Characters</option><option value="19556">Cryostasis</option><option value="117">Cultivation</option><option value="4596">Cunnilingus</option><option value="14593">Cunning Protagonist</option><option value="14990">Curious Protagonist</option><option value="579">Curses</option><option value="16457">Cute Children</option><option value="14987">Cute Protagonist</option><option value="5525">Cute Story</option><option value="8248">Dancers</option><option value="10343">Dao Companion</option><option value="6242">Dao Comprehension</option><option value="3824">Daoism</option><option value="4599">Dark</option><option value="2215">Dead Protagonist</option><option value="2142">Death</option><option value="1615">Death of Loved Ones</option><option value="5333">Debts</option><option value="749">Delinquents</option><option value="1362">Delusions</option><option value="5526">Demi-Humans</option><option value="2386">Demon Lord</option><option value="21268">Demonic Cultivation Technique</option><option value="86">Demons</option><option value="14541">Dense Protagonist</option><option value="12964">Depictions of Cruelty</option><option value="1946">Depression</option><option value="3135">Destiny</option><option value="1200">Detectives</option><option value="12959">Determined Protagonist</option><option value="15456">Devoted Love Interests</option><option value="2243">Different Social Status</option><option value="19241">Disabilities</option><option value="258">Discrimination</option><option value="13407">Disfigurement</option><option value="2343">Dishonest Protagonist</option><option value="7949">Distrustful Protagonist</option><option value="2019">Divination</option><option value="4706">Divine Protection</option><option value="2305">Divorce</option><option value="2876">Doctors</option><option value="16412">Dolls/Puppets</option><option value="2663">Domestic Affairs</option><option value="15026">Doting Love Interests</option><option value="15025">Doting Older Siblings</option><option value="14674">Doting Parents</option><option value="509">Dragon Riders</option><option value="897">Dragon Slayers</option><option value="72">Dragons</option><option value="1406">Dreams</option><option value="311">Drugs</option><option value="8126">Druids</option><option value="174">Dungeon Master</option><option value="175">Dungeons</option><option value="3569">Dwarfs</option><option value="1283">Dystopia</option><option value="15108">e-Sports</option><option value="7711">Early Romance</option><option value="1661">Earth Invasion</option><option value="5431">Easy Going Life</option><option value="784">Economics</option><option value="1878">Editors</option><option value="3796">Eidetic Memory</option><option value="13177">Elderly Protagonist</option><option value="9855">Elemental Magic</option><option value="165">Elves</option><option value="13305">Emotionally Weak Protagonist</option><option value="3570">Empires</option><option value="840">Enemies Become Allies</option><option value="113">Enemies Become Lovers</option><option value="263">Engagement</option><option value="2735">Engineer</option><option value="4480">Enlightenment</option><option value="1798">Episodic</option><option value="10593">Eunuch</option><option value="7778">European Ambience</option><option value="10172">Evil Gods</option><option value="5411">Evil Organizations</option><option value="12874">Evil Protagonist</option><option value="7087">Evil Religions</option><option value="2049">Evolution</option><option value="6365">Exhibitionism</option><option value="711">Exorcism</option><option value="200">Eye Powers</option><option value="525">Fairies</option><option value="134">Fallen Angels</option><option value="734">Fallen Nobility</option><option value="12009">Familial Love</option><option value="2152">Familiars</option><option value="641">Family</option><option value="2013">Family Business</option><option value="8664">Family Conflict</option><option value="1833">Famous Parents</option><option value="14555">Famous Protagonist</option><option value="5399">Fanaticism</option><option value="5691">Fanfiction</option><option value="771">Fantasy Creatures</option><option value="99">Fantasy World</option><option value="3068">Farming</option><option value="5068">Fast Cultivation</option><option value="7599">Fast Learner</option><option value="13721">Fat Protagonist</option><option value="15974">Fat to Fit</option><option value="2574">Fated Lovers</option><option value="14675">Fearless Protagonist</option><option value="3953">Fellatio</option><option value="5262">Female Master</option><option value="2879">Female Protagonist</option><option value="4178">Female to Male</option><option value="10919">Feng Shui</option><option value="9340">Firearms</option><option value="1569">First Love</option><option value="4880">First-time Intercourse</option><option value="6061">Flashbacks</option><option value="17020">Fleet Battles</option><option value="923">Folklore</option><option value="1353">Forced into a Relationship</option><option value="606">Forced Living Arrangements</option><option value="3052">Forced Marriage</option><option value="2367">Forgetful Protagonist</option><option value="7405">Former Hero</option><option value="6222">Fox Spirits</option><option value="1576">Friends Become Enemies</option><option value="1476">Friendship</option><option value="797">Fujoshi</option><option value="3468">Futanari</option><option value="2616">Futuristic Setting</option><option value="7643">Galge</option><option value="9425">Gambling</option><option value="93">Game Elements</option><option value="16272">Game Ranking System</option><option value="225">Gamers</option><option value="313">Gangs</option><option value="5382">Gate to Another World</option><option value="9295">Genderless Protagonist</option><option value="13888">Generals</option><option value="17629">Genetic Modifications</option><option value="2014">Genies</option><option value="14581">Genius Protagonist</option><option value="515">Ghosts</option><option value="2044">Gladiators</option><option value="15547">Glasses-wearing Love Interests</option><option value="15548">Glasses-wearing Protagonist</option><option value="11494">Goblins</option><option value="6683">God Protagonist</option><option value="1354">God-human Relationship</option><option value="1355">Goddesses</option><option value="73">Godly Powers</option><option value="177">Gods</option><option value="4437">Golems</option><option value="455">Gore</option><option value="1736">Grave Keepers</option><option value="807">Grinding</option><option value="2005">Guardian Relationship</option><option value="438">Guilds</option><option value="792">Gunfighters</option><option value="1753">Hackers</option><option value="9445">Half-human Protagonist</option><option value="9873">Handjob</option><option value="192">Handsome Male Lead</option><option value="14431">Hard-Working Protagonist</option><option value="14992">Harem-seeking Protagonist</option><option value="8149">Harsh Training</option><option value="2307">Hated Protagonist</option><option value="12141">Healers</option><option value="5267">Heartwarming</option><option value="242">Heaven</option><option value="4770">Heavenly Tribulation</option><option value="4720">Hell</option><option value="15545">Helpful Protagonist</option><option value="808">Herbalist</option><option value="100">Heroes</option><option value="2111">Heterochromia</option><option value="7227">Hidden Abilities</option><option value="6384">Hiding True Abilities</option><option value="10134">Hiding True Identity</option><option value="765">Hikikomori</option><option value="2741">Homunculus</option><option value="1772">Honest Protagonist</option><option value="1387">Hospital</option><option value="13502">Hot-blooded Protagonist</option><option value="12338">Human Experimentation</option><option value="1687">Human Weapon</option><option value="259">Human-Nonhuman Relationship</option><option value="17218">Humanoid Protagonist</option><option value="494">Hunters</option><option value="1711">Hypnotism</option><option value="10075">Identity Crisis</option><option value="3555">Imaginary Friend</option><option value="233">Immortals</option><option value="9033">Imperial Harem</option><option value="626">Incest</option><option value="3398">Incubus</option><option value="2319">Indecisive Protagonist</option><option value="12135">Industrialization</option><option value="506">Inferiority Complex</option><option value="1400">Inheritance</option><option value="5132">Inscriptions</option><option value="550">Insects</option><option value="4539">Interconnected Storylines</option><option value="11229">Interdimensional Travel</option><option value="14993">Introverted Protagonist</option><option value="9064">Investigations</option><option value="3978">Invisibility</option><option value="3687">Jack of All Trades</option><option value="1864">Jealousy</option><option value="6665">Jiangshi</option><option value="7086">Jobless Class</option><option value="5300">JSDF</option><option value="1250">Kidnappings</option><option value="13901">Kind Love Interests</option><option value="3869">Kingdom Building</option><option value="1904">Kingdoms</option><option value="137">Knights</option><option value="571">Kuudere</option><option value="5086">Lack of Common Sense</option><option value="5426">Language Barrier</option><option value="5920">Late Romance</option><option value="4551">Lawyers</option><option value="338">Lazy Protagonist</option><option value="6706">Leadership</option><option value="3430">Legends</option><option value="7443">Level System</option><option value="1402">Library</option><option value="2222">Limited Lifespan</option><option value="3154">Living Abroad</option><option value="1437">Living Alone</option><option value="3715">Loli</option><option value="865">Loneliness</option><option value="651">Loner Protagonist</option><option value="2891">Long Separations</option><option value="1859">Long-distance Relationship</option><option value="1696">Lost Civilizations</option><option value="7151">Lottery</option><option value="1561">Love at First Sight</option><option value="19605">Love Interest Falls in Love First</option><option value="627">Love Rivals</option><option value="607">Love Triangles</option><option value="1616">Lovers Reunited</option><option value="9161">Low-key Protagonist</option><option value="12035">Loyal Subordinates</option><option value="6526">Lucky Protagonist</option><option value="60">Magic</option><option value="6095">Magic Beasts</option><option value="3038">Magic Formations</option><option value="516">Magical Girls</option><option value="3460">Magical Space</option><option value="14900">Magical Technology</option><option value="336">Maids</option><option value="3257">Male Protagonist</option><option value="171">Male to Female</option><option value="557">Male Yandere</option><option value="2604">Management</option><option value="2781">Mangaka</option><option value="12963">Manipulative Characters</option><option value="2084">Manly Gay Couple</option><option value="642">Marriage</option><option value="706">Marriage of Convenience</option><option value="10145">Martial Spirits</option><option value="19604">Masochistic Characters</option><option value="667">Master-Disciple Relationship</option><option value="260">Master-Servant Relationship</option><option value="2501">Masturbation</option><option value="6437">Matriarchy</option><option value="279">Mature Protagonist</option><option value="9021">Medical Knowledge</option><option value="4498">Medieval</option><option value="9573">Mercenaries</option><option value="14516">Merchants</option><option value="589">Military</option><option value="6073">Mind Break</option><option value="456">Mind Control</option><option value="11503">Misandry</option><option value="1516">Mismatched Couple</option><option value="172">Misunderstandings</option><option value="105">MMORPG</option><option value="12487">Mob Protagonist</option><option value="14781">Models</option><option value="2606">Modern Day</option><option value="2666">Modern Knowledge</option><option value="3754">Money Grubber</option><option value="510">Monster Girls</option><option value="8988">Monster Society</option><option value="253">Monster Tamer</option><option value="261">Monsters</option><option value="345">Movies</option><option value="20139">Mpreg</option><option value="7163">Multiple Identities</option><option value="8266">Multiple Personalities</option><option value="14639">Multiple POV</option><option value="441">Multiple Protagonists</option><option value="3706">Multiple Realms</option><option value="5149">Multiple Reincarnated Individuals</option><option value="7802">Multiple Timelines</option><option value="16138">Multiple Transported Individuals</option><option value="885">Murders</option><option value="1231">Music</option><option value="5036">Mutated Creatures</option><option value="68">Mutations</option><option value="1141">Mute Character</option><option value="8440">Mysterious Family Background</option><option value="2224">Mysterious Illness</option><option value="4783">Mysterious Past</option><option value="3537">Mystery Solving</option><option value="8995">Mythical Beasts</option><option value="1474">Mythology</option><option value="14644">Naive Protagonist</option><option value="14994">Narcissistic Protagonist</option><option value="6204">Nationalism</option><option value="130">Near-Death Experience</option><option value="186">Necromancer</option><option value="1728">Neet</option><option value="4126">Netorare</option><option value="11561">Netorase</option><option value="3862">Netori</option><option value="6164">Nightmares</option><option value="1725">Ninjas</option><option value="265">Nobles</option><option value="17199">Non-humanoid Protagonist</option><option value="2514">Non-linear Storytelling</option><option value="1524">Nudity</option><option value="7551">Nurses</option><option value="1477">Obsessive Love</option><option value="9331">Office Romance</option><option value="16202">Older Love Interests</option><option value="14574">Omegaverse</option><option value="10980">Oneshot</option><option value="5122">Online Romance</option><option value="2060">Onmyouji</option><option value="6211">Orcs</option><option value="2645">Organized Crime</option><option value="11994">Orgy</option><option value="125">Orphans</option><option value="277">Otaku</option><option value="466">Otome Game</option><option value="726">Outcasts</option><option value="1222">Outdoor Intercourse</option><option value="875">Outer Space</option><option value="4598">Overpowered Protagonist</option><option value="145">Overprotective Siblings</option><option value="8832">Pacifist Protagonist</option><option value="6049">Paizuri</option><option value="318">Parallel Worlds</option><option value="1300">Parasites</option><option value="4741">Parent Complex</option><option value="2350">Parody</option><option value="628">Part-Time Job</option><option value="220">Past Plays a Big Role</option><option value="6457">Past Trauma</option><option value="15546">Persistent Love Interests</option><option value="488">Personality Changes</option><option value="14643">Perverted Protagonist</option><option value="2710">Pets</option><option value="5623">Pharmacist</option><option value="1799">Philosophical</option><option value="1817">Phobias</option><option value="8790">Phoenixes</option><option value="1371">Photography</option><option value="3019">Pill Based Cultivation</option><option value="9071">Pill Concocting</option><option value="1215">Pilots</option><option value="3025">Pirates</option><option value="1311">Playboys</option><option value="13369">Playful Protagonist</option><option value="7813">Poetry</option><option value="2674">Poisons</option><option value="83">Police</option><option value="14996">Polite Protagonist</option><option value="298">Politics</option><option value="11890">Polyandry</option><option value="2684">Polygamy</option><option value="12909">Poor Protagonist</option><option value="8801">Poor to Rich</option><option value="13481">Popular Love Interests</option><option value="94">Possession</option><option value="12966">Possessive Characters</option><option value="1301">Post-apocalyptic</option><option value="2551">Power Couple</option><option value="673">Power Struggle</option><option value="564">Pragmatic Protagonist</option><option value="2020">Precognition</option><option value="3330">Pregnancy</option><option value="1763">Pretend Lovers</option><option value="1124">Previous Life Talent</option><option value="3534">Priestesses</option><option value="2341">Priests</option><option value="1426">Prison</option><option value="701">Proactive Protagonist</option><option value="6302">Programmer</option><option value="13215">Prophecies</option><option value="1892">Prostitutes</option><option value="19606">Protagonist Falls in Love First</option><option value="167">Protagonist Strong from the Start</option><option value="18652">Protagonist with Multiple Bodies</option><option value="13480">Psychic Powers</option><option value="846">Psychopaths</option><option value="9950">Puppeteers</option><option value="13370">Quiet Characters</option><option value="931">Quirky Characters</option><option value="2738">R-15</option><option value="4074">R-18</option><option value="2903">Race Change</option><option value="3314">Racism</option><option value="431">Rape</option><option value="11714">Rape Victim Becomes Lover</option><option value="5574">Rebellion</option><option value="447">Reincarnated as a Monster</option><option value="9480">Reincarnated as an Object</option><option value="7297">Reincarnated in a Game World</option><option value="6304">Reincarnated in Another World</option><option value="120">Reincarnation</option><option value="15178">Religions</option><option value="179">Reluctant Protagonist</option><option value="1684">Reporters</option><option value="1835">Restaurant</option><option value="1209">Resurrection</option><option value="13303">Returning from Another World</option><option value="121">Revenge</option><option value="558">Reverse Harem</option><option value="4500">Reverse Rape</option><option value="28883">Reversible Couple</option><option value="11448">Rich to Poor</option><option value="7780">Righteous Protagonist</option><option value="614">Rivalry</option><option value="334">Romantic Subplot</option><option value="106">Roommates</option><option value="335">Royalty</option><option value="12916">Ruthless Protagonist</option><option value="13496">Sadistic Characters</option><option value="7288">Saints</option><option value="1189">Salaryman</option><option value="826">Samurai</option><option value="788">Saving the World</option><option value="24959">Schemes And Conspiracies</option><option value="2288">Schizophrenia</option><option value="843">Scientists</option><option value="426">Sculptors</option><option value="732">Sealed Power</option><option value="2571">Second Chance</option><option value="1475">Secret Crush</option><option value="214">Secret Identity</option><option value="827">Secret Organizations</option><option value="1190">Secret Relationship</option><option value="14997">Secretive Protagonist</option><option value="1623">Secrets</option><option value="7536">Sect Development</option><option value="2595">Seduction</option><option value="201">Seeing Things Other Humans Can't</option><option value="1463">Selfish Protagonist</option><option value="1372">Selfless Protagonist</option><option value="14319">Seme Protagonist</option><option value="629">Senpai-Kouhai Relationship</option><option value="13520">Sentient Objects</option><option value="13498">Sentimental Protagonist</option><option value="1328">Serial Killers</option><option value="254">Servants</option><option value="3863">Seven Deadly Sins</option><option value="13993">Seven Virtues</option><option value="4432">Sex Friends</option><option value="656">Sex Slaves</option><option value="3252">Sexual Abuse</option><option value="6245">Sexual Cultivation Technique</option><option value="13802">Shameless Protagonist</option><option value="10602">Shapeshifters</option><option value="7874">Sharing A Body</option><option value="13996">Sharp-tongued Characters</option><option value="17746">Shield User</option><option value="2280">Shikigami</option><option value="3358">Short Story</option><option value="5580">Shota</option><option value="7147">Shoujo-Ai Subplot</option><option value="7146">Shounen-Ai Subplot</option><option value="565">Showbiz</option><option value="13499">Shy Characters</option><option value="1971">Sibling Rivalry</option><option value="8186">Sibling's Care</option><option value="1411">Siblings</option><option value="65">Siblings Not Related by Blood</option><option value="13199">Sickly Characters</option><option value="8631">Sign Language</option><option value="209">Singers</option><option value="1985">Single Parent</option><option value="668">Sister Complex</option><option value="3820">Skill Assimilation</option><option value="7032">Skill Books</option><option value="6753">Skill Creation</option><option value="3936">Slave Harem</option><option value="5420">Slave Protagonist</option><option value="180">Slaves</option><option value="3014">Sleeping</option><option value="3185">Slow Growth at Start</option><option value="76">Slow Romance</option><option value="831">Smart Couple</option><option value="652">Social Outcasts</option><option value="674">Soldiers</option><option value="7828">Soul Power</option><option value="4627">Souls</option><option value="8777">Spatial Manipulation</option><option value="4842">Spear Wielder</option><option value="323">Special Abilities</option><option value="2069">Spies</option><option value="2586">Spirit Advisor</option><option value="11422">Spirit Users</option><option value="202">Spirits</option><option value="1713">Stalkers</option><option value="1373">Stockholm Syndrome</option><option value="13500">Stoic Characters</option><option value="13633">Store Owner</option><option value="1191">Straight Seme</option><option value="1594">Straight Uke</option><option value="675">Strategic Battles</option><option value="4777">Strategist</option><option value="14788">Strength-based Social Hierarchy</option><option value="12881">Strong Love Interests</option><option value="4971">Strong to Stronger</option><option value="1643">Stubborn Protagonist</option><option value="608">Student Council</option><option value="1224">Student-Teacher Relationship</option><option value="645">Succubus</option><option value="898">Sudden Strength Gain</option><option value="9574">Sudden Wealth</option><option value="1743">Suicides</option><option value="2990">Summoned Hero</option><option value="4127">Summoning Magic</option><option value="347">Survival</option><option value="348">Survival Game</option><option value="4302">Sword And Magic</option><option value="18792">Sword Wielder</option><option value="7357">System Administrator</option><option value="1749">Teachers</option><option value="1847">Teamwork</option><option value="16962">Technological Gap</option><option value="8760">Tentacles</option><option value="2225">Terminal Illness</option><option value="2196">Terrorists</option><option value="1360">Thieves</option><option value="1420">Threesome</option><option value="2970">Thriller</option><option value="886">Time Loop</option><option value="2054">Time Manipulation</option><option value="887">Time Paradox</option><option value="360">Time Skip</option><option value="92">Time Travel</option><option value="5085">Timid Protagonist</option><option value="267">Tomboyish Female Lead</option><option value="355">Torture</option><option value="14264">Toys</option><option value="148">Tragic Past</option><option value="16178">Transformation Ability</option><option value="3046">Transmigration</option><option value="4323">Transplanted Memories</option><option value="7663">Transported into a Game World</option><option value="6559">Transported Modern Structure</option><option value="15008">Transported to Another World</option><option value="1279">Trap</option><option value="5825">Tribal Society</option><option value="4856">Trickster</option><option value="795">Tsundere</option><option value="518">Twins</option><option value="10488">Twisted Personality</option><option value="12155">Ugly Protagonist</option><option value="4851">Ugly to Beautiful</option><option value="1595">Unconditional Love</option><option value="12907">Underestimated Protagonist</option><option value="3718">Unique Cultivation Technique</option><option value="13875">Unique Weapon User</option><option value="13874">Unique Weapons</option><option value="38534">Unlimited Flow</option><option value="2314">Unlucky Protagonist</option><option value="4697">Unreliable Narrator</option><option value="1268">Unrequited Love</option><option value="17588">Valkyries</option><option value="149">Vampires</option><option value="11404">Villainess Noble Girls</option><option value="109">Virtual Reality</option><option value="1634">Vocaloid</option><option value="2887">Voice Actors</option><option value="3256">Voyeurism</option><option value="1216">Waiters</option><option value="5128">War Records</option><option value="101">Wars</option><option value="12813">Weak Protagonist</option><option value="71">Weak to Strong</option><option value="13334">Wealthy Characters</option><option value="17533">Werebeasts</option><option value="1338">Wishes</option><option value="1829">Witches</option><option value="634">Wizards</option><option value="21767">World Hopping</option><option value="364">World Travel</option><option value="13198">World Tree</option><option value="548">Writers</option><option value="291">Yandere</option><option value="926">Youkai</option><option value="11677">Younger Brothers</option><option value="16201">Younger Love Interests</option><option value="11678">Younger Sisters</option><option value="350">Zombies</option></select> 
		</div> 
	</div> </div>
</div>      
<div class="extopsf"></div>   
<span class="sftitle sfpad new">Reading Lists <select name="releases_mm" id="rli_mm" class="select sf" style="width:70px;">
						<option value="exclude" selected>Exclude</option>
						<option value="include">Include</option>
					</select>]<i onclick="sf_getnovelinfo('readinglists', 'Reading Lists');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 

<div class="extrasf"></div>     

You must be logged in to use this option.  
<div class="extopsf"></div>      

 
<span class="sftitle sfpad new">Story Status</span> 
<br />

<div>
 <select name="storystatus" id="storystatus" class="storystatus" onchange="">
  <option value="1" selected="">All</option>
  <option value="2">Completed</option>
  <option value="3">Ongoing</option>
  <option value="4">Hiatus</option>
  </select>
</div>
 <div class="extopsf"></div>
    
<span class="sftitle sfpad new">Sort Results By... <i onclick="sf_getnovelinfo('sorting', 'Sort Results By...');" class="fa fa-info-circle sfinder" title="Help" aria-hidden="true"></i></span> 

<div class="extrasf"></div>
<div class="g-cols wpb_row offset_default">
<div class="one-half">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			
 <select name="sortmyresults" id="sortresults" class="sortresults" style="width: 40%;" onchange="">
  <option value="srel">Chapters</option>
  <option value="sfrel">Frequency</option>
  <option value="srank">Rank</option>
  <option value="srate">Rating</option>
  <option value="sread">Readers</option>
  <option value="sreview">Reviews</option>
  <option value="abc">Title</option>
  <option value="sdate" selected="">Last Updated</option>
  </select> 
 
		</div> 
	</div> </div>
<div class="one-half">
	<div class="wpb_text_column ">
		<div class="wpb_wrapper">
			 
<select name="sortmyorder" id="sortorder" class="sortorder" style="width: 40%;" onchange="">
  <option value="asc">Ascending</option>
  <option value="desc" selected="">Descending</option>
  </select> 
 
		</div> 
	</div> </div>
</div>  
<div class="extopsf"></div>
<a class="crfr finder" href="//www.novelupdates.com/series-finder/">Clear Filters</a>
</div>  <input type="submit" id="rk_filter" class="btnrev review" style="float:none; margin-top: 15px;" onclick="applyfilter('//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc');" value="Apply Filters"> 
</div>
</form>


   	                 
                
            <div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_46672.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (3.4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid46672" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/another-world-awakening-transcendental-create-skill/"></a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i></span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i></span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i></span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i></span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a></div>A bus accident during a school trip sends the entire class to another world. After being transported to another world, they are attacked by monsters and bandits, and while their classmates awaken sword and magic skills one after another, Takaya Nagami’s abilities never awaken, and he ends up being kicked out of the class party. Finding himself all alone in a world where he can’t rely on anyone, Takaya despairs and impulsively tries to commit suicide, but he survived and is rescued by local adventurers while in a miserable state.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
In the course of his adventures with these adventurers, he gradually gains strength, confidence, and trust (as well as cute girls!) by using his production and processing skills. An exhilarating and thrilling tale of growth in a different world that follows the dullest boy in his class as he becomes an unstoppable adventurer!
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_30279.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (4.1)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid30279" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/ill-become-a-villainess-that-will-go-down-in-history/"></a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i></span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i></span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i></span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i></span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="9" href="https://www.novelupdates.com/genre/drama/" title="View All Drama Related Series" >Drama</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="157" href="https://www.novelupdates.com/genre/shoujo/" title="View All Shoujo Related Series" >Shoujo</a></div>I’ve always dreamed of becoming a villainess, but I never thought that I would actually become one….!<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
This is the story of a young girl who aspires to become a villainous noble girl who’s capable of growing stronger through every confrontation that she is faced with.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_26659.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.1)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid26659" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/i-help-the-richest-man-spend-money-to-prevent-disasters/"></a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i></span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i></span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i></span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i></span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a></div>Ye Zhi not only inherited the run-down house, but also the engagement with a man. The man who met his true love disliked his fiancée in all ways and asked for divorce.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
In order to express their guilt, the man’s family decided to make up for it with Ye Zhi. But she did not want their ‘compensation’ and simply left without turning her back.<p style="margin-top:-5px;"></p>
The next second, Ye Zhi was picked up by a Rolls Royce.<p style="margin-top:-5px;"></p>
Gu Ren, the son of the richest man, had a tough life and couldn’t live past 30 years old. Only Ye Zhi, who was born on the cloudy day of the lunar year, was the most suitable match that could make him survive.<p style="margin-top:-5px;"></p>
Ye Zhi and Gu Ren got married. Gu Ren gave her a bank card. All she has to do is—<p style="margin-top:-5px;"></p>
Spend money for him! Squander his fortune!<p style="margin-top:-5px;"></p>
Ye Zhi bought a Lamborghini for Gu Ren, but he had an accident as soon as he drove the car.<p style="margin-top:-5px;"></p>
As soon as she bought some Hermes platinum bags for herself, Gu Ren signed a more than one billion worth contract.<p style="margin-top:-5px;"></p>
It turned out that the money could only be spent on herself.  The more she spent, the more prosperous she became.<p style="margin-top:-5px;"></p>
Since then her life has changed completely. She became a rich woman who owns luxury items, limited edition bags and a villa overlooking the Forbidden City.<p style="margin-top:-5px;"></p>
The man thought that Ye Zhi, who he had broken up with, must have suffered a lot. Until one day, he saw his former fiancée walking in the street, followed by a striking Lamborghini.<p style="margin-top:-5px;"></p>
Wait! Who is the man next to her?<p style="margin-top:-5px;"></p>
This man is a super-rich man—Hundreds of times richer than his own family!
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_45277.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (3.8)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid45277" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/i-just-want-to-be-a-salted-fish-quietly/">I Just Want To Be a Salted Fish Quietly</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 71 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 10.2 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 649 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 5 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="9" href="https://www.novelupdates.com/genre/drama/" title="View All Drama Related Series" >Drama</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="245" href="https://www.novelupdates.com/genre/mystery/" title="View All Mystery Related Series" >Mystery</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="6" href="https://www.novelupdates.com/genre/school-life/" title="View All School Life Related Series" >School Life</a> <a class="gennew search" gid="7" href="https://www.novelupdates.com/genre/slice-of-life/" title="View All Slice of Life Related Series" >Slice of Life</a> <a class="gennew search" gid="16" href="https://www.novelupdates.com/genre/supernatural/" title="View All Supernatural Related Series" >Supernatural</a></div>Lin Xian said bitterly: I really just want to be a salted fish quietly.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
System: No, you can’t!<p style="margin-top:-5px;"></p>
System: Congratulations to the host for obtaining the VIP privilege gift package and gaining the skill, 【Driving a Ferrari with one hand】…<p style="margin-top:-5px;"></p>
Congratulations to the host for obtaining props, 【Absolute focus】…<p style="margin-top:-5px;"></p>
Congratulations to the host for obtaining props, 【Consumption Rebate Critical Card】…<p style="margin-top:-5px;"></p>
…<p style="margin-top:-5px;"></p>
With the help of the system, Lin Xian found that he seemed to be on the path to becoming an Almighty Male God, getting farther and farther…<p style="margin-top:-5px;"></p>
(Salted Fish – a term used to refer to a person who has no intention of doing anything or person who wants to live a carefree life)<p style="margin-top:-5px;"></p>
(Male God – Prince Charming / Mister Right)
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_39787.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.1)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid39787" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/tai-sui/">Tai Sui</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 251 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 1.1 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 907 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 5 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="330" href="https://www.novelupdates.com/genre/historical/" title="View All Historical Related Series" >Historical</a> <a class="gennew search" gid="480" href="https://www.novelupdates.com/genre/xianxia/" title="View All Xianxia Related Series" >Xianxia</a></div>If I had a choice, I would only want to be a little insect in the mundane dust, born in confusion, dying in mediocrity, never seeing the light of day beneath the fog of Jinping City.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
Better than taking this wrong road to heaven.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_27879.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (3.7)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid27879" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/hellmode-a-hardcore-gamer-becomes-peerless-in-another-world-with-retro-game-settings/">Hellmode ~A Hardcore Gamer Becomes Peerless in Another World with Retro Game Settings~</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 264 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 1.3 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 4730 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 23 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="6" href="https://www.novelupdates.com/genre/school-life/" title="View All School Life Related Series" >School Life</a> <a class="gennew search" gid="16" href="https://www.novelupdates.com/genre/supernatural/" title="View All Supernatural Related Series" >Supernatural</a></div>Kenichi Yamada was a 35-year-old salaryman. As a hardcore game enthusiast, he was saddened by the modern trend towards casual games.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
So, when a site claimed to be “for people who like to do things the hard way,” he just couldn’t resist. Thus, he was reincarnated into another world as Allen, playing on Hell Mode.<p style="margin-top:-5px;"></p>
This is fantasy light novel about Allen’s journey as a summoner.<p style="margin-top:-5px;"></p>
Reincarnated as a serf, he starts from nothing.<p style="margin-top:-5px;"></p>
His journey begins with absolutely no knowledge of the world around him, just like those games he played 10, 20 years ago.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_33838.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.2)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid33838" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/little-tyrant-doesnt-want-to-meet-with-a-bad-end/">Little Tyrant Doesn&#8217;t Want to Meet with a Bad End</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 562 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 1.7 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 10072 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 244 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a></div>The moment the little tyrant of the nobles, Roel Ascart, saw his stepsister, he recalled his memories. He realized that he was in the world of a gal game he played in his previous life. To make things worse, he was the greatest villain in the common route of the game!<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
“I’ll be killed by the main character and the four capture targets ten years from now. Is there still any hope for me?”<p style="margin-top:-5px;"></p>
Just thinking about the fearsome glints of those sharp swords those beautiful capture targets held in their hands, Roel couldn’t help but tremble in fear.<p style="margin-top:-5px;"></p>
Till a voice finally sounded in his head.<p style="margin-top:-5px;"></p>
【Welcome to the House Resurgence System】
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_22254.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid22254" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/star-odyssey/">Star Odyssey</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 1264 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.5 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 2122 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 17 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="9" href="https://www.novelupdates.com/genre/drama/" title="View All Drama Related Series" >Drama</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="11" href="https://www.novelupdates.com/genre/sci-fi/" title="View All Sci-fi Related Series" >Sci-fi</a> <a class="gennew search" gid="12" href="https://www.novelupdates.com/genre/shounen/" title="View All Shounen Related Series" >Shounen</a> <a class="gennew search" gid="3954" href="https://www.novelupdates.com/genre/xuanhuan/" title="View All Xuanhuan Related Series" >Xuanhuan</a></div>Join Lu Yin on an epic journey across the Universe, pursuing the truth and tragedy of his past. This is a world of science fantasy<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"> where the older generations step back and allow the young to take charge of affairs. Heart-wrenching separations, terrifying situations, all with comic relief that will leave you coming back for more. This is a world where the other characters actually matter, and are revisited frequently as their own lives unfold. Dotting Lu Yin’s path are monumental feats of kingdom-building and treacherous political situations where he must tread carefully if he wants to get to the truth of his history.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_52799.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgkr" style="font-weight: bold;">KR</span> (3.9)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid52799" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/how-to-live-as-the-vampire-lord/">How to Live As the Vampire Lord</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 67 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.9 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 619 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 3 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a></div>Vampire Eugene -- a sacrificial lamb slaughtered after half a year of running to fulfil the desires of a templar for fame.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
He was given a second chance at life after concluding his life with regrets.<p style="margin-top:-5px;"></p>
"I will never again die the same way. If I really did return to the past, no matter what it takes... I will never regret again."
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_52696.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.3)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid52696" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/the-reincarnated-villain-makes-the-heroines-tearfully-beg-for-forgiveness/">The Reincarnated Villain Makes The Heroines Tearfully Beg for Forgiveness</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 10 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 3.9 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 441 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 6 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="280" href="https://www.novelupdates.com/genre/adult/" title="View All Adult Related Series" >Adult</a> <a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="4" href="https://www.novelupdates.com/genre/mature/" title="View All Mature Related Series" >Mature</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="3954" href="https://www.novelupdates.com/genre/xuanhuan/" title="View All Xuanhuan Related Series" >Xuanhuan</a></div>Lin Yan had transmigrated into the world of Immortals and had become a Heavenly Emperor who was already defeated by the Male Lead along with the ten Female Leads at the beginning of the plot.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
But as Lin Yan was about to die, he awakened a power that he could use to rewrite his past doings.<p style="margin-top:-5px;"></p>
So, Lin Yan has begun to rewrite his past.<p style="margin-top:-5px;"></p>
When Lin Yan was still a child, he struggled to feed Xiao Yanran, one of the female leads, who was then younger than him. He had to strip his Supreme Bone and give it to her, protect her from the obstacles. He made a vow, “As long as I am here, no one will be allowed to bully you.”<p style="margin-top:-5px;"></p>
Later when Lin Yan’s cultivation rose , he had run into Xiao Mei, yet another heroine, and the two of them spent time together and eventually got married. When the Devils were about to invade their world, Lin Yan said staring right into his wife’s eyes, “I will protect you with my life…”<p style="margin-top:-5px;"></p>
And soon he rose to become the Heavenly Emperor.<p style="margin-top:-5px;"></p>
Everything he did was to shelter his loved ones and to protect the world.<p style="margin-top:-5px;"></p>
However, his actions were misunderstood by the world.<p style="margin-top:-5px;"></p>
…<p style="margin-top:-5px;"></p>
Let’s see what all Lin Yan had done in the past…
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_23882.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (3.3)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid23882" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/even-if-im-reborn-as-a-cute-dragon-girl-i-will-still-make-a-harem/">Even If I&#8217;m Reborn as a Cute Dragon Girl, I Will Still Make a Harem</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 296 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 3.3 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 3234 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 28 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="168" href="https://www.novelupdates.com/genre/gender-bender/" title="View All Gender Bender Related Series" >Gender Bender</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="6" href="https://www.novelupdates.com/genre/school-life/" title="View All School Life Related Series" >School Life</a> <a class="gennew search" gid="851" href="https://www.novelupdates.com/genre/shoujo-ai/" title="View All Shoujo Ai Related Series" >Shoujo Ai</a></div>We, Aristides Castro G Morris Brooklyn Washington Napoleon George I, are the lord of this empire. Our demise art the result of saving our people from impending disaster. In order to award us for our glorious achievement, the goddess hast granted us the chance at rebirth. However… why did we turn into a girl, and the princess of dragons at that?!<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
This is the story of a man who suffered from Eight-Grade Syndrome (chuunibyou) and has been reincarnated as the dragons’ princess. However, because of his chuunibyou behavior, one of the goddesses was annoyed at him and cast a curse on him. One that summoned lightning upon him whenever his symptom acts up.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_49371.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (3)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid49371" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/i-reincarnated-but-will-try-to-live-without-using-my-cheat-ability/">I Reincarnated But Will Try To Live Without Using My Cheat Ability</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 294 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.6 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 939 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 8 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="12" href="https://www.novelupdates.com/genre/shounen/" title="View All Shounen Related Series" >Shounen</a></div>When someone says, “I’ll give you a cheat ability, so just use it”, don’t you feel like not using it?<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
There is nothing more expensive than free stuff.<p style="margin-top:-5px;"></p>
It’s not cool to get excited about something you got for free.<p style="margin-top:-5px;"></p>
In a sense, the protagonist is a reincarnation target, but he perseveres while messing with God’s plans.<p style="margin-top:-5px;"></p>
God makes a mistake in choosing his reincarnation target, and later learns of it.<p style="margin-top:-5px;"></p>
He later realizes that it’s not like everyone wants to be strong and not everyone wants to make a harem.<p style="margin-top:-5px;"></p>
And that he has chosen someone who is more attentive than others and is a real pain in the ass…
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_50621.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (4.1)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid50621" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/the-story-of-a-manga-artist-who-was-imprisoned-by-a-strange-high-school-girl/">The Story of a Manga Artist Who Was Imprisoned by a Strange High School Girl</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 18 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 6.5 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 251 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 0 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="9" href="https://www.novelupdates.com/genre/drama/" title="View All Drama Related Series" >Drama</a> <a class="gennew search" gid="292" href="https://www.novelupdates.com/genre/ecchi/" title="View All Ecchi Related Series" >Ecchi</a> <a class="gennew search" gid="343" href="https://www.novelupdates.com/genre/horror/" title="View All Horror Related Series" >Horror</a> <a class="gennew search" gid="245" href="https://www.novelupdates.com/genre/mystery/" title="View All Mystery Related Series" >Mystery</a> <a class="gennew search" gid="486" href="https://www.novelupdates.com/genre/psychological/" title="View All Psychological Related Series" >Psychological</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="18" href="https://www.novelupdates.com/genre/seinen/" title="View All Seinen Related Series" >Seinen</a> <a class="gennew search" gid="7" href="https://www.novelupdates.com/genre/slice-of-life/" title="View All Slice of Life Related Series" >Slice of Life</a></div>"Where am I?" When I awoke, I noticed a strange ceiling, no lights around me, and a――collar and chain around my neck!? The room's light<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"> suddenly turns on. There was a high school girl holding a kitchen knife. A JK and a manga artist, their life of confinement is about to begin!
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_490.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (3.8)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid490" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/manowa/">Manowa</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 442 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 9.1 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 4682 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 12 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a></div>When she came to, she was in another world. A world just like the stage of a popular game, the girl “Kazane” collects and fights with monster skills in a Hyper Learning Story.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
Before she realizes it, she’s saved a town, exposed a scandal, met a ghost, saved a princess, went into a hot spring, and spent her time going to the hot spring daily with her party. In the fight against the approaching devil, Kazane confronts her own dark side during her destined reunion with the Big Cat. Did I mention she also spends time at a hot spring?
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/noimagemid.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (3.7)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid53841" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/an-s-class-adventurer-banished-as-the-tank-of-the-party-uses-his-slave-release-skill-to-build-the-strongest-country-in-history/">An S-class Adventurer, Banished As The Tank of the Party, uses his &#8220;S*ave Release&#8221; skill to Build the Strongest Country in History!</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 5 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.4 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 106 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 0 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a></div>I became an S-class adventurer with my childhood friends, using the strength that I’m proud of to attract enemies and have my friends defeat them in the meantime, continuing to play the so-called “tank” role.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
I was a paladin, a job specializing in protecting my companions, but perhaps because everyone was protected by me, they did not learn any defense or evasion skills, but only improved their attack skills.<p style="margin-top:-5px;"></p>
In such a situation, they forgot that they are protected by my skills, and they fired me …… Wait, I don’t care what happens, okay!? Even though your natural defensive power are as good as paper.<p style="margin-top:-5px;"></p>
But since I had been through a terrible time on my way up to become an S-class adventurer, I decided to take it easy for once and went to the Adventurers’ Guild to accept a request for a relaxing day alone.<p style="margin-top:-5px;"></p>
I was told that they had tried to cultivate the land in the past, but the monsters that appeared were so strong that all the farmers had run away, so the vast land would be given as a reward to anyone who succeeds. they were rewarded with a large tract of land.<p style="margin-top:-5px;"></p>
When I went there after being told that “Alex, an S-class adventurer, would be fine!” …… I found that there was not even a village or a house around it, and it’s beyond of what you call the countryside, because it was a land inhabited by monsters.<p style="margin-top:-5px;"></p>
The land was deserted, with farming tools left unattended, and there was a monster that I had never seen before, and when I managed to defeat it, I acquired the “s*ave Release” skill.<p style="margin-top:-5px;"></p>
I was told that if I used this skill, it would summon one person of the opposite sex, who is ens*aved somewhere in the world, to me.<p style="margin-top:-5px;"></p>
I thought it would be better for the world if I used this skill, and I activated it immediately, and a beautiful elf girl appeared.<p style="margin-top:-5px;"></p>
After that, a girl with architecture skill, a princess from a foreign country who was a prisoner …… and I kept freeing people who were ens*aved, and before I knew it, a country was formed and I was treated like a king.<p style="margin-top:-5px;"></p>
We received a special award in the Otherworldly Fantasy category at Kakuyomukon 6!<p style="margin-top:-5px;"></p>
≪With the decision to publish a book, the old title “No need for a tank role! I was banished from the S-class party even though I helped my paper defense friends, but I was given the «s*ave Release» skill in the undeveloped land, and while being loved by the beautiful girls who helped me … the strongest country in history was made. The title has been changed.≫<p style="margin-top:-5px;"></p>
*Episode XX: Protagonist’s Point of View<p style="margin-top:-5px;"></p>
Episode X: The point of view of the character written in the title.<p style="margin-top:-5px;"></p>
The following is an example.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/noimagemid.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid53749" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/i-the-clock-user-was-kicked-out-of-the-craftsmens-guild-for-being-incompetent-but-at-the-dungeons-depths-i-awakened-my-true-power-even-though-i-was-told-to-return-because-they-couldnt/">I, the &#8220;Clock User&#8221;, was kicked out of the craftsmen&#8217;s guild for being incompetent, but at the dungeon&#8217;s depths, I awakened my true power &#8211; even though I was told to return because they couldn&#8217;t handle the work, it&#8217;s too late now, and I will live freely as an SSS-class adventurer.</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 14 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.4 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 89 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 0 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a></div>Cyclo Owen. As a genius boy engineer, he was a person of great promise in the circles of magical tool production. However, at the age of 16, he was given the job skill of "Clock User" in a skill-giving ceremony. In this world, it was forbidden for those with job skills to work in any other occupation. Thus, Cyclo's path as a genius engineer was cut off.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
However, Cyclo did not despair and worked diligently as a watchmaker and contributed to the craftsmen's guild by using his job skill as a Clock User. One day, however, the guildmaster of the craftsmen's guild banished him from the guild, saying that the guild did not need a watchmaker, a profession with little usefulness. Cyclo, who has lost his job, faces further despair. He happens to hear that a police officer is planning to r*pe Miranda Lilibel, the neighborhood apothecary's lady. Since the police is involved, he cannot carelessly report the crime. So Cyclo decides to stop the plan that night. He tries to restrain the police officer who is about to attack Miranda's house, but he is detained by the officer instead. He is then accused of trying to r*pe Miranda, and when Miranda hears the commotion, she tells Cyclo she has misjudged him.<p style="margin-top:-5px;"></p>
Cyclo is thus made a criminal for attempted r*pe, and is sent to a remote area to explore the worst dungeon in the world. In the frontier, imprisoned adventurers are considered to be at the bottom of the heap, so Cyclo is looked down upon. Furthermore, due to rumors spread by his younger sister, Alice Owen, a well-known genius alchemist among adventurers, he is treated as incompetent here as well.<p style="margin-top:-5px;"></p>
He is then mistreated by adventurers who makes uses him as a baggage carrier, and is pushed down to the bottom of a valley in the dungeon.<p style="margin-top:-5px;"></p>
At the very bottom of the dungeon, Cyclo wished he could use his Clock User abilities to stop his internal clock. Then his hunger stopped. This led Cyclo to understand that his skill coul manipulate anything that could be counted as a clock.<p style="margin-top:-5px;"></p>
Thus awakening to his true power, Cyclo escaped from the deepest part of the worst dungeon. In recognition of his achievements and abilities, he is recognized as an SSS-class adventurer.<p style="margin-top:-5px;"></p>
On the other hand, those who had driven Cyclo to misfortune are now faced with misfortune in Cyclo's absence.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_12586.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (3.5)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid12586" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/mechanical-god-emperor/">Mechanical God Emperor</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 1233 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.9 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 3899 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 37 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="10" href="https://www.novelupdates.com/genre/mecha/" title="View All Mecha Related Series" >Mecha</a> <a class="gennew search" gid="11" href="https://www.novelupdates.com/genre/sci-fi/" title="View All Sci-fi Related Series" >Sci-fi</a> <a class="gennew search" gid="3954" href="https://www.novelupdates.com/genre/xuanhuan/" title="View All Xuanhuan Related Series" >Xuanhuan</a></div>Yang Feng somehow transmigrated into a different world and received a legacy of an ‘ancient high tech’ family, which does not directly raise his power, but gives him the technology to build things which are way more advanced than the seemingly medieval world.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
But to build something you need resources and energy. To receive resources you need strength. To gain strength you need knowledge. To gain knowledge you… need strength? or a background? Or maybe a fully armed army of high tech robots who aren’t afraid of death?<p style="margin-top:-5px;"></p>
But is this legacy really for him to keep?
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_47256.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (4.6)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid47256" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/strange-memories-with-you-my-dear/">Strange Memories With You, My Dear</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 16 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 13.3 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 261 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 0 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="245" href="https://www.novelupdates.com/genre/mystery/" title="View All Mystery Related Series" >Mystery</a> <a class="gennew search" gid="486" href="https://www.novelupdates.com/genre/psychological/" title="View All Psychological Related Series" >Psychological</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="16" href="https://www.novelupdates.com/genre/supernatural/" title="View All Supernatural Related Series" >Supernatural</a> <a class="gennew search" gid="132" href="https://www.novelupdates.com/genre/tragedy/" title="View All Tragedy Related Series" >Tragedy</a></div>I’ll come to your aid. I’m not going to give up no matter what―<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
In his dreams, Yukinari has been interacting and spending time with a mysterious girl. He experiences a sensation of déjà vu one day when he visits a place for the first time, and there he meets Yuuko, the girl from his dream.And, to her astonishment, she had met Yukinari in a dream as well. Then they start to go out together to confirm the phenomenon, and eventually fall in love. But one day, when everything appeared to be going well, Yuuko fainted and went into a coma….. A story about two people who have the same dream and are looking for a miracle――
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_41062.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (3.8)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid41062" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/youre-my-belated-happiness/">You&#8217;re My Belated Happiness</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 111 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 3.3 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 835 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 6 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="7" href="https://www.novelupdates.com/genre/slice-of-life/" title="View All Slice of Life Related Series" >Slice of Life</a></div>-- All the reunions with the person you had a crush on in the past are long-cherished schemes.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
Ruan Yu’s web novel on Jinjiang, “Really Want to Nibble on Your Ear”, has been accused of plagiarism.<p style="margin-top:-5px;"></p>
She posted on Weibo: God knows, this story about a crush is from my own personal experience back when I was still a student.<p style="margin-top:-5px;"></p>
The other author called someone on the phone after seeing it: “Brother, I seem…… to have found the person you have a crush on.”<p style="margin-top:-5px;"></p>
A few days later, Ruan Yu, who had been “exposed,” looked at the prototype of the male lead in her novel and kept shaking her hand: “I don’t know him, don’t know him……” <p style="margin-top:-5px;"></p>
Xu Huaisong smiled with gritted teeth: You’ve already nibbled my ear. It’s too late to play dumb.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_53052.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid53052" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/cannon-fodder-fake-master-was-stunned-after-being-reborn/">Cannon Fodder Fake Master Was Stunned After Being Reborn</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 24 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 2.2 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 1143 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 10 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="9" href="https://www.novelupdates.com/genre/drama/" title="View All Drama Related Series" >Drama</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="7" href="https://www.novelupdates.com/genre/slice-of-life/" title="View All Slice of Life Related Series" >Slice of Life</a> <a class="gennew search" gid="16" href="https://www.novelupdates.com/genre/supernatural/" title="View All Supernatural Related Series" >Supernatural</a> <a class="gennew search" gid="560" href="https://www.novelupdates.com/genre/yaoi/" title="View All Yaoi Related Series" >Yaoi</a></div>It was only after Jian Xingsui’s death that he realized that he was the cannon fodder in a book with a dog-blooded plot. He was a dumb and s*upid fake young master in it.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
​In the book, he was reluctant to return to his poor ”biological” parents. He was greedy of the Jian Family’s wealth and also wanted to suppress the protagonist shou of the book, the real young master. With the tricks that he had personally thought out of his pig-headed brain, he finally succeeded in making everyone in the family annoy the hell out of him:<p style="margin-top:-5px;"></p>
His former CEO elder brother: “I don’t have a younger brother like you.”<p style="margin-top:-5px;"></p>
His former star second brother: “Never mention to anyone in the entertainment industry that I used to be your brother!!!”<p style="margin-top:-5px;"></p>
His former father: “Get out, I’m so disappointed in you.”<p style="margin-top:-5px;"></p>
​​After that, he was kicked out of the house and had a terrible ending.<p style="margin-top:-5px;"></p>
After the rebirth, just when it was the day when the protagonist shou came up to his door, instead of waiting for his family to talk to him, he took the initiative: “You guys won’t have to say anything, I am willing to leave the house!”
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_48680.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.1)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid48680" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/dressed-as-the-school-grass-ex-boyfriend/">Dressed as The School Grass&#8217; Ex-boyfriend</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 77 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 2 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 1091 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 5 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="17" href="https://www.novelupdates.com/genre/comedy/" title="View All Comedy Related Series" >Comedy</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="6" href="https://www.novelupdates.com/genre/school-life/" title="View All School Life Related Series" >School Life</a> <a class="gennew search" gid="560" href="https://www.novelupdates.com/genre/yaoi/" title="View All Yaoi Related Series" >Yaoi</a></div>Jing Ci: Math problems are not easy to do, and aren’t exams not fun? So why not fall in love?<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
No interest, impossible, time-consuming. <p style="margin-top:-5px;"></p>
The school bully’s arrogant words—-<p style="margin-top:-5px;"></p>
“Go bother someone else if you want to talk about love. I’m not interested.”<p style="margin-top:-5px;"></p>
“Just kidding, what kind of thing is Jing Ci? You think dad will look at him more?” <p style="margin-top:-5px;"></p>
Later–<p style="margin-top:-5px;"></p>
“You see the one that scored first place in the exam? That’s my boyfriend.” <p style="margin-top:-5px;"></p>
“Come on, Jing Ci, which one do you choose? Mathematics or me.”<p style="margin-top:-5px;"></p>
Later–<p style="margin-top:-5px;"></p>
The school bully pressed Jing Ci against the wall: “Good, say something nice, and I will let you go.”
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/noimagemid.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid29237" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/heros-redo-a-hero-that-once-saved-the-world-reborn-as-a-girl/">Hero&#8217;s Redo ~ A Hero That Once Saved the World Reborn as a Girl ~</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 145 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 1.2 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 2169 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 11 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="168" href="https://www.novelupdates.com/genre/gender-bender/" title="View All Gender Bender Related Series" >Gender Bender</a> <a class="gennew search" gid="7" href="https://www.novelupdates.com/genre/slice-of-life/" title="View All Slice of Life Related Series" >Slice of Life</a></div>The Hero was traumatized by politics and people afraid of him for being too powerful, so the mage in his party decided to turn him<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"> into a girl. But he was so traumatized to interact with humans that he just hid inside the mountains until two girls who came to the mountain were caught in danger. After saving them, they decided to bring her to their town to pay her back for helping them.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_25273.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgjp" style="font-weight: bold;">JP</span> (3.8)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid25273" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/shangrila-frontier-shitty-games-hunter-challenges-godly-game/">ShangriLa Frontier ~ Shitty Games Hunter Challenges Godly Game ~</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 390 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 2 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 1367 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 5 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="11" href="https://www.novelupdates.com/genre/sci-fi/" title="View All Sci-fi Related Series" >Sci-fi</a> <a class="gennew search" gid="12" href="https://www.novelupdates.com/genre/shounen/" title="View All Shounen Related Series" >Shounen</a></div>For every God out there, there are about thousand shitty games in this world.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
Bugs, Errors, Texture Collapses, Inconsistent Scenarios…… Those are the things that fill the players around the world with grief and remorse.<p style="margin-top:-5px;"></p>
A certain boy who loves such shitty games decides to challenge a Godly game recognized by the general public for a change.<p style="margin-top:-5px;"></p>
As a result, both the gaming world and the real world in which the boy is revolving begin to change. Still, the specs of the Godly game still fills the boy’s heart with dread.<p style="margin-top:-5px;"></p>
「Are Encounter Rates like that really a common sense in Godly Games……?」
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_53511.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (4.4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid53511" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/after-becoming-the-alpha-protagonist-i-snatched-the-cannon-fodder-omega/">After Becoming the Alpha Protagonist, I Snatched the Cannon Fodder Omega</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 6 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 2.2 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 546 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 3 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="10" href="https://www.novelupdates.com/genre/mecha/" title="View All Mecha Related Series" >Mecha</a> <a class="gennew search" gid="15" href="https://www.novelupdates.com/genre/romance/" title="View All Romance Related Series" >Romance</a> <a class="gennew search" gid="6" href="https://www.novelupdates.com/genre/school-life/" title="View All School Life Related Series" >School Life</a> <a class="gennew search" gid="11" href="https://www.novelupdates.com/genre/sci-fi/" title="View All Sci-fi Related Series" >Sci-fi</a> <a class="gennew search" gid="1692" href="https://www.novelupdates.com/genre/shounen-ai/" title="View All Shounen Ai Related Series" >Shounen Ai</a> <a class="gennew search" gid="560" href="https://www.novelupdates.com/genre/yaoi/" title="View All Yaoi Related Series" >Yaoi</a></div>The thick-skinned Qiu Zhenyang just transmigrated, and became the crazy, cool alpha protagonist gong in an interstellar ABO campus story.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
He silently looked at the young man glaring at him with misty eyes. This young man was the cannon fodder omega the original owner had recently broken the marriage contract with. The pheromone overflowing all over his body was indicating he was about to enter estrus.<p style="margin-top:-5px;"></p>
Do or not do?<p style="margin-top:-5px;"></p>
This is a question worth thinking about...<p style="margin-top:-5px;"></p>
---------<p style="margin-top:-5px;"></p>
Transmigrating into the young master of the Federation's Chief Legion, the First Consortium, and a Pharmacy Family, Qiu Zhenyang has whatever he wants. Whoever is not pleasing to the eye, he can deal with.<p style="margin-top:-5px;"></p>
On his first day in the Interstellar Academy, facing the confession of the omega school flower: Sorry, you aren't my type.<p style="margin-top:-5px;"></p>
Inadvertently, hero saves the beauty. Facing the protagonist shou in the original text: Go away, I don't drink green tea.<p style="margin-top:-5px;"></p>
Besieged at home, facing the pursuit of the coquettish childhood sweetheart: Why are you wearing such showy clothes?<p style="margin-top:-5px;"></p>
Qiu Zhenyang's pheromone tastes like sunshine, and his personality is like a little sun. But only when facing Ling Mu, the cannon fodder covered in thorns, did heart soften again and again. His affection grew stronger, and he loved him uncontrollably.<p style="margin-top:-5px;"></p>
---------<p style="margin-top:-5px;"></p>
After the marriage contract with the Qiu family's young master was called off, Ling Mu's difficult life became even worse.<p style="margin-top:-5px;"></p>
But, just over a month later, he had an affair with the young master, and was even temporarily marked!<p style="margin-top:-5px;"></p>
That bastard!<p style="margin-top:-5px;"></p>
Qiu Zhenyang, who temporarily marked him: "I will be responsible for you!"<p style="margin-top:-5px;"></p>
Ling Mu: "Shut up, go away!"<p style="margin-top:-5px;"></p>
Qiu Zhenyang, who officially introduced Ling Mu to his friends: "This is my wife!"<p style="margin-top:-5px;"></p>
Ling Mu: "What wife, go away!"<p style="margin-top:-5px;"></p>
Qiu Zhenyang, who had just rejected the protagonist's confession: "Oh, are you jealous?"<p style="margin-top:-5px;"></p>
Ling Mu: "F*ck, go away!"<p style="margin-top:-5px;"></p>
Qiu Zhenyang: "Are you serious? Then... Xiao Mumu, let's do something harmonious?"<p style="margin-top:-5px;"></p>
Ling Mu: "..."<p style="margin-top:-5px;"></p>
After a few seconds, Qiu Zhenyang: "Ow! Wife, your fists are so powerful!"<p style="margin-top:-5px;"></p>
---------<p style="margin-top:-5px;"></p>
My pheromone is sunshine, your pheromone is sunflower, we are born to fit together.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div><div style="clear: both;"></div><div class="search_main_box_nu" dp="yes"><div class="search_img_nu" dp="yes"><img dp="yes" src="https://cdn.novelupdates.com/imgmid/series_19687.jpg" /><div dp="yes" class="search_stars"><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i><i class="fa fa-star-o"></i></div>
			<div dp="yes" class="search_ratings"><span class="orgcn" style="font-weight: bold;">CN</span> (3.4)</div>
			</div><div class="search_body_nu"><div class="search_title"><span id="sid19687" class="rl_icons_en"></span><a href="https://www.novelupdates.com/series/the-sovereigns-ascension/">The Sovereign&#8217;s Ascension</a></div><div class="search_stats"><span class="ss_desk"><i title="Chapter Count" style="padding-left: 0px;" class="fa fa-list-alt pad" aria-hidden="true"></i> 1083 Chapters</span><span class="ss_desk"><i title="Updates Frequency" class="fa fa-bolt pad" aria-hidden="true"></i> Every 0.5 Day(s)</span><span class="ss_desk"><i title="Readers" class="fa fa-user-o pad" aria-hidden="true"></i> 1351 Readers</span><span class="ss_desk"><i title="Reviews" class="fa fa-pencil-square-o pad" aria-hidden="true"></i> 14 Reviews</span><span class="ss_desk"><i title="Last Updated" class="fa fa-calendar pad" aria-hidden="true"></i> 04-07-2022</span></div><div class="search_genre"><a class="gennew search" gid="8" href="https://www.novelupdates.com/genre/action/" title="View All Action Related Series" >Action</a> <a class="gennew search" gid="13" href="https://www.novelupdates.com/genre/adventure/" title="View All Adventure Related Series" >Adventure</a> <a class="gennew search" gid="5" href="https://www.novelupdates.com/genre/fantasy/" title="View All Fantasy Related Series" >Fantasy</a> <a class="gennew search" gid="3" href="https://www.novelupdates.com/genre/harem/" title="View All Harem Related Series" >Harem</a> <a class="gennew search" gid="3954" href="https://www.novelupdates.com/genre/xuanhuan/" title="View All Xuanhuan Related Series" >Xuanhuan</a></div>As a lawyer back on earth, Lin Yun had never lost a case. He owed his success to three things: two blessings he received at birth (a photographic memory and the ability to comprehend anything he studies) and an indomitable will he forged himself.<span class="dots">... </span><span class="morelink list" href="#" onclick="showtext(this); return false;">more>></span><span class="testhide" style="display:none"><p style="margin-top:-5px;"></p>
While on a trip in Shandong province, he decided he would pay a visit to Mt. Tai. Just as he was cresting the peak, he felt a sharp pain in his chest and his vision went black.<p style="margin-top:-5px;"></p>
Upon waking up, he found himself in the world of Profound Amber occupying the body of a sword s*ave who had shared his name. The last thing he remembered before dying was the image of a sword piercing his chest.<p style="margin-top:-5px;"></p>
Through the memories his body retained the sword s*ave’s life, Lin Yun came to understand the brutality of this world. If he sought respect, he would have to earn it through strength. The weak found no compassion here.<p style="margin-top:-5px;"></p>
Refusing to leave his fate in the hands of others, Lin Yun set out to become a sovereign. No man or beast would stop him from achieving his destiny.<p style="margin-top:-5px;"></p>
With his sword in hand, he would overcome any obstacle.
<span class="morelink list" href="#" onclick="hidetext(this); return false;"> &lt;&lt;less</span></span></div></div>               								</div>
							</div>
                                
			<div class="digg_pagination" style=""><em class="current">1</em><a href="//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc&pg=2">2</a><a href="//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc&pg=3">3</a><span class="my_popup_a_open">…</span><a href="//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc&pg=496">496</a><a href="//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc&pg=2" rel="next" class="next_page"> →</a></div><div id="my_popup_b">
  	<div class="page_digg_bg">
  	<div class="page_digg_title">Go to Page...</div>
    <input type="text" name="p_num_b" id="p_num_b" onkeydown="if (event.keyCode == 13) { document.getElementById('mypu_b_button').click(); }">
    <button onclick="gotoPage('//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc&pg=xxxxx','1');" class="my_popup_b_button" id="mypu_b_button">Go</button>
    </div>
  </div>
  
    <div id="my_popup_a">
  	<div class="page_digg_bg">
  	<div class="page_digg_title">Go to Page...</div>
    <input type="text" name="p_num_a" id="p_num_a" onkeydown="if (event.keyCode == 13) { document.getElementById('mypu_a_button').click(); }">
    <button onclick="gotoPage('//www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=desc&pg=xxxxx','2');" class="my_popup_a_button" id="mypu_a_button">Go</button>
    </div>
  </div>        

<script>
    $(document).ready(function() {
     // Initialize the plugin
$('.fa.fa-info-circle.sfinder').on({
    click: function(event) {
        $('#my_popupsfinder').popup({
            tooltipanchor: event.target,
			vertical: 'top',
            autoopen: true,
            type: 'tooltip'
        });
    }
});
    });
</script>
  
  <div id="my_popupsfinder" style="display:none;">
  <div class="pop-title">Title</div>
  <div class="pop-content">...popup content...</div>
  <div class="pop-footer"><button class="my_popupsfinder_close">OK</button></div>
  <div class="arrow_down_sfinder"></div>
  </div>
  
  <span class="popcontent_noveltype" style="display:none;">The type of novel.</span>
  <span class="popcontent_language" style="display:none;">The language of the novel.</span>
  
  <span class="popcontent_releases" style="display:none;">This is the number of releases (chapters).</br></br>
  <ul>
  <li><b>Min</b> - Minimum amount of chapters.</li>
  <li><b>Max</b>- Maximum amount of chapters.</li>
 </ul></span>
 
  <span class="popcontent_frequency" style="display:none;">The release frequency of a novel. Higher frequency means the novel is updated more often.</br></br>
  <ul>
  <li><b>Min</b> - Minimum release frequency. Lower frequency = slower updates.</li>
  <li><b>Max</b>- Maximum release frequency. Higher frequency = quicker updates.</li>
 </ul></span>
 
   <span class="popcontent_rating" style="display:none;">Novel rating on a scale of 1 to 5.</br></br>
  <ul>
  <li><b>Min</b> - Minimum rating (1 to 5)</li>
  <li><b>Max</b> - Maximum rating (1 to 5)</li>
 </ul></span>
 
   <span class="popcontent_numberratings" style="display:none;">The amount of ratings for a novel.</br></br>
  <ul>
  <li><b>Min</b> - Minimum number of ratings.</li>
  <li><b>Max</b> - Maximum number of ratings.</li>
 </ul></span>
 
    <span class="popcontent_readers" style="display:none;">The number of readers a novel has.</br></br>
  <ul>
  <li><b>Min</b> - Minimum number of readers.</li>
  <li><b>Max</b> - Maximum number of readers.</li>
 </ul></span>
 
   <span class="popcontent_lastdate" style="display:none;">The last release date.</span>
   
   <span class="popcontent_genre" style="display:none;">The genres of the novel.</br></br>
     <ul>
  <li><b>Include</b> - Click a genre once to include it in your search.</li>
  <li><b>Exclude</b> - Click a genre twice to exclude it in your search.</li>
   <li><b>AND</b> - Matches ALL genres selected.</li>
  <li><b>OR</b> - Matches ANY of the genres selected.</li>
 </ul></span>
 
   <span class="popcontent_tags" style="display:none;">The tags of the novel.</br></br>   <ul>
   <li><b>AND</b> - Matches ALL tags selected.</li>
  <li><b>OR</b> - Matches ANY of the tags selected.</li>
 </ul></span>
 
   <span class="popcontent_readinglists" style="display:none;">Your reading lists are listed here.</br></br><ul>
   <li><b>Include</b> - Searches only your reading lists.</li>
  <li><b>Exclude</b> - Exclude series that matches your reading lists.</li>
 </ul> </span>
 
   <span class="popcontent_complete" style="display:none;">Completely translated option.</span>
   <span class="popcontent_sorting" style="display:none;">You can sort results by selecting the options from the drop down list.</span>
<script>
$(document).ready(function() 
    { 
		$('#sortorder option[value="desc"]').prop('selected', true);
$('#sortresults option[value="sdate"]').prop('selected', true);
		$(".sortresults").chosen({disable_search_threshold: 20,width: "90%"});
		$(".sortorder").chosen({disable_search_threshold: 20,width: "90%"});
		$(".storystatus").chosen({disable_search_threshold: 20,width: "40%"});
		$(".ctranslate").chosen({width: "40%"});
		$("#ardate").datepicker();
		$("#ardate_first").datepicker();
		
		$(".tablesorter").tablesorter(); 
		
		$(".chzn-select").chosen({width: "90%", search_contains:true});$("#hdlist").chosen({width: "95%"});	} 
); 
</script>
<script>

function getgrouptags()
{
		var strType = '';
		
		$('#tags_include').empty();
		$('#tags_exclude').empty();
		
		$("#tags_include").attr('disabled', true).trigger("chosen:updated")
		$("#tags_exclude").attr('disabled', true).trigger("chosen:updated")
		
		var chkIcon = $('.tag_icon').attr('current');
		
		if ( chkIcon == 'org' )
		{
			$('.fa.fa-tag').addClass('fa-spin');
			strType = 'tag';
		}
		else
		{
		   	$('.fa.fa-undo').addClass('fa-spin');
			strType = 'undo';
		}
		
	
        $.ajax({
            type:"POST",
            url: "https://www.novelupdates.com/wp-admin/admin-ajax.php",
            data: {
                action:'nd_gettags',
				strType : strType
            },
            success:function(response){
				response = response.slice(0,-1);


				$("#tags_include").append(response);		
				$('#tags_include').trigger("chosen:updated");
				$("#tags_exclude").append(response);
				$('#tags_exclude').trigger("chosen:updated");
			
				$("#tags_include").attr('disabled', false).trigger("chosen:updated")
				$("#tags_exclude").attr('disabled', false).trigger("chosen:updated")
			
				if ( chkIcon == 'org' )
				{
					$('.fa.fa-tag').removeClass('fa-spin');
					$('.tag_icon').replaceWith('<span class="tag_icon" current="extra"><i style="cursor:pointer;" onclick="getgrouptags();" class="fa fa-undo" aria-hidden="true"></i></span>');
				}
				else
				{
					$('.fa.fa-undo').removeClass('fa-spin');
					$('.tag_icon').replaceWith('<span class="tag_icon" current="org"><i style="cursor:pointer;" onclick="getgrouptags();" class="fa fa-tag" aria-hidden="true"></i></span>');
				}
				
            }}); 
}

rl_open_start();
function showtext(n){$(n).parent().find("span").show(),$(n).parent().find(".dots").hide(),$(n).hide()}function hidetext(n){$(n).parent().parent().find(".testhide").hide(),$(n).parent().parent().find(".dots").show(),$(n).parent().parent().find("span.morelink").show()}

</script>	
			</div>

			<div class="l-sidebar at_right">
				

<!-- begin generated sidebar -->
<div id="text-2" class="widget widget_text"><h3 class="widgettitle">Novel Updates</h3>			<div class="textwidget"><ul>
<li><a href="https://forum.novelupdates.com/">Forum</a></li>
<li><a href="https://www.novelupdates.com/random-novel/">Random Novel</a></li>
<li><a href="https://www.novelupdates.com/series-finder/">Series Finder</a></li>
<li><a href="https://www.novelupdates.com/novelslisting/">Series Listing</a></li>
<li><a href="https://www.novelupdates.com/series-ranking/">Series Ranking</a></li>
<li><a href="https://www.novelupdates.com/latest-series/">Latest Series</a></li>
<li><a href="https://www.novelupdates.com/recommendation-lists/">Rec Lists</a></li>
</ul></div>
		</div><div id="text-3" class="widget widget_text"><h3 class="widgettitle">User Tools</h3>			<div class="textwidget"><ul>
<li><a href="https://www.novelupdates.com/reading-list/">Reading List</a></li>
<li><a href="https://www.novelupdates.com/release-filtering/">Release Filtering</a></li>
<li><a href="https://www.novelupdates.com/series-filtering/">Series Filtering</a></li>
</ul></div>
		</div><style>.search_img { float: left;}.search_img img { height: 71px; width: 50px; }</style><div style="margin-bottom: 25px;"><div class="search_main_box"><div class="search_img">
			<a href="https://www.scribblehub.com/series/471295/a-dragon-idols-reincarnation-tale/"><img src="https://cdn.scribblehub.com/seriesimg/mid/23/mid_471295.jpg"></a>
			</div><div class="search_body" style="line-height: 20px; margin-left: 55px;"><div style="font-size:14px; position: relative; top: -3px; max-height: 40px; overflow: hidden;"><a href="https://www.scribblehub.com/series/471295/a-dragon-idols-reincarnation-tale/">A Dragon Idol's Reincarnation Tale</a></div><div><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i></div></div></div><div style="clear: both;"></div><div class="search_main_box"><div class="search_img">
			<a href="https://www.scribblehub.com/series/278470/retribution-engine-arc-/"><img src="https://cdn.scribblehub.com/seriesimg/mid/13/mid_278470.jpg"></a>
			</div><div class="search_body" style="line-height: 20px; margin-left: 55px;"><div style="font-size:14px; position: relative; top: -3px; max-height: 40px; overflow: hidden;"><a href="https://www.scribblehub.com/series/278470/retribution-engine-arc-/">Retribution Engine ARC 2</a></div><div><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div></div></div><div style="clear: both;"></div><div class="search_main_box"><div class="search_img">
			<a href="https://www.scribblehub.com/series/366931/the-hero-has-many-flaws-a-journey-in-another-world/"><img src="https://cdn.scribblehub.com/seriesimg/mid/18/mid_366931.jpg"></a>
			</div><div class="search_body" style="line-height: 20px; margin-left: 55px;"><div style="font-size:14px; position: relative; top: -3px; max-height: 40px; overflow: hidden;"><a href="https://www.scribblehub.com/series/366931/the-hero-has-many-flaws-a-journey-in-another-world/">The Hero has many flaws, a...</a></div><div><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star-half-o"></i></div></div></div><div style="clear: both;"></div></div><h3 class="widgettitle_nuf" style="margin-bottom: 5px;">Latest Discussions</h3><div class="nuf_bd"><ul id="nuf_ul"><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="Regressor Instruction Manual" class="nuftrend_link" href="http://forum.novelupdates.com/threads/120915/unread">Regressor Instruction Manual</a></div><div class="additionalRow muted">4h, 14m ago</div></div></li><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="The Villain Wants to Live" class="nuftrend_link" href="http://forum.novelupdates.com/threads/134041/unread">The Villain Wants...</a></div><div class="additionalRow muted">4h, 23m ago</div></div></li><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="Clearing an Isekai with the Zero-Believers Goddess – The Weakest Mage among the Classmates (WN)" class="nuftrend_link" href="http://forum.novelupdates.com/threads/126095/unread">Clearing an Isekai...</a></div><div class="additionalRow muted">10h, 7m ago</div></div></li><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="Absolute Resonance" class="nuftrend_link" href="http://forum.novelupdates.com/threads/133372/unread">Absolute Resonance</a></div><div class="additionalRow muted">13h, 46m ago</div></div></li><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="The Protagonists Are Murdered by Me" class="nuftrend_link" href="http://forum.novelupdates.com/threads/110159/unread">The Protagonists Are...</a></div><div class="additionalRow muted">16h, 0m ago</div></div></li><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="The Death Mage Who Doesn&#8217;t Want a Fourth Time" class="nuftrend_link" href="http://forum.novelupdates.com/threads/91000/unread">The Death Mage...</a></div><div class="additionalRow muted">23h, 44m ago</div></div></li><li id="nuf_li"><div class="messageInfo"><div class="nuf_msgcontent"><a title="Possessing Nothing" class="nuftrend_link" href="http://forum.novelupdates.com/threads/40682/unread">Possessing Nothing</a></div><div class="additionalRow muted">1d, 19h, 33m ago</div></div></li></ul></div><!-- TAGNAME: 300x600_Parent_SmartSlot --><div id='div-gpt-ad-novelupdatescom40265'></div>
<!-- end generated sidebar -->

			</div>
		</div>
	</div>



</div>
<!-- /MAIN -->

</div>
<!-- /CANVAS -->

<!-- FOOTER -->
<div class="l-footer">
			<!-- subfooter: bottom -->
	<div class="l-subfooter at_bottom">
		<div class="l-subfooter-h i-cf">

									
			

			<div class="w-copyright"><script>
(function waitGEO() {
    var readyGEO;
    if (window['UnicI'] && window['UnicI'].geo && window['UnicI'].geo !== '-' ) {
        readyGEO = true;
        if (window['UnicI'].geo === 'EU') {
            if(document.getElementById("unic-gdpr")) {
              document.getElementById("unic-gdpr").style.display = 'inline-block';
            }
        }
        if (window['UnicI'].geo === 'CA') {
            if(document.getElementById("unic-ccpa")) {
              document.getElementById("unic-ccpa").style.display = 'inline-block';
            }
        }
    }
    if (!readyGEO) {
        setTimeout(waitGEO, 200);
    }
})();
</script>

<!-- Consent links -->

<a id='unic-gdpr' onclick='__tcfapi("openunic");return false;' style='display:none;cursor:pointer;'>Change Ad Consent</a>
<a id='unic-ccpa' onclick="window.__uspapi('openunic')" style='display:none;cursor:pointer;'>Do not sell my data</a> | <a href="https://www.novelupdates.com/privacy-policy/">Privacy Policy</a> | <a href="https://www.novelupdates.com/terms-of-service/">Terms of Service</a> | <a href="https://www.novelupdates.com/contact-us/">Contact Us</a></div>
					</div>
	</div>
	
</div>
<!-- /FOOTER -->
<a class="w-toplink" href="#"><i class="fa fa-angle-up"></i></a>
<script type="text/javascript">
	if (window.$us === undefined) window.$us = {};
	$us.canvasOptions = ($us.canvasOptions || {});
	$us.canvasOptions.headerDisableStickyHeaderWidth = parseInt('1000');
	$us.canvasOptions.headerDisableAnimationWidth = parseInt('900');
	$us.canvasOptions.headerMainHeight = parseInt('60');
	$us.canvasOptions.headerMainShrinkedHeight = parseInt('60');
	$us.canvasOptions.headerExtraHeight = parseInt('36');
	$us.canvasOptions.mobileNavWidth = parseInt('1000');

	$us.navOptions = ($us.navOptions || {});
	$us.navOptions.togglable = 1;

</script>
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-66522525-1', 'auto');
  ga('send', 'pageview');

</script><script type='text/javascript' src='https://www.novelupdates.com/wp-content/plugins/user-meta-display/assets/js/scripts-user_meta_display.js?ver=1.2.2'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/plugins/yet-another-stars-rating/js/jquery.rateit.min.js?ver=1.0.22'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/plugins/yet-another-stars-rating/js/yasr-front.js?ver=1.0.1'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates/js/jquery.easing.min.js?ver=5.2.7'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates/js/us.core.js?ver=4.4.8'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-content/themes/ndupdates/js/us.widgets.js?ver=4.4.8'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-includes/js/comment-reply.min.js?ver=5.2.7'></script>
<script type='text/javascript' src='https://www.novelupdates.com/wp-includes/js/wp-embed.min.js?ver=5.2.7'></script>
</body>
</html>
`)

func TestDoSearchRequest_DoesNotReturn_Nil(t *testing.T) {

	response, err := DoSearchRequest(TestSuccessRequester, nil)

	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestDoSearchRequest_Returns_ResultsCount_As_25(t *testing.T) {

	response, _ := DoSearchRequest(TestSuccessRequester, nil)

	assert.Len(t, response, 25)
}

func TestDoSearchRequest_Parse_Title_Success(t *testing.T) {

	response, _ := DoSearchRequest(TestSuccessRequester, nil)

	assert.Equal(t, "Another World Awakening Transcendental Create Skill", response[0].Title)
	assert.Equal(t, "I’ll Become a Villainess That Will Go Down in History", response[1].Title)
	assert.Equal(t, "I Help the Richest Man Spend Money to Prevent Disasters", response[2].Title)
}

func TestDoSearchRequest_Parse_Title_Failure(t *testing.T) {

	response, _ := DoSearchRequest(TestFailureRequester, nil)

	assert.Equal(t, "ERR", response[0].Title)
	assert.Equal(t, "ERR", response[1].Title)
	assert.Equal(t, "ERR", response[2].Title)
}

func TestDoSearchRequest_Parse_ChaptersCount_Success(t *testing.T) {

	response, _ := DoSearchRequest(TestSuccessRequester, nil)

	assert.Equal(t, 100, response[0].Chapters)
	assert.Equal(t, 225, response[1].Chapters)
	assert.Equal(t, 372, response[2].Chapters)
}

func TestDoSearchRequest_Parse_ChaptersCount_Failure(t *testing.T) {

	response, _ := DoSearchRequest(TestFailureRequester, nil)

	assert.Equal(t, -1, response[0].Chapters)
	assert.Equal(t, -1, response[1].Chapters)
	assert.Equal(t, -1, response[2].Chapters)
}

func TestDoSearchRequest_Parse_UpdateFrequency_Success(t *testing.T) {

	response, _ := DoSearchRequest(TestSuccessRequester, nil)

	assert.Equal(t, 1.1, response[0].ReleaseFrequencyInDays)
	assert.Equal(t, 1.6, response[1].ReleaseFrequencyInDays)
	assert.Equal(t, 1.7, response[2].ReleaseFrequencyInDays)
}

func TestDoSearchRequest_Parse_UpdateFrequency_Failure(t *testing.T) {

	response, _ := DoSearchRequest(TestFailureRequester, nil)

	assert.Equal(t, -1.0, response[0].ReleaseFrequencyInDays)
	assert.Equal(t, -1.0, response[1].ReleaseFrequencyInDays)
	assert.Equal(t, -1.0, response[2].ReleaseFrequencyInDays)
}

func TestDoSearchRequest_Parse_Readers_Success(t *testing.T) {

	response, _ := DoSearchRequest(TestSuccessRequester, nil)

	assert.Equal(t, 1039, response[0].Readers)
	assert.Equal(t, 7819, response[1].Readers)
	assert.Equal(t, 5617, response[2].Readers)
}

func TestDoSearchRequest_Parse_Readers_Failure(t *testing.T) {

	response, _ := DoSearchRequest(TestFailureRequester, nil)

	assert.Equal(t, -1, response[0].Readers)
	assert.Equal(t, -1, response[1].Readers)
	assert.Equal(t, -1, response[2].Readers)
}

func TestDoSearchRequest_Parse_Reviews_Success(t *testing.T) {

	response, _ := DoSearchRequest(TestSuccessRequester, nil)

	assert.Equal(t, 2, response[0].Reviews)
	assert.Equal(t, 49, response[1].Reviews)
	assert.Equal(t, 83, response[2].Reviews)
}

func TestDoSearchRequest_Parse_Reviews_Failure(t *testing.T) {

	response, _ := DoSearchRequest(TestFailureRequester, nil)

	assert.Equal(t, -1, response[0].Reviews)
	assert.Equal(t, -1, response[1].Reviews)
	assert.Equal(t, -1, response[2].Reviews)
}
