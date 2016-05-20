package regionid

var (
	builtinWorld = []byte(
		`
country,004,AF,AFG,004,阿富汗
country,248,AX,ALA,248,奥兰群岛
country,008,AL,ALB,008,阿尔巴尼亚
country,012,DZ,DZA,012,阿尔及利亚
country,016,AS,ASM,016,美属萨摩亚
country,020,AD,AND,020,安道尔
country,024,AO,AGO,024,安哥拉
country,660,AI,AIA,660,安圭拉
country,010,AQ,ATA,010,南极洲
country,028,AG,ATG,028,安提瓜和巴布达
country,032,AR,ARG,032,阿根廷
country,051,AM,ARM,051,亚美尼亚
country,533,AW,ABW,533,阿鲁巴
country,036,AU,AUS,036,澳大利亚
country,040,AT,AUT,040,奥地利
country,031,AZ,AZE,031,阿塞拜疆
country,044,BS,BHS,044,巴哈马
country,048,BH,BHR,048,巴林
country,050,BD,BGD,050,孟加拉
country,052,BB,BRB,052,巴巴多斯
country,112,BY,BLR,112,白俄罗斯
country,056,BE,BEL,056,比利时
country,084,BZ,BLZ,084,伯利兹
country,204,BJ,BEN,204,贝宁
country,060,BM,BMU,060,百慕大
country,064,BT,BTN,064,不丹
country,068,BO,BOL,068,玻利维亚
country,535,BQ,BES,535,荷兰加勒比区
country,070,BA,BIH,070,波斯尼亚和黑塞哥维那
country,072,BW,BWA,072,博茨瓦纳
country,074,BV,BVT,074,布韦岛
country,076,BR,BRA,076,巴西
country,086,IO,IOT,086,英属印度洋领地
country,096,BN,BRN,096,文莱
country,100,BG,BGR,100,保加利亚
country,854,BF,BFA,854,布基纳法索
country,108,BI,BDI,108,布隆迪
country,132,CV,CPV,132,佛得角
country,116,KH,KHM,116,柬埔寨
country,120,CM,CMR,120,喀麦隆
country,124,CA,CAN,124,加拿大
country,136,KY,CYM,136,开曼群岛
country,140,CF,CAF,140,中非
country,148,TD,TCD,148,乍得
country,152,CL,CHL,152,智利
country,156,CN,CHN,156,中国

subdivision,110000,CN-11,北京
subdivision,120000,CN-12,天津
subdivision,130000,CN-13,河北
city,130100,石家庄
city,130200,唐山
city,130300,秦皇岛
city,130400,邯郸
city,130500,邢台
city,130600,保定
city,130700,张家口
city,130800,承德
city,130900,沧州
city,131000,廊坊
city,131100,衡水
city,131082,三河
city,130982,任丘
city,130110,鹿泉
subdivision,140000,CN-14,山西
city,140100,太原
city,140200,大同
city,140300,阳泉
city,140400,长治
city,140500,晋城
city,140600,朔州
city,140700,晋中
city,140781,介休
city,140800,运城
city,140900,忻州
city,141000,临汾
city,141100,吕梁
city,141182,汾阳
subdivision,150000,CN-15,内蒙古
city,150100,呼和浩特
city,150200,包头
city,150300,乌海
city,150400,赤峰
city,150500,通辽
city,150600,鄂尔多斯
city,150700,呼伦贝尔
city,150800,巴彦淖尔
city,150900,乌兰察布
city,152200,兴安盟
city,152500,锡林郭勒盟
city,152900,阿拉善盟
subdivision,210000,CN-21,辽宁
city,210100,沈阳
city,210200,大连
city,210300,鞍山
city,210400,抚顺
city,210500,本溪
city,210600,丹东
city,210700,锦州
city,210800,营口
city,210900,阜新
city,211000,辽阳
city,211100,盘锦
city,211200,铁岭
city,211300,朝阳
city,211400,葫芦岛
subdivision,220000,CN-22,吉林
city,220100,长春
city,220200,吉林市
city,220300,四平
city,220400,辽源
city,220500,通化
city,220600,白山
city,220700,松原
city,220800,白城
city,222400,延边朝鲜族自治州
city,222401,延吉
subdivision,230000,CN-23,黑龙江
city,230100,哈尔滨
city,230182,双城
city,230200,齐齐哈尔
city,230300,鸡西
city,230400,鹤岗
city,230500,双鸭山
city,230600,大庆
city,230700,伊春
city,230800,佳木斯
city,230900,七台河
city,231000,牡丹江
city,231100,黑河
city,231200,绥化
city,232700,大兴安岭地区
subdivision,310000,CN-31,上海
subdivision,320000,CN-32,江苏
city,320100,南京
city,320200,无锡
city,320300,徐州
city,320400,常州
city,320500,苏州
city,320581,常熟
city,320582,张家港
city,320583,昆山
city,320585,太仓
city,320600,南通
city,320700,连云港
city,320800,淮安
city,320826,涟水
city,320900,盐城
city,321000,扬州
city,321100,镇江
city,321181,丹阳
city,321183,句容
city,321200,泰州
city,321300,宿迁
subdivision,330000,CN-33,浙江
city,330100,杭州
city,330109,萧山
city,330183,富阳
city,330185,临安
city,330200,宁波
city,330281,余姚
city,330300,温州
city,330381,瑞安
city,330382,乐清
city,330400,嘉兴
city,330481,海宁
city,330500,湖州
city,330600,绍兴
city,330700,金华
city,330782,义乌
city,330783,东阳
city,330800,衢州
city,330900,舟山
city,331000,台州
city,331100,丽水
subdivision,340000,CN-34,安徽
city,340100,合肥
city,340181,巢湖
city,340200,芜湖
city,340300,蚌埠
city,340400,淮南
city,340500,马鞍山
city,340600,淮北
city,340700,铜陵
city,340800,安庆
city,341000,黄山
city,341100,滁州
city,341200,阜阳
city,341300,宿州
city,341500,六安
city,341600,亳州
city,341700,池州
city,341800,宣城
subdivision,350000,CN-35,福建
city,350100,福州
city,350200,厦门
city,350300,莆田
city,350400,三明
city,350481,永安
city,350500,泉州
city,350600,漳州
city,350700,南平
city,350800,龙岩
city,350900,宁德
subdivision,360000,CN-36,江西
city,360100,南昌
city,360200,景德镇
city,360300,萍乡
city,360400,九江
city,360500,新余
city,360600,鹰潭
city,360700,赣州
city,360800,吉安
city,360900,宜春
city,361000,抚州
city,361100,上饶
subdivision,370000,CN-37,山东
city,370100,济南
city,370181,章丘
city,370200,青岛
city,370281,胶州
city,370300,淄博
city,370400,枣庄
city,370500,东营
city,370600,烟台
city,370682,莱阳
city,370700,潍坊
city,370781,青州
city,370782,诸城
city,370800,济宁
city,370881,曲阜
city,370900,泰安
city,370902,泰山
city,371000,威海
city,371100,日照
city,371200,莱芜
city,371300,临沂
city,371400,德州
city,371500,聊城
city,371600,滨州
city,371700,菏泽
subdivision,410000,CN-41,河南
city,410100,郑州
city,410185,登封
city,410200,开封
city,410300,洛阳
city,410400,平顶山
city,410500,安阳
city,410600,鹤壁
city,410700,新乡
city,410800,焦作
city,410900,濮阳
city,411000,许昌
city,411100,漯河
city,411200,三门峡
city,411300,南阳
city,411400,商丘
city,411481,永城
city,411500,信阳
city,411600,周口
city,411700,驻马店
city,419001,济源
subdivision,420000,CN-42,湖北
city,420100,武汉
city,420200,黄石
city,420300,十堰
city,420500,宜昌
city,420600,襄阳
city,420700,鄂州
city,420800,荆门
city,420900,孝感
city,421000,荆州
city,421100,黄冈
city,421200,咸宁
city,421300,随州
city,422800,恩施
city,422800,恩施土家族苗族自治州
city,429004,仙桃
city,429005,潜江
city,429006,天门
city,429021,神农架
city,429021,神农架林区
subdivision,430000,CN-43,湖南
city,430100,长沙
city,430200,株洲
city,430300,湘潭
city,430400,衡阳
city,430481,耒阳
city,430500,邵阳
city,430600,岳阳
city,430700,常德
city,430800,张家界
city,430900,益阳
city,431000,郴州
city,431100,永州
city,431200,怀化
city,431300,娄底
city,433100,湘西土家族苗族自治州
city,433101,吉首
subdivision,440000,CN-44,广东
city,440100,广州
city,440200,韶关
city,440300,深圳
city,440400,珠海
city,440500,汕头
city,440600,佛山
city,440605,南海
city,440506,顺德
city,440700,江门
city,440800,湛江
city,440900,茂名
city,440981,高州
city,441200,肇庆
city,441300,惠州
city,441303,惠阳
city,441400,梅州
city,441500,汕尾
city,441600,河源
city,441700,阳江
city,441800,清远
city,441900,东莞
city,442000,中山
city,445100,潮州
city,445200,揭阳
city,445281,普宁
city,445300,云浮
subdivision,450000,CN-45,广西
city,450100,南宁
city,450200,柳州
city,450300,桂林
city,450400,梧州
city,450500,北海
city,450600,防城港
city,450700,钦州
city,450800,贵港
city,450900,玉林
city,451000,百色
city,451100,贺州
city,451200,河池
city,451300,来宾
city,451400,崇左
subdivision,460000,CN-46,海南
city,460100,海口
city,460200,三亚
city,460300,三沙
city,469001,五指山
city,469002,琼海
city,469003,儋州
city,469005,文昌
city,469006,万宁
city,469007,东方
city,469021,定安县
city,469022,屯昌县
city,469023,澄迈县
city,469024,临高县
city,469025,白沙黎族自治县
city,469026,昌江黎族自治县
city,469027,乐东黎族自治县
city,469028,陵水黎族自治县
city,469029,保亭黎族苗族自治县
city,469030,琼中黎族苗族自治县
subdivision,500000,CN-50,重庆
subdivision,510000,CN-51,四川
city,510100,成都
city,510181,都江堰
city,510300,自贡
city,510400,攀枝花
city,510500,泸州
city,510600,德阳
city,510681,广汉
city,510700,绵阳
city,510781,江油
city,510800,广元
city,510900,遂宁
city,511000,内江
city,511100,乐山
city,511300,南充
city,511400,眉山
city,511500,宜宾
city,511600,广安
city,511700,达州
city,511800,雅安
city,511900,巴中
city,512000,资阳
city,513200,阿坝藏族羌族自治州
city,513300,甘孜藏族自治州
city,513322,康定
city,513400,凉山彝族自治州
city,513401,西昌
subdivision,520000,CN-52,贵州
city,520100,贵阳
city,520200,六盘水
city,520300,遵义
city,520400,安顺
city,520500,毕节
city,520600,铜仁
city,522300,黔西南布依族苗族自治州
city,522301,兴义
city,522600,黔东南苗族侗族自治州
city,522601,凯里
city,522700,黔南布依族苗族自治州
city,522701,都匀
subdivision,530000,CN-53,云南
city,530100,昆明
city,530300,曲靖
city,530400,玉溪
city,530500,保山
city,530600,昭通
city,530700,丽江
city,530800,普洱
city,530802,思茅
city,530900,临沧
city,532300,楚雄彝族自治州
city,532500,红河哈尼族彝族自治州
city,532503,蒙自
city,532600,文山壮族苗族自治州
city,532800,西双版纳傣族自治州
city,532900,大理白族自治州
city,533100,德宏傣族景颇族自治州
city,533300,怒江傈僳族自治州
city,533400,迪庆藏族自治州
subdivision,540000,CN-54,西藏
city,540100,拉萨
city,540200,日喀则地区
city,540200,日喀则
city,542100,昌都地区
city,542100,昌都
city,542200,山南地区
city,542400,那曲地区
city,542500,阿里地区
city,542600,林芝地区
city,542600,林芝
subdivision,610000,CN-61,陕西
city,610100,西安
city,610200,铜川
city,610300,宝鸡
city,610400,咸阳
city,610500,渭南
city,610600,延安
city,610700,汉中
city,610800,榆林
city,610900,安康
city,611000,商洛
subdivision,620000,CN-62,甘肃
city,620100,兰州
city,620123,榆中
city,620200,嘉峪关
city,620300,金昌
city,620400,白银
city,620500,天水
city,620600,武威
city,620700,张掖
city,620800,平凉
city,620900,酒泉
city,621000,庆阳
city,621100,定西
city,621200,陇南
city,622900,临夏回族自治州
city,623000,甘南藏族自治州
city,623001,合作
subdivision,630000,CN-63,青海
city,630100,西宁
city,630200,海东
city,632200,海北藏族自治州
city,632300,黄南藏族自治州
city,632500,海南藏族自治州
city,632600,果洛藏族自治州
city,632700,玉树藏族自治州
city,632800,海西蒙古族藏族自治州
subdivision,640000,CN-64,宁夏
city,640100,银川
city,640200,石嘴山
city,640300,吴忠
city,640400,固原
city,640500,中卫
subdivision,650000,CN-65,新疆
city,650100,乌鲁木齐
city,650200,克拉玛依
city,652100,吐鲁番地区
city,652100,吐鲁番
city,652200,哈密地区
city,652300,昌吉回族自治州
city,652700,博尔塔拉蒙古自治州
city,652800,巴音郭楞蒙古自治州
city,652801,库尔勒
city,652900,阿克苏地区
city,653000,克孜勒苏柯尔克孜自治州
city,653100,喀什地区
city,653200,和田地区
city,654000,伊犁哈萨克自治州
city,654002,伊宁
city,654200,塔城地区
city,654300,阿勒泰地区
city,659001,石河子
city,659002,阿拉尔
subdivision,710000,CN-71,台湾
city,710000,澎湖县
city,710000,新北市
city,710000,台北市
city,710000,嘉义县
city,710000,高雄市
city,710000,台中市
city,710000,新竹市
city,710000,桃园市
city,710000,基隆市
city,710000,嘉义市
city,710000,云林县
city,710000,屏东县
city,710000,金门县
city,710000,台南市
city,710000,宜兰县
city,710000,花莲县
subdivision,810000,CN-81,香港
subdivision,820000,CN-82,澳门

country,162,CX,CXR,162,圣诞岛
country,166,CC,CCK,166,科科斯（基林）群岛
country,170,CO,COL,170,哥伦比亚
country,174,KM,COM,174,科摩罗
country,178,CG,COG,178,刚果(布)
country,180,CD,COD,180,刚果(金)
country,184,CK,COK,184,库克群岛
country,188,CR,CRI,188,哥斯达黎加
country,384,CI,CIV,384,科特迪瓦
country,191,HR,HRV,191,克罗地亚
country,192,CU,CUB,192,古巴
country,531,CW,CUW,531,库拉索
country,196,CY,CYP,196,塞浦路斯
country,203,CZ,CZE,203,捷克
country,208,DK,DNK,208,丹麦
country,262,DJ,DJI,262,吉布提
country,212,DM,DMA,212,多米尼克
country,214,DO,DOM,214,多米尼加
country,218,EC,ECU,218,厄瓜多尔
country,818,EG,EGY,818,埃及
country,222,SV,SLV,222,萨尔瓦多
country,226,GQ,GNQ,226,赤道几内亚
country,232,ER,ERI,232,厄立特里亚
country,233,EE,EST,233,爱沙尼亚
country,231,ET,ETH,231,埃塞俄比亚
country,238,FK,FLK,238,福克兰群岛
country,234,FO,FRO,234,法罗群岛
country,242,FJ,FJI,242,斐济
country,246,FI,FIN,246,芬兰
country,250,FR,FRA,250,法国
country,254,GF,GUF,254,法属圭亚那
country,258,PF,PYF,258,法属波利尼西亚
country,260,TF,ATF,260,法属南部领地
country,266,GA,GAB,266,加蓬
country,270,GM,GMB,270,冈比亚
country,268,GE,GEO,268,格鲁吉亚
country,276,DE,DEU,276,德国
country,288,GH,GHA,288,加纳
country,292,GI,GIB,292,直布罗陀
country,300,GR,GRC,300,希腊
country,304,GL,GRL,304,格陵兰
country,308,GD,GRD,308,格林纳达
country,312,GP,GLP,312,瓜德罗普
country,316,GU,GUM,316,关岛
country,320,GT,GTM,320,危地马拉
country,831,GG,GGY,831,根西岛
country,324,GN,GIN,324,几内亚
country,624,GW,GNB,624,几内亚比绍
country,328,GY,GUY,328,圭亚那
country,332,HT,HTI,332,海地
country,334,HM,HMD,334,赫德岛和麦克唐纳群岛
country,336,VA,VAT,336,梵蒂冈
country,340,HN,HND,340,洪都拉斯
country,344,HK,HKG,344,香港
country,348,HU,HUN,348,匈牙利
country,352,IS,ISL,352,冰岛
country,356,IN,IND,356,印度
country,360,ID,IDN,360,印度尼西亚
country,364,IR,IRN,364,伊朗
country,368,IQ,IRQ,368,伊拉克
country,372,IE,IRL,372,爱尔兰
country,833,IM,IMN,833,马恩岛
country,376,IL,ISR,376,以色列
country,380,IT,ITA,380,意大利
country,388,JM,JAM,388,牙买加
country,392,JP,JPN,392,日本
country,832,JE,JEY,832,泽西岛
country,400,JO,JOR,400,约旦
country,398,KZ,KAZ,398,哈萨克斯坦
country,404,KE,KEN,404,肯尼亚
country,296,KI,KIR,296,基里巴斯
country,408,KP,PRK,408,朝鲜
country,410,KR,KOR,410,韩国
country,414,KW,KWT,414,科威特
country,417,KG,KGZ,417,吉尔吉斯斯坦
country,418,LA,LAO,418,老挝
country,428,LV,LVA,428,拉脱维亚
country,422,LB,LBN,422,黎巴嫩
country,426,LS,LSO,426,莱索托
country,430,LR,LBR,430,利比里亚
country,434,LY,LBY,434,利比亚
country,438,LI,LIE,438,列支敦士登
country,440,LT,LTU,440,立陶宛
country,442,LU,LUX,442,卢森堡
country,446,MO,MAC,446,卢森堡
country,807,MK,MKD,807,马其顿
country,450,MG,MDG,450,马达加斯加
country,454,MW,MWI,454,马拉维
country,458,MY,MYS,458,马来西亚
country,462,MV,MDV,462,马尔代夫
country,466,ML,MLI,466,马里
country,470,MT,MLT,470,马耳他
country,584,MH,MHL,584,马绍尔群岛
country,474,MQ,MTQ,474,马提尼克岛
country,478,MR,MRT,478,毛里塔尼亚
country,480,MU,MUS,480,毛里求斯
country,175,YT,MYT,175,马约特
country,484,MX,MEX,484,墨西哥
country,583,FM,FSM,583,密克罗尼西亚联邦
country,498,MD,MDA,498,摩尔多瓦
country,492,MC,MCO,492,摩纳哥
country,496,MN,MNG,496,蒙古
country,499,ME,MNE,499,黑山
country,500,MS,MSR,500,蒙塞拉特岛
country,504,MA,MAR,504,摩洛哥
country,508,MZ,MOZ,508,莫桑比克
country,104,MM,MMR,104,缅甸
country,516,NA,NAM,516,纳米比亚
country,520,NR,NRU,520,瑙鲁
country,524,NP,NPL,524,尼泊尔
country,528,NL,NLD,528,荷兰
country,540,NC,NCL,540,新喀里多尼亚
country,554,NZ,NZL,554,新西兰
country,558,NI,NIC,558,尼加拉瓜
country,562,NE,NER,562,尼日尔
country,566,NG,NGA,566,尼日利亚
country,570,NU,NIU,570,纽埃岛
country,574,NF,NFK,574,诺福克岛
country,580,MP,MNP,580,北马里亚纳群岛
country,578,NO,NOR,578,挪威
country,512,OM,OMN,512,阿曼
country,586,PK,PAK,586,巴基斯坦
country,585,PW,PLW,585,帕劳
country,275,PS,PSE,275,巴勒斯坦
country,591,PA,PAN,591,巴拿马
country,598,PG,PNG,598,巴布亚新几内亚
country,600,PY,PRY,600,巴拉圭
country,604,PE,PER,604,秘鲁
country,608,PH,PHL,608,菲律宾
country,612,PN,PCN,612,皮特凯恩群岛
country,616,PL,POL,616,波兰
country,620,PT,PRT,620,葡萄牙
country,630,PR,PRI,630,波多黎各
country,634,QA,QAT,634,卡达
country,638,RE,REU,638,留尼汪岛
country,642,RO,ROU,642,罗马尼亚
country,643,RU,RUS,643,俄罗斯
country,646,RW,RWA,646,卢旺达
country,652,BL,BLM,652,圣巴泰勒米
country,654,SH,SHN,654,圣赫勒拿
country,659,KN,KNA,659,圣基茨和尼维斯
country,662,LC,LCA,662,圣卢西亚
country,663,MF,MAF,663,法属圣马丁
country,666,PM,SPM,666,圣皮埃尔和密克隆群岛
country,670,VC,VCT,670,圣文森特和格林纳丁斯
country,882,WS,WSM,882,萨摩亚
country,674,SM,SMR,674,圣马力诺
country,678,ST,STP,678,圣多美和普林西比
country,682,SA,SAU,682,沙特阿拉伯
country,686,SN,SEN,686,塞内加尔
country,688,RS,SRB,688,塞尔维亚
country,690,SC,SYC,690,塞舌尔
country,694,SL,SLE,694,塞拉利昂
country,702,SG,SGP,702,新加坡
country,534,SX,SXM,534,荷属圣马丁
country,703,SK,SVK,703,斯洛伐克
country,705,SI,SVN,705,斯洛文尼亚
country,090,SB,SLB,090,所罗门群岛
country,706,SO,SOM,706,索马里
country,710,ZA,ZAF,710,南非
country,239,GS,SGS,239,南乔治亚岛和南桑威奇群岛
country,728,SS,SSD,728,南苏丹
country,724,ES,ESP,724,西班牙
country,144,LK,LKA,144,斯里兰卡
country,729,SD,SDN,729,苏丹
country,740,SR,SUR,740,苏里南
country,744,SJ,SJM,744,斯瓦尔巴特群岛
country,748,SZ,SWZ,748,斯威士兰
country,752,SE,SWE,752,瑞典
country,756,CH,CHE,756,瑞士
country,760,SY,SYR,760,叙利亚
country,158,TW,TWN,158,中华民国
country,762,TJ,TJK,762,塔吉克斯坦
country,834,TZ,TZA,834,坦桑尼亚
country,764,TH,THA,764,泰国
country,626,TL,TLS,626,东帝汶
country,768,TG,TGO,768,多哥
country,772,TK,TKL,772,托克劳群岛
country,776,TO,TON,776,汤加
country,780,TT,TTO,780,特立尼达和多巴哥
country,788,TN,TUN,788,突尼斯
country,792,TR,TUR,792,土耳其
country,795,TM,TKM,795,土库曼斯坦
country,796,TC,TCA,796,特克斯和凯科斯群岛
country,798,TV,TUV,798,图瓦卢
country,800,UG,UGA,800,乌干达
country,804,UA,UKR,804,乌克兰
country,784,AE,ARE,784,阿联酋
country,826,GB,GBR,826,英国
country,840,US,USA,840,美国
country,581,UM,UMI,581,美国本土外小岛屿
country,858,UY,URY,858,乌拉圭
country,860,UZ,UZB,860,乌兹别克斯坦
country,548,VU,VUT,548,瓦努阿图
country,862,VE,VEN,862,委内瑞拉
country,704,VN,VNM,704,越南
country,092,VG,VGB,092,英属维尔京群岛
country,850,VI,VIR,850,美属维京群岛
country,876,WF,WLF,876,瓦利斯和富图纳群岛
country,732,EH,ESH,732,撒拉威阿拉伯民主共和国
country,887,YE,YEM,887,也门
country,894,ZM,ZMB,894,赞比亚
country,716,ZW,ZWE,716,津巴布韦

isp,0,移动
isp,0,联通,CHINAUNICOM
isp,0,电信,CHINATELECOM
isp,0,华数
isp,0,教育网
isp,0,鹏博士
isp,0,阿里云
isp,0,广电网
isp,0,铁通
isp,0,科技网
isp,0,有线通
isp,0,歌华有线
isp,0,数讯
isp,0,珠江宽频
isp,0,方正宽带
isp,0,京宽网络
isp,0,天威视讯
isp,0,中企通信
`)
)

func BuiltinWorld() []byte {
	return builtinWorld
}
