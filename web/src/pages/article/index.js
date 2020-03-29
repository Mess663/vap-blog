import './index.less'

import $ from 'jquery'
import showdown from 'showdown'


const converter = new showdown.Converter();
console.log(converter.makeHtml($('.tem').text()))
$('article').html(converter.makeHtml($('.tem').text()))
// document.querySelector('article').appendChild(parser.parseFromString(temHtml, "text/xml"))