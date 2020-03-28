import './index.less'

import $ from 'jquery'

console.log($('.tem').text())
$('article').html($('.tem').text())
// document.querySelector('article').appendChild(parser.parseFromString(temHtml, "text/xml"))