import './index.less'

import $ from 'jquery'
import {markdown} from 'markdown';

$('article').html(markdown.toHTML($('.tem').text()))
// document.querySelector('article').appendChild(parser.parseFromString(temHtml, "text/xml"))