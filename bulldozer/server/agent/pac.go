package agent

import (
	"github.com/dollarkillerx/remote_networking/bulldozer/conf"
	"github.com/dollarkillerx/remote_networking/bulldozer/utils"
	"github.com/dollarkillerx/remote_networking/bulldozer/utils/ip2region"
	"log"
	"net"
	"strings"
)

var defaultPacDomain = []string{
	"google.com",
	"google.cn",
	"google.co.jp",
	"github.com",
	"githubusercontent.com",
	"bbc.com",
	"huaglad.com",
	"slideshare.net",
	"tumblr.com",
	"bbc.co.uk",
	"pinterest.com",
	"googleadsensetvsite.com",
	"googlesyndication.com",
	"googletagmanager.com",
	"googletagservices.com",
	"googleusercontent.com",
	"google-analytics.com",
	"qi-gong.me",
	"3tui.net",
	"androidebook.org",
	"androidmall.org",
	"bannednews.org",
	"bnews.co",
	"breakgfw.com",
	"hyperspaceproxy.info",
	"singlelogin.org",
	"bookos-z1.org",
	"android-x86.org",
	"thelaunchbook.com",
	"freefq.com",
	"faluninfo.net",
	"falundafaradio.org",
	"guangming.org",
	"zhengwunet.org",
	"yuanming.net",
	"99cn.info",
	"jinpianwang.com",
	"bookepub.com",
	"100ke.org",
	"dtwang.org",
	"blogspot.fr",
	"zhengjian.org",
	"shenyun.com",
	"bbc.com",
	"dw.com",
	"watchinese.com",
	"ntdtv.com",
	"live.com",
	"keepvid.com",
	"unblockdmm.com",
	"lvv2.com",
	"usembassy.gov",
	"iobit.com",
	"telegram.org",
	"abc.xyz",
	"amazonaws.com",
	"chenqiwei.com",
	"blogspot.com",
	"vpsdime.com",
	"teamviewer.com",
	"truste.com",
	"a.fsdn.com",
	"akamaihd.net",
	"cdninstagram.com",
	"namecheap.com",
	"chromium.org",
	"nexon.net",
	"nexon.com",
	"nexoneu.com",
	"nexon.co.jp",
	"konachan.com",
	"vultr.com",
	"flashfxp.com",
	"nyaa.eu",
	"nyaa.se",
	"dmhy.org",
	"jwpcdn.com",
	"jwplayer.com",
	"0to255.com",
	"123rf.com",
	"12bet.com",
	"12vpn.com",
	"17t17p.com",
	"1984bbs.com",
	"1bao.org",
	"2-hand.info",
	"moegirl.org",
	"wzyboy.im",
	"2000fun.com",
	"2008xianzhang.info",
	"21andy.com",
	"2shared.com",
	"301works.org",
	"36rain.com",
	"4bluestones.biz",
	"4chan.org",
	"4shared.com",
	"4sq.com",
	"50webs.com",
	"64tianwang.com",
	"64wiki.com",
	"666kb.com",
	"6park.com",
	"7capture.com",
	"881903.com",
	"89-64.org",
	"twitch.tv",
	"9001700.com",
	"908taiwan.org",
	"9bis.com",
	"9bis.net",
	"a-normal-day.com",
	"a5.com.ru",
	"abc.pp.ru",
	"ablwang.com",
	"aboluowang.com",
	"aculo.us",
	"addictedtocoffee.de",
	"advanscene.com",
	"aenhancers.com",
	"af.mil",
	"aiweiwei.com",
	"akiba-online.com",
	"aliengu.com",
	"alkasir.com",
	"all-that-is-interesting.com",
	"moegirl.org",
	"allaboutalpha.com",
	"allgirlsallowed.org",
	"allmovie.com",
	"alternate-tools.com",
	"altrec.com",
	"alvinalexander.com",
	"alwaysdata.com",
	"alwaysdata.net",
	"amazon.com",
	"ameblo.jp",
	"americangreencard.com",
	"amiblockedornot.com",
	"amnesty.org",
	"amnestyusa.org",
	"anchorfree.com",
	"ancsconf.org",
	"andfaraway.net",
	"android.com",
	"angularjs.org",
	"animecrazy.net",
	"anobii.com",
	"anontext.com",
	"anonymizer.com",
	"answering-islam.org",
	"antd.org",
	"anthonycalzadilla.com",
	"aol.ca",
	"aol.com",
	"aolnews.com",
	"aomiwang.com",
	"ap.org",
	"apiary.io",
	"apigee.com",
	"appledaily.com",
	"appspot.com",
	"archive.is",
	"archive.org",
	"tuo8.co",
	"areca-backup.org",
	"army.mil",
	"artsy.net",
	"asahichinese.com",
	"asdfg.jp",
	"asiaharvest.org",
	"asianews.it",
	"askstudent.com",
	"assembla.com",
	"atc.org.au",
	"atchinese.com",
	"atgfw.org",
	"allinfa.com",
	"atnext.com",
	"avaaz.org",
	"avidemux.org",
	"avoision.com",
	"awardwinningfjords.com",
	"awflasher.com",
	"axureformac.com",
	"baby-kingdom.com",
	"backchina.com",
	"backtotiananmen.com",
	"badassjs.com",
	"badoo.com",
	"baidu.jp",
	"barenakedislam.com",
	"bayvoice.net",
	"baywords.com",
	"bbc.in",
	"bbcchinese.com",
	"bbg.gov",
	"bbsland.com",
	"bebo.com",
	"beijing1989.com",
	"beijingspring.com",
	"benjaminste.in",
	"berlintwitterwall.com",
	"bestvpnservice.com",
	"bet365.com",
	"betfair.com",
	"bettween.com",
	"betvictor.com",
	"beyondfirewall.com",
	"bignews.org",
	"bigsound.org",
	"bill2-software.com",
	"bipic.net",
	"birdhouseapp.com",
	"bit.ly",
	"bitcointalk.org",
	"bitly.com",
	"bitshare.com",
	"bjzc.org",
	"blinkx.com",
	"blip.tv",
	"blog.de",
	"blogcatalog.com",
	"blogger.com",
	"bloglovin.com",
	"blogs.com",
	"blogtd.org",
	"bloodshed.net",
	"bloomberg.cn",
	"bloomberg.com",
	"bloomberg.de",
	"bnrmetal.com",
	"boardreader.com",
	"bobulate.com",
	"bonjourlesgeeks.com",
	"bot.nu",
	"botanwang.com",
	"bowenpress.com",
	"box.net",
	"boxcar.io",
	"boxun.com",
	"boxunblog.com",
	"br.st",
	"bralio.com",
	"braumeister.org",
	"break.com",
	"breakwall.net",
	"brightkite.com",
	"brizzly.com",
	"brucewang.net",
	"budaedu.org",
	"bugclub.org",
	"bullog.org",
	"bullogger.com",
	"businessinsider.com.au",
	"businesstimes.com.cn",
	"businessweek.com",
	"bx.tl",
	"c-spanvideo.org",
	"provideocoalition.com",
	"stupidvideos.com",
	"cactusvpn.com",
	"cafepress.com",
	"calameo.com",
	"calebelston.com",
	"cams.org.sg",
	"canadameet.com",
	"cantonese.asia",
	"canyu.org",
	"caobian.info",
	"caochangqing.com",
	"catch22.net",
	"cbsnews.com",
	"ccim.org",
	"cclife.org",
	"ccthere.com",
	"cctongbao.com",
	"ccue.ca",
	"ccue.com",
	"cdjp.org",
	"cdp1998.org",
	"cdp2006.org",
	"cdpwu.org",
	"cdw.com",
	"cecc.gov",
	"cellulo.info",
	"cenci.tk",
	"cenews.eu",
	"centralnation.com",
	"centurys.net",
	"cftfc.com",
	"cgdepot.org",
	"chandoo.org",
	"change.org",
	"changp.com",
	"chapm25.com",
	"chengmingmag.com",
	"chenpokong.com",
	"cherrysave.com",
	"chicagoncmtv.com",
	"china-week.com",
	"china101.com",
	"china21.org",
	"china5000.us",
	"chinaaffairs.org",
	"chinaaid.net",
	"chinaaid.org",
	"chinachange.org",
	"chinacomments.org",
	"chinadigitaltimes.net",
	"chinagate.com",
	"chinagfw.org",
	"chinahush.com",
	"chinainperspective.com",
	"chinalawandpolicy.com",
	"chinalawtranslate.com",
	"chinarightsia.org",
	"chinasoul.org",
	"chinatimes.com",
	"chinaworker.info",
	"chinese-memorial.org",
	"chinesen.de",
	"chinesepen.org",
	"chinesetalks.net",
	"chosun.com",
	"chrispederick.com",
	"christianstudy.com",
	"chrome.com",
	"citizenlab.org",
	"citizensradio.org",
	"civilhrfront.org",
	"cjb.net",
	"cl.ly",
	"classicalguitarblog.net",
	"clientsfromhell.net",
	"clipfish.de",
	"cmoinc.org",
	"cms.gov",
	"cnd.org",
	"cnn.com",
	"cnyes.com",
	"codeboxapp.com",
	"codeshare.io",
	"collateralmurder.com",
	"collateralmurder.org",
	"comedycentral.com",
	"compileheart.com",
	"contactmagazine.net",
	"convio.net",
	"coolaler.com",
	"coolder.com",
	"corumcollege.com",
	"cotweet.com",
	"cpj.org",
	"crackle.com",
	"crd-net.org",
	"creaders.net",
	"csdparty.com",
	"csuchen.de",
	"cubicle17.com",
	"cuhkacs.org",
	"cuihua.org",
	"curvefish.com",
	"cyberctm.com",
	"cyberghostvpn.com",
	"d0z.net",
	"dabr.mobi",
	"dadazim.com",
	"dafahao.com",
	"dailyme.com",
	"dailymotion.com",
	"dajiyuan.com",
	"dajiyuan.eu",
	"dalailama.com",
	"dalailama.ru",
	"dalailamaworld.com",
	"dalianmeng.org",
	"danke4china.net",
	"danwei.org",
	"daolan.net",
	"darpa.mil",
	"date.fm",
	"davidslog.com",
	"davidziegler.net",
	"dayabook.com",
	"daylife.com",
	"dayoneapp.com",
	"de-sci.org",
	"debian.org",
	"deck.ly",
	"delcamp.net",
	"democrats.org",
	"desc.se",
	"deutsche-welle.de",
	"dev102.com",
	"deviantart.com",
	"deviantart.net",
	"devio.us",
	"dfanning.com",
	"dfas.mil",
	"diaoyuislands.org",
	"digg.com",
	"diigo.com",
	"directcreative.com",
	"disp.cc",
	"dit-inc.us",
	"djangosnippets.org",
	"dmcdn.net",
	"dns2go.com",
	"dnscrypt.org",
	"dongtaiwang.com",
	"dongtaiwang.net",
	"dontfilter.us",
	"dotsub.com",
	"doubleaf.com",
	"dougscripts.com",
	"dowei.org",
	"doxygen.org",
	"dphk.org",
	"drewolanoff.com",
	"nicovideo.jp",
	"drgan.net",
	"dribbble.com",
	"dropbox.com",
	"dropboxusercontent.com",
	"drsunacademy.com",
	"dtic.mil",
	"dtiserv.com",
	"duckduckgo.com",
	"duckload.com",
	"duihua.org",
	"duihuahrjournal.org",
	"duping.net",
	"duplicati.com",
	"dupola.com",
	"dupola.net",
	"dvorak.org",
	"dw-world.com",
	"dw-world.de",
	"dw.de",
	"dwheeler.com",
	"dwnews.com",
	"dynawebinc.com",
	"dyndns.org",
	"e-gold.com",
	"eamonnbrennan.com",
	"ebookee.com",
	"echofon.com",
	"ecministry.net",
	"ecstart.com",
	"edgecastcdn.net",
	"edicypages.com",
	"edoors.com",
	"edubridge.com",
	"efksoft.com",
	"efmoe.club",
	"electionsmeter.com",
	"emacsblog.org",
	"emory.edu",
	"emuparadise.me",
	"engadget.com",
	"epochtimes-bg.com",
	"epochtimes-romania.com",
	"epochtimes.co.kr",
	"epochtimes.com",
	"epochtimes.de",
	"epochtimes.fr",
	"epochtimes.ie",
	"epochtimes.jp",
	"epochtimes.ru",
	"epochtimes.se",
	"epochtimestr.com",
	"epochweekly.com",
	"erabaru.net",
	"erepublik.com",
	"erights.net",
	"eriversoft.com",
	"ernestmandel.org",
	"etaiwannews.com",
	"ettoday.net",
	"eventful.com",
	"everyday-carry.com",
	"excite.co.jp",
	"expatshield.com",
	"ezpeer.com",
	"facebook.com",
	"facebook.net",
	"facesofnyfw.com",
	"faiththedog.info",
	"falsefire.com",
	"falunart.org",
	"falunasia.info",
	"falundafa.org",
	"falundafamuseum.org",
	"falunhr.org",
	"fanglizhi.info",
	"fangong.org",
	"fangongheike.com",
	"fanqianghou.com",
	"fanswong.com",
	"fanyue.info",
	"farwestchina.com",
	"fastpic.ru",
	"faststone.org",
	"favstar.fm",
	"fawanghuihui.org",
	"faydao.com",
	"fb.com",
	"fb.me",
	"fbcdn.net",
	"fbsbx.com",
	"feedburner.com",
	"feministteacher.com",
	"ff.im",
	"fflick.com",
	"fgmtv.net",
	"fgmtv.org",
	"filefactory.com",
	"fileserve.com",
	"fillthesquare.org",
	"firstfivefollowers.com",
	"flecheinthepeche.fr",
	"flickr.com",
	"flickrhivemind.net",
	"flnet.org",
	"fly4ever.me",
	"focusvpn.com",
	"fofg.org",
	"foolsmountain.com",
	"fooooo.com",
	"forum4hk.com",
	"forums-free.com",
	"fotop.net",
	"foxbusiness.com",
	"freakshare.com",
	"fredwilson.vc",
	"free-gate.org",
	"free-hada-now.org",
	"free-ssh.com",
	"free.fr",
	"freealim.com",
	"freedomhouse.org",
	"freegao.com",
	"freelotto.com",
	"freeman2.com",
	"50megs.com",
	"freenetproject.org",
	"freeoz.org",
	"freerk.com",
	"freetibet.org",
	"freewallpaper4.me",
	"freewebs.com",
	"freeweibo.com",
	"friendfeed.com",
	"fring.com",
	"frommel.net",
	"frontlinedefenders.org",
	"fscked.org",
	"ftchinese.com",
	"fuckgfw.org",
	"funp.com",
	"furinkan.com",
	"futureme.org",
	"fw.cm",
	"fxnetworks.com",
	"fzh999.net",
	"g.co",
	"gabocorp.com",
	"game735.com",
	"ganges.com",
	"gaoming.net",
	"gaozhisheng.net",
	"gardennetworks.com",
	"gardennetworks.org",
	"gartlive.com",
	"canton8.com",
	"genuitec.com",
	"geocities.co.jp",
	"geocities.com",
	"geocities.jp",
	"geohot.com",
	"geometrictools.com",
	"get-digital-help.com",
	"getcloudapp.com",
	"getfoxyproxy.org",
	"getjetso.com",
	"getlantern.org",
	"ggssl.com",
	"ghost.org",
	"ghostery.com",
	"ghut.org",
	"giganews.com",
	"gimpshop.com",
	"git-scm.com",
	"glennhilton.com",
	"globaljihad.net",
	"globalmuseumoncommunism.org",
	"globalvoicesonline.org",
	"gmail.com",
	"gmodules.com",
	"goagent.biz",
	"goagentplus.com",
	"golang.org",
	"goldbetsports.com",
	"goldwave.com",
	"gongm.in",
	"gongwt.com",
	"goo.gl",
	"goodreads.com",
	"google.co.id",
	"google.co.jp",
	"google.co.kr",
	"google.co.uk",
	"google.ca",
	"google.cn",
	"google.de",
	"google.fr",
	"google.it",
	"google.nl",
	"google.com",
	"google.com.au",
	"google.com.hk",
	"google.com.my",
	"google.com.tw",
	"googleadservices.com",
	"googleapis.com",
	"googlecode.com",
	"googledomains.com",
	"googledrive.com",
	"googleearth.com",
	"googlehosted.com",
	"googlelabs.com",
	"googlemail.com",
	"googlepages.com",
	"googleplus.com",
	"googlesource.com",
	"googlevideo.com",
	"gopetition.com",
	"gospelherald.com",
	"gotw.ca",
	"gowalla.com",
	"gradconnection.com",
	"grandtrial.org",
	"gravatar.com",
	"graylog2.org",
	"greatfire.org",
	"greatfirewall.biz",
	"greatfirewallofchina.org",
	"greatzhonghua.org",
	"gstatic.com",
	"gtricks.com",
	"guishan.org",
	"gunsamerica.com",
	"gyalwarinpoche.com",
	"gzone-anime.info",
	"h-china.org",
	"hacken.cc",
	"hdtvb.net",
	"heartyit.com",
	"helloandroid.com",
	"hellonewyork.us",
	"hellouk.org",
	"helplinfen.com",
	"heqinglian.net",
	"heywire.com",
	"hideipvpn.com",
	"hidemyass.com",
	"hikinggfw.org",
	"hinet.net",
	"hkbf.org",
	"hkchurch.org",
	"hkday.net",
	"hkej.com",
	"hkepc.com",
	"hkfront.org",
	"hkheadline.com",
	"hkhkhk.com",
	"hkjc.com",
	"hkjp.org",
	"hkptu.org",
	"hnjhj.com",
	"hola.com",
	"holyspiritspeaks.org",
	"homeservershow.com",
	"honeynet.org",
	"hongzhi.li",
	"hootsuite.com",
	"hotspotshield.com",
	"howtoforge.com",
	"hqcdp.org",
	"hrichina.org",
	"hrw.org",
	"hsjp.net",
	"hsselite.com",
	"ht.ly",
	"htl.li",
	"htmldog.com",
	"huanghuagang.org",
	"huaxia-news.com",
	"hudatoriq.web.id",
	"huhaitai.com",
	"huhamhire.com",
	"hulu.com",
	"huping.net",
	"hwinfo.com",
	"hyperrate.com",
	"i-cable.com",
	"i2p2.de",
	"ialmostlaugh.com",
	"ibiblio.org",
	"iblogserv-f.net",
	"ibtimes.com",
	"icerocket.com",
	"icij.org",
	"icl-fi.org",
	"iconpaper.org",
	"icu-project.org",
	"idaiwan.com",
	"idemocracy.asia",
	"identi.ca",
	"idiomconnection.com",
	"idlcoyote.com",
	"ifanqiang.com",
	"ifanr.com",
	"ifcss.org",
	"ifjc.org",
	"ifttt.com",
	"ig.com.br",
	"igfw.net",
	"igvita.com",
	"ihakka.net",
	"illusionfactory.com",
	"imageshack.us",
	"imagevenue.com",
	"imagezilla.net",
	"imdb.com",
	"img.ly",
	"in.com",
	"incredibox.fr",
	"inmediahk.net",
	"instagram.com",
	"instapaper.com",
	"internationalrivers.org",
	"internet.org",
	"internetdefenseleague.org",
	"internetfreedom.org",
	"iphone-dev.org",
	"iphone4hongkong.com",
	"iphonehacks.com",
	"ipicture.ru",
	"ipobar.com",
	"ippotv.com",
	"ipvanish.com",
	"iredmail.org",
	"ironicsoftware.com",
	"ironpython.net",
	"isaacmao.com",
	"isgreat.org",
	"islamicity.com",
	"ismprofessional.net",
	"isohunt.com",
	"israbox.com",
	"istef.info",
	"istockphoto.com",
	"isunaffairs.com",
	"isuntv.com",
	"itaboo.info",
	"itshidden.com",
	"itweet.net",
	"iu45.com",
	"iverycd.com",
	"ixquick.com",
	"izaobao.us",
	"izihost.org",
	"izles.net",
	"jackjia.com",
	"japan-whores.com",
	"jayparkinsonmd.com",
	"jbtalks.cc",
	"jbtalks.com",
	"jbtalks.my",
	"jeanyim.com",
	"jgoodies.com",
	"jiaoyou8.com",
	"jiehua.cz",
	"jiepang.com",
	"jieshibaobao.com",
	"jimoparty.com",
	"jinbushe.org",
	"jingpin.org",
	"jinhai.de",
	"jiruan.net",
	"jitouch.com",
	"joachims.org",
	"jobso.tv",
	"joeedelman.com",
	"joeyrobert.org",
	"journalofdemocracy.org",
	"jpopforum.net",
	"jqueryui.com",
	"juliereyc.com",
	"junauza.com",
	"junefourth-20.net",
	"justfreevpn.com",
	"justin.tv",
	"justtristan.com",
	"juziyue.com",
	"jwmusic.org",
	"jyxf.net",
	"ka-wai.com",
	"kaiyuan.de",
	"kakao.com",
	"kangye.org",
	"kanzhongguo.com",
	"kanzhongguo.eu",
	"karayou.com",
	"kcome.org",
	"kcsoftwares.com",
	"kechara.com",
	"keepandshare.com",
	"kendincos.net",
	"kenengba.com",
	"keontech.net",
	"keso.cn",
	"kickstarter.com",
	"killwall.com",
	"kingdomsalvation.org",
	"kinghost.com",
	"kissbbao.cn",
	"kl.am",
	"klip.me",
	"knowledgerush.com",
	"kodingen.com",
	"kompozer.net",
	"koolsolutions.com",
	"koornk.com",
	"kui.name",
	"kun.im",
	"kurtmunger.com",
	"kusocity.com",
	"kwongwah.com.my",
	"kyohk.net",
	"kzeng.info",
	"la-forum.org",
	"labiennale.org",
	"ladbrokes.com",
	"lagranepoca.com",
	"lalulalu.com",
	"laogai.org",
	"laomiu.com",
	"laoyang.info",
	"laptoplockdown.com",
	"laqingdan.net",
	"larsgeorge.com",
	"lastfm.es",
	"latelinenews.com",
	"latimes.com",
	"law.com",
	"lazarsearlymusic.com",
	"leecheukyan.org",
	"lematin.ch",
	"lemonde.fr",
	"lenwhite.com",
	"lerosua.org",
	"lesoir.be",
	"lesscss.org",
	"lester850.info",
	"letscorp.net",
	"liansi.org",
	"lianyue.net",
	"liaowangxizang.net",
	"lidecheng.com",
	"lightbox.com",
	"limiao.net",
	"line.me",
	"linglingfa.com",
	"lingvodics.com",
	"linkideo.com",
	"linksalpha.com",
	"linpie.com",
	"linux-engineer.net",
	"linuxconfig.org",
	"linuxreviews.org",
	"linuxtoy.org",
	"lipuman.com",
	"list.ly",
	"listentoyoutube.com",
	"listorious.com",
	"littlebigdetails.com",
	"liu.lu",
	"liudejun.com",
	"liuhanyu.com",
	"liujianshu.com",
	"liuxiaotong.com",
	"liveleak.com",
	"livestation.com",
	"livestream.com",
	"livingonline.us",
	"livingstream.com",
	"lizhizhuangbi.com",
	"lkcn.net",
	"localpresshk.com",
	"lockdown.com",
	"lockestek.com",
	"logbot.net",
	"logiqx.com",
	"logmike.com",
	"loiclemeur.com",
	"longtermly.net",
	"lookatgame.com",
	"lookingglasstheatre.org",
	"lookpic.com",
	"lovequicksilver.com",
	"lrfz.com",
	"lrip.org",
	"lsforum.net",
	"lsm.org",
	"lsmchinese.org",
	"lsmkorean.org",
	"lsxszzg.com",
	"lupm.org",
	"lushstories.com",
	"lvhai.org",
	"lyricsquote.com",
	"m-team.cc",
	"macrovpn.com",
	"mad-ar.ch",
	"madmenunbuttoned.com",
	"maiio.net",
	"mail-archive.com",
	"maiplus.com",
	"makemymood.com",
	"malaysiakini.com",
	"marc.info",
	"marco.org",
	"marguerite.su",
	"marines.mil",
	"markmail.org",
	"markmilian.com",
	"martau.com",
	"martincartoons.com",
	"maruta.be",
	"marxist.com",
	"marxist.net",
	"marxists.org",
	"mash.to",
	"matainja.com",
	"mathiew-badimon.com",
	"matsushimakaede.com",
	"maxgif.com",
	"mayimayi.com",
	"mcadforums.com",
	"mcfog.com",
	"md-t.org",
	"mediafire.com",
	"meetup.com",
	"mefeedia.com",
	"megarotic.com",
	"megurineluka.com",
	"meirixiaochao.com",
	"melon-peach.com",
	"memedia.cn",
	"memehk.com",
	"memrijttm.org",
	"mesotw.com",
	"metacafe.com",
	"metarthunter.com",
	"meteorshowersonline.com",
	"metrolife.ca",
	"mgoon.com",
	"mgstage.com",
	"mh4u.org",
	"mhradio.org",
	"michaelanti.com",
	"michaelmarketl.com",
	"middle-way.net",
	"mihua.org",
	"mimivip.com",
	"minghui-a.org",
	"minghui-b.org",
	"minghui-school.org",
	"minghui.org",
	"mingjinglishi.com",
	"mingjingnews.com",
	"mingpao.com",
	"mingpaocanada.com",
	"mingpaomonthly.com",
	"mingpaonews.com",
	"mingpaony.com",
	"mingpaosf.com",
	"mingpaotor.com",
	"mingpaovan.com",
	"minimalmac.com",
	"minzhuhua.net",
	"minzhuzhanxian.com",
	"minzhuzhongguo.org",
	"miroguide.com",
	"mirrorbooks.com",
	"mitbbs.com",
	"mixedmedialabs.com",
	"mixero.com",
	"mixpod.com",
	"mizzmona.com",
	"mk5000.com",
	"mlcool.com",
	"mmdays.com",
	"mmmca.com",
	"mobatek.net",
	"mobile01.com",
	"mobileways.de",
	"moby.to",
	"mobypicture.com",
	"modfetish.com",
	"mog.com",
	"molihua.org",
	"mondex.org",
	"mongodb.org",
	"monitorchina.org",
	"monlamit.org",
	"mooo.com",
	"morbell.com",
	"morningsun.org",
	"movabletype.com",
	"moviefap.com",
	"moztw.org",
	"mp",
	"mp3ye.eu",
	"mpettis.com",
	"mpfinance.com",
	"mpinews.com",
	"mrdoob.com",
	"mrtweet.com",
	"msguancha.com",
	"mthruf.com",
	"mtw.tl",
	"multiply.com",
	"multiproxy.org",
	"multiupload.com",
	"muouju.com",
	"muselinks.co.jp",
	"muzi.com",
	"muzi.net",
	"muzu.tv",
	"mx981.com",
	"my-addr.com",
	"my-proxy.com",
	"my903.com",
	"myactimes.com",
	"myaudiocast.com",
	"mychat.to",
	"mychinamyhome.com",
	"mycould.com",
	"myeclipseide.com",
	"myfreshnet.com",
	"mymaji.com",
	"myopenid.com",
	"myparagliding.com",
	"mypopescu.com",
	"mysinablog.com",
	"myspace.com",
	"naacoalition.org",
	"nabble.com",
	"naitik.net",
	"nakido.com",
	"namsisi.com",
	"nanyang.com",
	"nanyangpost.com",
	"nanzao.com",
	"naol.ca",
	"natado.com",
	"navicat.com",
	"navigeaters.com",
	"navy.mil",
	"nbc.com",
	"ncn.org",
	"ncol.com",
	"nde.de",
	"ndoors.com",
	"ndr.de",
	"ned.org",
	"neighborhoodr.com",
	"nekoslovakia.net",
	"neolee.cn",
	"netcolony.com",
	"netfirms.com",
	"netflix.com",
	"netlog.com",
	"netme.cc",
	"networkedblogs.com",
	"neverforget8964.org",
	"new-3lunch.net",
	"new-akiba.com",
	"newcenturymc.com",
	"newcenturynews.com",
	"newchen.com",
	"newgrounds.com",
	"newlandmagazine.com.au",
	"newsancai.com",
	"newscn.org",
	"newsminer.com",
	"newspeak.cc",
	"newstapa.org",
	"newyorktimes.com",
	"nextmedia.com",
	"nexton-net.jp",
	"nf.id.au",
	"nga.mil",
	"ngensis.com",
	"nighost.org",
	"ning.com",
	"nintendium.com",
	"njactb.org",
	"njuice.com",
	"nlfreevpn.com",
	"nobel.se",
	"nobelprize.org",
	"nobodycanstop.us",
	"nodesnoop.com",
	"nokogiri.org",
	"nokola.com",
	"noobbox.com",
	"novelasia.com",
	"nownews.com",
	"nowtorrents.com",
	"noypf.com",
	"npa.go.jp",
	"nps.gov",
	"nrk.no",
	"ntdtv.ca",
	"ntdtv.co",
	"ntdtv.org",
	"ntdtv.ru",
	"nuexpo.com",
	"nurgo-software.com",
	"nuvid.com",
	"nuzcom.com",
	"nvquan.org",
	"nydus.ca",
	"nysingtao.com",
	"nyt.com",
	"nytco.com",
	"nytimes.com",
	"nytimg.com",
	"oauth.net",
	"observechina.net",
	"october-review.org",
	"offbeatchina.com",
	"ogaoga.org",
	"oiktv.com",
	"oizoblog.com",
	"okayfreedom.com",
	"old-cat.net",
	"olumpo.com",
	"olympicwatch.org",
	"omgili.com",
	"omnitalk.com",
	"omy.sg",
	"on.cc",
	"onlylady.cn",
	"onmoon.com",
	"onmoon.net",
	"ontrac.com",
	"oopsforum.com",
	"opendemocracy.net",
	"openid.net",
	"openinkpot.org",
	"openleaks.org",
	"openvpn.net",
	"openwebster.com",
	"opera-mini.net",
	"opera.com",
	"opnir.com",
	"orchidbbs.com",
	"org.uk",
	"orient-doll.com",
	"orientaldaily.com.my",
	"orn.jp",
	"orzdream.com",
	"orzistic.org",
	"osfoora.com",
	"oulove.org",
	"ourdearamy.com",
	"oursogo.com",
	"oursteps.com.au",
	"over-blog.com",
	"overlapr.com",
	"ovi.com",
	"ow.ly",
	"owind.com",
	"owl.li",
	"oxid.it",
	"oyax.com",
	"ozchinese.com",
	"ozyoyo.com",
	"pacificpoker.com",
	"packetix.net",
	"page2rss.com",
	"pagodabox.com",
	"paint.net",
	"palacemoon.com",
	"palm.com",
	"palmislife.com",
	"pandora.com",
	"pandora.tv",
	"panluan.net",
	"panoramio.com",
	"pao-pao.net",
	"paper-replika.com",
	"paper.li",
	"paperb.us",
	"parade.com",
	"parislemon.com",
	"parkansky.com",
	"pastebin.com",
	"pastie.org",
	"path.com",
	"pathtosharepoint.com",
	"pbs.org",
	"pbwiki.com",
	"pbworks.com",
	"pbxes.com",
	"pbxes.org",
	"pcdiscuss.com",
	"pcij.org",
	"pdetails.com",
	"pdproxy.com",
	"peacefire.org",
	"peacehall.com",
	"peeasian.com",
	"peerpong.com",
	"pekingduck.org",
	"penchinese.com",
	"penchinese.net",
	"pengyulong.com",
	"pentalogic.net",
	"penthouse.com",
	"peopo.org",
	"percy.in",
	"perfectvpn.net",
	"perfspot.com",
	"perlhowto.com",
	"philly.com",
	"phonegap.com",
	"photofocus.com",
	"phuquocservices.com",
	"picidae.net",
	"picturesocial.com",
	"pidown.com",
	"pign.net",
	"pikchur.com",
	"pilotmoon.com",
	"pin6.com",
	"ping.fm",
	"pinoy-n.com",
	"piring.com",
	"pixelqi.com",
	"pixnet.in",
	"pixnet.net",
	"pk.com",
	"placemix.com",
	"plixi.com",
	"plunder.com",
	"plus28.com",
	"plusbb.com",
	"pmates.com",
	"po2b.com",
	"podictionary.com",
	"pokerstars.com",
	"pokerstrategy.com",
	"politicalchina.org",
	"popyard.com",
	"popyard.org",
	"phncdn.com",
	"pose.com",
	"post.ly",
	"post852.com",
	"postadult.com",
	"posterous.com",
	"power.com",
	"powerapple.com",
	"powercx.com",
	"powerpointninja.com",
	"prayforchina.net",
	"premeforwindows7.com",
	"presentationzen.com",
	"prestige-av.com",
	"printfriendly.com",
	"privacybox.de",
	"privateinternetaccess.com",
	"privatepaste.com",
	"privatetunnel.com",
	"procopytips.com",
	"prosiben.de",
	"proxifier.com",
	"proxlet.com",
	"proxomitron.info",
	"proxy.org",
	"proxypy.net",
	"proxyroad.com",
	"prozz.net",
	"psblog.name",
	"psiphon.ca",
	"ptt.cc",
	"puffinbrowser.com",
	"puffstore.com",
	"pullfolio.com",
	"pure18.com",
	"pureconcepts.net",
	"purepdf.com",
	"purevpn.com",
	"putlocker.com",
	"pwned.com",
	"python.com",
	"qanote.com",
	"qidian.ca",
	"qienkuen.org",
	"qiwen.lu",
	"qixianglu.cn",
	"qkshare.com",
	"qmzdd.com",
	"qoos.com",
	"qstatus.com",
	"qtrac.eu",
	"qtweeter.com",
	"quadedge.com",
	"qusi8.net",
	"qvodzy.org",
	"qx.net",
	"qxbbs.org",
	"radicalparty.org",
	"radioaustralia.net.au",
	"radiotime.com",
	"radiovaticana.org",
	"radiovncr.com",
	"rangzen.org",
	"ranxiang.com",
	"ranyunfei.com",
	"rapbull.net",
	"rapidgator.net",
	"rapidshare8.com",
	"rapidsharedata.com",
	"rcinet.ca",
	"rdio.com",
	"read100.com",
	"readmoo.com",
	"realraptalk.com",
	"recaptcha.net",
	"recordhistory.org",
	"redchinacn.org",
	"referer.us",
	"reflectivecode.com",
	"relaxbbs.com",
	"renminbao.com",
	"renyurenquan.org",
	"rerouted.org",
	"retweeteffect.com",
	"retweetist.com",
	"retweetrank.com",
	"reuters.com",
	"revleft.com",
	"revver.com",
	"rfa.org",
	"rfachina.com",
	"rfamobile.org",
	"rferl.org",
	"rfi.fr",
	"rfi.my",
	"rhcloud.com",
	"rightster.com",
	"riku.me",
	"rileyguide.com",
	"rlwlw.com",
	"rmjdw.com",
	"rnw.nl",
	"robtex.com",
	"robustnessiskey.com",
	"rockmelt.com",
	"rocmp.org",
	"rojo.com",
	"romanandreg.com",
	"ronjoneswriter.com",
	"roodo.com",
	"rotten.com",
	"rsf-chinese.org",
	"rsf.org",
	"rssmeme.com",
	"ruanyifeng.com",
	"rushbee.com",
	"rutube.ru",
	"ruyiseek.com",
	"s1heng.com",
	"s8forum.com",
	"sadpanda.us",
	"saiq.me",
	"samair.ru",
	"sammyjs.org",
	"samsoff.es",
	"sandnoble.com",
	"sankaizok.com",
	"sapikachu.net",
	"savemedia.com",
	"savetibet.de",
	"savetibet.fr",
	"savetibet.nl",
	"savetibet.org",
	"savetibet.ru",
	"savevid.com",
	"say2.info",
	"sciencemag.org",
	"scmp.com",
	"scmpchinese.com",
	"scribd.com",
	"scriptspot.com",
	"seapuff.com",
	"search.com",
	"secretchina.com",
	"secretgarden.no",
	"secureserver.net",
	"securitykiss.com",
	"seesmic.com",
	"seevpn.com",
	"seezone.net",
	"sejie.com",
	"sendoid.com",
	"sendspace.com",
	"seraph.me",
	"sesawe.net",
	"sesawe.org",
	"sethwklein.net",
	"sevenload.com",
	"sfileydy.com",
	"sftuk.org",
	"shadow.ma",
	"shadowsocks.org",
	"shahamat-english.com",
	"shangfang.org",
	"shapeservices.com",
	"sharebee.com",
	"sharecool.org",
	"sharkdolphin.com",
	"shaunthesheep.com",
	"sheikyermami.com",
	"shellmix.com",
	"shenshou.org",
	"shenyunperformingarts.org",
	"shenzhoufilm.com",
	"shinychan.com",
	"shitaotv.org",
	"shixiao.org",
	"shizhao.org",
	"shkspr.mobi",
	"shodanhq.com",
	"shopping.com",
	"showtime.jp",
	"shvoong.com",
	"shwchurch3.com",
	"sidelinesnews.com",
	"sidelinessportseatery.com",
	"simplecd.org",
	"simpleproductivityblog.com",
	"sina.com",
	"singtao.ca",
	"singtao.com",
	"sino-monthly.com",
	"sinoants.com",
	"sinocast.com",
	"sinocism.com",
	"sinomontreal.ca",
	"sinonet.ca",
	"sinopitt.info",
	"sinoquebec.com",
	"sis001.com",
	"sis001.us",
	"site90.net",
	"sitemaps.org",
	"sitetag.us",
	"sjum.cn",
	"skybet.com",
	"skyhighpremium.com",
	"skykiwi.com",
	"skype.com",
	"skyvegas.com",
	"slacker.com",
	"slandr.net",
	"slavasoft.com",
	"slheng.com",
	"slickvpn.com",
	"slideshare.net",
	"slinkset.com",
	"slutload.com",
	"smhric.org",
	"snapchat.com",
	"snaptu.com",
	"sndcdn.com",
	"sneakme.net",
	"so-ga.net",
	"so-news.com",
	"sobees.com",
	"soc.mil",
	"socialwhale.com",
	"sockslist.net",
	"sod.co.jp",
	"softether-download.com",
	"softether.co.jp",
	"softether.org",
	"softwarebychuck.com",
	"sogclub.com",
	"sogoo.org",
	"sogrady.me",
	"sohcradio.com",
	"sohfrance.org",
	"soifind.com",
	"sokamonline.com",
	"solozorro.tk",
	"somee.com",
	"songjianjun.com",
	"sonidodelaesperanza.org",
	"sopcast.com",
	"sopcast.org",
	"sorting-algorithms.com",
	"soumo.info",
	"soundcloud.com",
	"soundofhope.kr",
	"soundofhope.org",
	"soup.io",
	"soupofmedia.com",
	"sourceforge.net",
	"sowiki.net",
	"space-scape.com",
	"spankwire.com",
	"sparrowmailapp.com",
	"spb.com",
	"speckleapp.com",
	"spencertipping.com",
	"spinejs.com",
	"spotify.com",
	"sproutcore.com",
	"squarespace.com",
	"ssh91.com",
	"stackfile.com",
	"standupfortibet.org",
	"starp2p.com",
	"startpage.com",
	"state.gov",
	"state168.com",
	"staticflickr.com",
	"steel-storm.com",
	"sthoo.com",
	"stickam.com",
	"stickeraction.com",
	"stonegames.net",
	"stoneip.info",
	"stoptibetcrisis.net",
	"storagenewsletter.com",
	"storify.com",
	"streamingthe.net",
	"streetvoice.com",
	"strongvpn.com",
	"studentsforafreetibet.org",
	"stuffimreading.com",
	"stuffimreading.net",
	"sugarsync.com",
	"summify.com",
	"sun1911.com",
	"suoluo.org",
	"supertweet.net",
	"surfeasy.com.au",
	"svwind.com",
	"sweux.com",
	"swift-tools.net",
	"sydneytoday.com",
	"sylfoundation.org",
	"syncback.com",
	"sysadmin1138.net",
	"sysresccd.org",
	"sytes.net",
	"syx86.cn",
	"syx86.com",
	"szbbs.net",
	"t.co",
	"t35.com",
	"t66y.com",
	"taa-usa.org",
	"tabtter.jp",
	"tacem.org",
	"tafaward.com",
	"tagwalk.com",
	"taipeisociety.org",
	"taiwandaily.net",
	"taiwankiss.com",
	"taiwannation.com",
	"taiwantp.net",
	"taiwanus.net",
	"taiwanyes.com",
	"tamiaode.tk",
	"tampabay.com",
	"tanc.org",
	"tangben.com",
	"taolun.info",
	"tap11.com",
	"taragana.com",
	"target.com",
	"taweet.com",
	"tbpic.info",
	"tbsec.org",
	"tbsn.org",
	"tbsseattle.org",
	"tchrd.org",
	"tcno.net",
	"teamseesmic.com",
	"teashark.com",
	"techlifeweb.com",
	"techparaiso.com",
	"teck.in",
	"telecomspace.com",
	"tenacy.com",
	"theampfactory.com",
	"theappleblog.com",
	"theatrum-belli.com",
	"thebcomplex.com",
	"theblemish.com",
	"thebodyshop-usa.com",
	"thechinabeat.org",
	"thechinastory.org",
	"thedailywh.at",
	"thedieline.com",
	"thedw.us",
	"thegatesnotes.com",
	"thegioitinhoc.vn",
	"theguardian.co",
	"thehots.info",
	"thehousenews.com",
	"thehun.net",
	"thehungrydudes.com",
	"theinternetwishlist.com",
	"thelifeyoucansave.com",
	"thelius.org",
	"thepiratebay.org",
	"thepiratebay.se",
	"theqii.info",
	"thereallove.kr",
	"thesartorialist.com",
	"thespeeder.com",
	"thestandnews.com",
	"thetibetpost.com",
	"thetrotskymovie.com",
	"thevivekspot.com",
	"thewgo.org",
	"thinkingtaiwan.com",
	"thisav.com",
	"thisiswhyyouarefat.com",
	"thkphoto.com",
	"thomasbernhard.org",
	"threatchaos.com",
	"throughnightsfire.com",
	"thumbzilla.com",
	"thywords.com",
	"tiananmenmother.org",
	"tiananmenuniv.com",
	"tiananmenuniv.net",
	"tiandixing.org",
	"tianhuayuan.com",
	"tiantibooks.org",
	"tianzhu.org",
	"tibet.at",
	"tibet.com",
	"tibet.net",
	"tibetalk.com",
	"tibetanyouthcongress.org",
	"tibetcorps.org",
	"tibetfund.org",
	"tibetjustice.org",
	"tibetoffice.org",
	"tibetonline.com",
	"tibetonline.tv",
	"tibetsun.com",
	"tibetwrites.org",
	"tidyread.com",
	"tiffanyarment.com",
	"time.com",
	"tiney.com",
	"tinychat.com",
	"tinypaste.com",
	"tistory.com",
	"tjholowaychuk.com",
	"tkcs-collins.com",
	"tkforum.tk",
	"tl.gd",
	"tmagazine.com",
	"tmi.me",
	"tnaflix.com",
	"togetter.com",
	"tokyo-247.com",
	"tokyo-hot.com",
	"tokyocn.com",
	"tomayko.com",
	"tomsc.com",
	"tono-oka.jp",
	"tonyyan.net",
	"toodoc.com",
	"toonel.net",
	"topify.com",
	"topnews.in",
	"topshare.us",
	"topshareware.com",
	"topstyle4.com",
	"topsy.com",
	"tora.to",
	"torproject.org",
	"torrentcrazy.com",
	"torrentproject.se",
	"torvpn.com",
	"touch99.com",
	"toutfr.com",
	"transgressionism.org",
	"transparency.org",
	"travelinlocal.com",
	"trendsmap.com",
	"trialofccp.org",
	"tripod.com",
	"trouw.nl",
	"trulyergonomic.com",
	"trustedbi.com",
	"truthcn.com",
	"truveo.com",
	"tsctv.net",
	"tsemtulku.com",
	"tsquare.tv",
	"tsunagarumon.com",
	"tsuru-bird.net",
	"tt1069.com",
	"tttan.com",
	"tuanzt.com",
	"tuidang.net",
	"tuidang.org",
	"tuitui.info",
	"tumblweed.org",
	"tumutanzi.com",
	"tunein.com",
	"tunnelbear.com",
	"turbobit.net",
	"turbotwitter.com",
	"turningtorso.com",
	"turntable.fm",
	"tuxtraining.com",
	"tv-intros.com",
	"tv.com",
	"tvants.com",
	"tvb.com",
	"tvboxnow.com",
	"tvider.com",
	"tvunetworks.com",
	"tw",
	"tw-npo.org",
	"twapperkeeper.com",
	"twaud.io",
	"twbbs.org",
	"twblogger.com",
	"tweepguide.com",
	"tweeplike.me",
	"tweepmag.com",
	"tweepml.org",
	"tweetbackup.com",
	"tweetboard.com",
	"tweetboner.biz",
	"tweetdeck.com",
	"tweetedtimes.com",
	"tweetmylast.fm",
	"tweetphoto.com",
	"tweetrans.com",
	"tweetree.com",
	"tweetwally.com",
	"tweetymail.com",
	"twftp.org",
	"twhirl.org",
	"twibase.com",
	"twibble.de",
	"twibbon.com",
	"twibs.com",
	"twicsy.com",
	"twifan.com",
	"twiffo.com",
	"twiggit.org",
	"twilio.com",
	"twilog.org",
	"twimbow.com",
	"twimg.com",
	"twip.me",
	"twipple.jp",
	"twistar.cc",
	"twisternow.com",
	"twistory.net",
	"twit2d.com",
	"twitbrowser.net",
	"twitcause.com",
	"twitgether.com",
	"twitgoo.com",
	"twitiq.com",
	"twitlonger.com",
	"twitoaster.com",
	"twitonmsn.com",
	"twitpic.com",
	"twitreferral.com",
	"twitstat.com",
	"twittbot.net",
	"twitter.com",
	"twitter.jp",
	"twitter4j.org",
	"twittercounter.com",
	"twitterfeed.com",
	"twittergadget.com",
	"twitterkr.com",
	"twittermail.com",
	"twittertim.es",
	"twitthat.com",
	"twitturk.com",
	"twitturly.com",
	"twitvid.com",
	"twitzap.com",
	"twiyia.com",
	"twreg.info",
	"twstar.net",
	"twt.fm",
	"twt.tl",
	"twtkr.com",
	"twtrland.com",
	"twttr.com",
	"twurl.nl",
	"twyac.org",
	"tycool.com",
	"tynsoe.org",
	"typepad.com",
	"tzangms.com",
	"ub0.cc",
	"uberproxy.net",
	"ucam.org",
	"ucdc1998.org",
	"uderzo.it",
	"udn.com",
	"ufreevpn.com",
	"ugo.com",
	"uhrp.org",
	"uighurbiz.net",
	"uk.to",
	"ulike.net",
	"ultravpn.fr",
	"ultraxs.com",
	"unblock.cn.com",
	"unblocksit.es",
	"uncyclomedia.org",
	"uncyclopedia.info",
	"unholyknight.com",
	"uni.cc",
	"unicode.org",
	"uniteddaily.com.my",
	"unix100.com",
	"unknownspace.org",
	"unpo.org",
	"uocn.org",
	"updatestar.com",
	"upholdjustice.org",
	"upload4u.info",
	"uploaded.net",
	"uploaded.to",
	"uploadstation.com",
	"urbanoutfitters.com",
	"urlborg.com",
	"urlparser.com",
	"us.to",
	"usa.gov",
	"usacn.com",
	"usejump.com",
	"usfk.mil",
	"usgs.gov",
	"usmc.mil",
	"ustream.tv",
	"ustwrap.info",
	"usus.cc",
	"utom.us",
	"uushare.com",
	"uwants.com",
	"uwants.net",
	"uyghuramerican.org",
	"uyghurcongress.org",
	"uygur.org",
	"v-state.org",
	"v70.us",
	"v7976888.info",
	"vaayoo.com",
	"value-domain.com",
	"van698.com",
	"vanemu.cn",
	"vanilla-jp.com",
	"vansky.com",
	"vapurl.com",
	"vatn.org",
	"vcf-online.org",
	"vcfbuilder.org",
	"veempiire.com",
	"vegorpedersen.com",
	"velkaepocha.sk",
	"venbbs.com",
	"venchina.com",
	"ventureswell.com",
	"veoh.com",
	"verizon.net",
	"verybs.com",
	"vevo.com",
	"views.fm",
	"viki.com",
	"vimeo.com",
	"vimgolf.com",
	"vimperator.org",
	"vincnd.com",
	"vinniev.com",
	"visiontimes.com",
	"vllcs.org",
	"vmixcore.com",
	"voa.mobi",
	"voacantonese.com",
	"voachinese.com",
	"voachineseblog.com",
	"voagd.com",
	"voanews.com",
	"voatibetan.com",
	"vocn.tv",
	"vot.org",
	"voy.com",
	"vpnbook.com",
	"vpncup.com",
	"vpnfire.com",
	"vpngate.jp",
	"vpngate.net",
	"vpnpop.com",
	"vpnpronet.com",
	"vtunnel.com",
	"w.org",
	"w3.org",
	"waffle1999.com",
	"wahas.com",
	"waigaobu.com",
	"waikeung.org",
	"waiwaier.com",
	"wallornot.org",
	"wallpapercasa.com",
	"wan-press.org",
	"wanderinghorse.net",
	"wangafu.net",
	"wangjinbo.org",
	"wanglixiong.com",
	"wangruoshui.net",
	"wangruowang.org",
	"want-daily.com",
	"wapedia.mobi",
	"waqn.com",
	"warehouse333.com",
	"waselpro.com",
	"washeng.net",
	"watchmygf.net",
	"wattpad.com",
	"wdf5.com",
	"wearn.com",
	"web2project.net",
	"webbang.net",
	"webfee.tk",
	"weblagu.com",
	"webmproject.org",
	"webs-tv.net",
	"webshots.com",
	"websitepulse.com",
	"webworkerdaily.com",
	"weeewooo.net",
	"weekmag.info",
	"wefong.com",
	"weiboleak.com",
	"weigegebyc.dreamhosters.com",
	"weijingsheng.org",
	"weiming.info",
	"weiquanwang.org",
	"weisuo.ws",
	"wellplacedpixels.com",
	"wengewang.com",
	"wengewang.org",
	"wenhui.ch",
	"wenku.com",
	"wenweipo.com",
	"wenxuecity.com",
	"wenyunchao.com",
	"wepn.info",
	"westca.com",
	"westernwolves.com",
	"westkit.net",
	"wet123.com",
	"wetplace.com",
	"wetpussygames.com",
	"wexiaobo.org",
	"wezhiyong.org",
	"wezone.net",
	"wforum.com",
	"whatblocked.com",
	"whereiswerner.com",
	"whippedass.com",
	"who.is",
	"whydidyoubuymethat.com",
	"whylover.com",
	"whyx.org",
	"wikia.com",
	"wikibooks.org",
	"wikileaks.ch",
	"wikileaks.de",
	"wikileaks.eu",
	"wikileaks.lu",
	"wikileaks.org",
	"wikileaks.pl",
	"wikilivres.info",
	"wikimapia.org",
	"wikimedia.org",
	"wikimedia.org.mo",
	"wikinews.org",
	"wikipedia.org",
	"wikisource.org",
	"wikiwiki.jp",
	"williamhill.com",
	"willw.net",
	"windowsphoneme.com",
	"winwhispers.info",
	"wiredbytes.com",
	"wiredpen.com",
	"wireshark.org",
	"wisevid.com",
	"witnessleeteaching.com",
	"witopia.net",
	"wo.tc",
	"woeser.com",
	"woesermiddle-way.net",
	"wolfax.com",
	"womensrightsofchina.org",
	"woopie.jp",
	"woopie.tv",
	"wordboner.com",
	"wordpress.com",
	"wordsandturds.com",
	"workatruna.com",
	"worldcat.org",
	"worldjournal.com",
	"worstthingieverate.com",
	"wow-life.net",
	"wowlegacy.ml",
	"woxinghuiguo.com",
	"wozy.in",
	"wp.com",
	"wpoforum.com",
	"wqlhw.com",
	"wqyd.org",
	"wrchina.org",
	"wretch.cc",
	"wsj.com",
	"wsj.net",
	"wtfpeople.com",
	"wuala.com",
	"wuerkaixi.com",
	"wuguoguang.com",
	"wujie.net",
	"wujieliulan.com",
	"wukangrui.net",
	"wwitv.com",
	"x-art.com",
	"x-berry.com",
	"x-wall.org",
	"x1949x.com",
	"x365x.com",
	"xanga.com",
	"xbabe.com",
	"xbookcn.com",
	"xcafe.in",
	"xcity.jp",
	"xcritic.com",
	"xfiles.to",
	"xfm.pp.ru",
	"xgmyd.com",
	"xh4n.cn",
	"xhamster.com",
	"xiaochuncnjp.com",
	"xiaod.in",
	"xiaohexie.com",
	"xiaoma.org",
	"xiezhua.com",
	"xing.com",
	"xinhuanet.org",
	"xinsheng.net",
	"xinshijue.com",
	"xinyubbs.net",
	"xizang-zhiye.org",
	"xjp.cc",
	"xlgames.com",
	"xml-training-guide.com",
	"xmovies.com",
	"xmusic.fm",
	"xpdo.net",
	"xpud.org",
	"xrea.com",
	"xskywalker.com",
	"xthost.info",
	"xuchao.net",
	"xuchao.org",
	"xuite.net",
	"xuzhiyong.net",
	"xuzhuoer.com",
	"yahoo.co.jp",
	"yahoo.com",
	"yam.com",
	"yasukuni.or.jp",
	"ydy.com",
	"yeelou.com",
	"yeeyi.com",
	"yegle.net",
	"yfrog.com",
	"yhcw.net",
	"yi.org",
	"yidio.com",
	"yilubbs.com",
	"yimg.com",
	"yipub.com",
	"yogichen.org",
	"yong.hu",
	"yorkbbs.ca",
	"youjizz.com",
	"youmaker.com",
	"youpai.org",
	"your-freedom.net",
	"yourepeat.com",
	"yousendit.com",
	"youthbao.com",
	"youthnetradio.org",
	"youtu.be",
	"youtube-nocookie.com",
	"youtube.com",
	"youversion.com",
	"ytht.net",
	"ytimg.com",
	"yuanming.net",
	"yunchao.net",
	"yvesgeleyn.com",
	"yx51.net",
	"yyii.org",
	"yymaya.com",
	"yzzk.com",
	"zacebook.com",
	"zannel.com",
	"zaobao.com",
	"zaobao.com.sg",
	"zaozon.com",
	"zarias.com",
	"zattoo.com",
	"zengjinyan.org",
	"zeutch.com",
	"zfreet.com",
	"zgzcjj.net",
	"zhanbin.net",
	"zhe.la",
	"zhenghui.org",
	"zhenlibu.info",
	"zhinengluyou.com",
	"zhong.pp.ru",
	"zhongguotese.net",
	"zhongmeng.org",
	"zhreader.com",
	"zhuichaguoji.org",
	"ziddu.com",
	"zillionk.com",
	"zinio.com",
	"ziplib.com",
	"zkaip.com",
	"zlib.net",
	"zmw.cn",
	"zoho.com",
	"zomobo.net",
	"filosoft.ee",
	"zonaeuropa.com",
	"zonble.net",
	"zootool.com",
	"zoozle.net",
	"zozotown.com",
	"zshare.net",
	"zsrhao.com",
	"zuo.la",
	"zuola.com",
}

var pacListGN = []string{
	"baidu.com",
	"bdstatic.com",
	"bilibili.com",
	"bilivideo.com",
	"qq.com",
	"bootcdn.net",
	"baidustatic.com",
}

var ip2Region *ip2region.Ip2Region

func init() {
	ip2, err := ip2region.New("ip2region.db")
	if err != nil {
		log.Println("ip2region not fund")
		return
	}

	ip2Region = ip2
}

func IsPac(domain string) bool {
	for _, v := range defaultPacDomain {
		if strings.Contains(domain, v) {
			return true
		}
	}

	for _, v := range pacListGN {
		if strings.Contains(domain, v) {
			return false
		}
	}

	for _, v := range conf.AgentConfig.ProxyList {
		if strings.Contains(domain, v) {
			return true
		}
	}

	for _, v := range conf.AgentConfig.NoProxyList {
		if strings.Contains(domain, v) {
			return false
		}
	}

	if ip2Region == nil {
		return false
	}

	dns, err := utils.SearchDns(domain, conf.AgentConfig.GetDNS())
	if err != nil {
		host, err := net.LookupHost(domain)
		if err != nil {
			log.Println(err)
		}
		if len(host) != 0 {
			dns = append(dns, host...)
		}
	}

	if len(dns) == 0 {
		return true
	}

	search, err := ip2Region.BinarySearch(dns[0])
	if err != nil {
		return false
	}
	if (search.Country == "中国" && search.Province != "台湾" && search.Province != "香港" && search.Province != " 澳门") || (search.Country == "0" && search.City == "内网IP") {
		return false
	}

	return false
}
