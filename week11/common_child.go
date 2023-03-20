package main

import "fmt"

type Key struct {
	i, j int
}

func lcs(memo map[Key]int32, s1 string, s2 string, pair Key) int32 {
	if memo[pair] == 0 {
		if (pair.i >= len(s1)) || (pair.j >= len(s2)) {
			return 0
		} else if s1[pair.i] == s2[pair.j] {
			memo[pair] = 1 + lcs(memo, s1, s2, Key{i: pair.i + 1, j: pair.j + 1})
		} else {
			s1up := lcs(memo, s1, s2, Key{i: pair.i + 1, j: pair.j})
			s2up := lcs(memo, s1, s2, Key{i: pair.i, j: pair.j + 1})

			if s1up > s2up {
				memo[pair] = s1up
			} else {
				memo[pair] = s2up
			}
		}
	}
	return memo[pair]
}

func commonChild(s1 string, s2 string) int32 {
	// Write your code here
	memo := make(map[Key]int32)
	defer fmt.Println(memo)
	return lcs(memo, s1, s2, Key{i: 0, j: 0})
}

func main() {
	fmt.Println(
		commonChild(
			"UBBJXJGKLXGXTFBJYNLHQPULFILXLMPDQFWIYVBRSRFNETTEGXOHLBVOAJMHLZMTMPSJCPWJGHISUUIKDPPAWVQMZECIEIQUPLMFKENHVLKCJVDDSUPDOZXBZSRMWHLIHENULKFEXVCZIOVHRQHZWMDAYLDLGXSTMPDBGAMEBHOMOGGEBFRITAQVALWGINAOWMTRJLJHGEVJOPCVXZQVDBKOUPFJWHMRXULNNKRUUITTVIXYYSZDECBIBBIIWPDOEMFHDJKUSIFZNOTGIVDIJJVHPOGQHXWERMYNYGYOHJYGNOVFNWWERWMTBZOAXNHTCJIBOCUBXERSJORHOAMALOODYOHXIDEINMDWSNKUKFLMFHZLUIFOVDDSGPRJLUOLSFVNCVZUIWRFLKGCKPHGBWMXGNATTRKPOPYYMTHSRYHVXOGFDVRSKDVHHCTJIRAWOLDFDCRXROEVQXMDVOJSGHUPHDSGAJUFWLILUNCGSEPBIGBDNNAGCHBZHFWVUDPAQXLAXIPLSZWDLRQJOMILYKVJNMGBPARKUUTIHOAFIKIRWWUAZASRTOOYHPNHCANPDHGZEWHKAIBGDAAZLAYRWKRWUKBVTRAMLQKSRYVSRXRYVVBOUEGBMGXAZMTQMAJHKGQCRLAJORKHTVLPWFDOZFWHJXMMHNMQIJOBNALQQMPBHMIZCGRGQKOSTJHQXYXUOFKDYKHEISOYNVJSMLYSTEEBWSPFNWIYVBWEFGTUFRZZJGXSMUQGLKQNJPBLKDTXOZQEHXXGKHJZPUUOCWMXUNIBNTJCUFXRQWZGZSLYADLXFYOOPLJRLSYYYHQQUCFEAHNWZOMOUWAUYRFAWHDUPAOJHHPIWYLQRFXGRUCUOAMDNFVTTWTXREEUFCSXADRQKRLRTWYSOGJQSIQZQJHLFWOHTMTYETCJRBRVNYAWIOTXTVEVJMXSGDBESHKJMWQZLWBPSRQAIVURCOZJUXTKCDVJSTEEEXWZCHIFXCROXLHDDVGSFLRPNLRPNXFUUBATGIZJUUVUXMUWAZGKJBRRJAIVIACJTZICMLIKDJRMPVOHEQNAYZMNDMQFMAJNMQOZBXYIXTZDLUJAPISDFMILADTNCJMWSINHWZNGNJNWCRQBSFZAAQVOVQAHJTGFAQTYBEHMSWYGQOJTNJKVVBXZQHOCBOSSBKVRRVFGEXXRZEEOSTOLDFXEATRULHACAQXYAXDEXHFHHGOBXZNIBHEJMRLLNNHICIFTPCRWIPNCXYYRJRERNDVSZAOXIULLFQROQLUAZXDRSYGWOVNCLUBKKRNBEYKALCZDJCJMPFBQXIIMNOHYWEHYGWBRVGNZYVRHVIOSMCATCNDNHKJQVILGTTJIIYPYWECZWDHUGMYBGOAANHFHRUHHFSNXIPYDTEACTOWCUYEXZIOFVOYOYPRQQSBGCBQSFQMTHMCKADQWTMSPHSDYIMOHHWJADQFCRIZANPLWJDCXYXUKOXEMWVKHBISONBQJLODRNPRSSUMBIUKAINIRZEZSHKPTYYVXTGXCRXAULWYFDTRPWWIYDVGFVLNSLIAZHKLWCHVBUFJWINZWKZNVYJYKJYKJIBOBJRQGGVPADSKQHXPSZHXVAVIWQTXZBBDJMTVVFDMZMCWIPEAXLUJJECGMHETBGCJZWIAMKJPXDRYNYZDVYIXHFNOBGOPKUCXTUDFVZFAMWKBTHZDXQNUMFQRNZZNQPYBSWCFJDGOGBAVORCVMWPPRLOKGFCSUMMUBDHABUEYTGOPRLPYESPONOIQNMUQGBFGULJYRLESRQCYNXLOPTBTOZKPRABIYIUWRWIPWONTDPKNGMPHSEWHGNIIAVKUVIXIQZGPNSXCNPPLMIOXKMNUUIAUABOQMMYRJDVDAYAAYOXOWQAVGBMCMIZTVDGKQVXCWRUCCUHSHWXKLZIKJDKIZZBSMBGNYARWXMMKNIPMBKUKSYBVYLTHLIDEHXBHSNYVYJVINFOKPNQCAZIDRYAWTWODCYHUTVHDLBPAQIJJNDDEMVVADJXIJHHHJEVSCHPNJLAUGUPEZAZJEVTJQWCAUUZISUMSYJPRUBAOCAFMTUHXOWBPRAGKFYPLJUPHWHEEOOWKDXYHISNGQYJTADSPAGQMIVZQZLIZQGMMIOKMTKMCOHDZQRYGIKEJCKOBGHWUJZNNSDSSUPUDTPXNBBOKYZRFCORNYJEQXDXGDTIKQMKZNCANALAIPXFIDROXSCHDSPXUMHIHYRCETHUZQJHLCSNGEMBDDFPNONFMKDFXPYANYIYHSKMLYOJRQFOBWIGTEBZVLQDXNUDHZXNMESUJWWQSMZYRXTHXXCEAFUJVFHYYBEOJEVGBPHYMKWVHCCIARIBUISSSAPDIZRBRMSUBWYEKVUMGPJCBBOCDCLMCZBWBVYRQRREQPLHUNHMHTRGLXXXKZJBLJWYKEOJWIQEYBHWJCLQGHHLASDDMWAHFYZHCZJPXJWZQTDGKTOJRAHCTIDIVHRQUHSCMRDICYAAPQYZEQBMXUEPOTILSIJKEEOJXXFFXWDXBVZYTQJSAUAQKELWVTTQGXAHXJWBKBZWZVZADQSCHFQBMSXLCGNNGDMIPFWNKAVZIQLIHBMABUBXSGXXTREJZYBTFBAVIIKTDAIRPYWQZANPYLDJSHRZPMWIPZJPUQUTXJPKMRNOAYOHPFCVLINPOTRDAZWXZCJMDLHGOOLZTNQXMVPNOUZCDFEYATNLCODMOQLDSPYOAPQITRJLQYEBBTOEOHKJKCCETKIADLYLJOMWLGZSKPJQAPFJFOTPTRDCTUCHKTJMTVIFSCRCADLSGNNIUHXXHKOSTZEZMJRYIYCIWCZQPMJPFNBGGIRBTWQYDTVSKDNXDVVIFEJSKELBNQUGYDVKHOUMFGBABLJAYLZKASVYVLCJALPJCUGCULUCLCZIZUYNRPBOQJZHTCKIHEQYMTYLOHJLOEAKZBCGKAQRJSATNRZUSFOSEQXEEZCDBALXHUALTQVBBAWRFBIWKBESOYXPBZINJPUJIKHUJWTJLKHKKQWQTWZNLHSTHYUVINLTDXFNEFRFEGODSFWXGDWZHAVWRNROFTNYEVEBWKMZZURJKOHOHQTANWSMKBNGQOWQHGWWGWAQPNSFNFAKUJNEHRMBYJTWYXKQCWYVGLSEUZGBIFDCEJKVTUGCEKGOFGXDOPADDEZQSHPOWIDPMRCPEEIBBAIUJFVCGAFVSPOJSUPSRYAFBBPALHFKIBQKZEWNUTESHZCHPJHSQUCIEAKQXQWVNJKYHWCSKUEPDRSYGLYQDLFAHYCOCYOXCUKIDVVROCZFOCQCLIHYKOBQPTRAOPZFMOBQNZXHIPZFTDZJALTSKOYTEYAOAOCYEXBGRQYWUOEAVHQUXBPTMVDMIVOWWURCVCXJBSKGNKDMRKQCIYRYHQJTROFIGSITZACPYJDAYODNXWWIDMMNEQWEVLKJAHXERTAPWJHYDTVGSVXADZLUEEMPRSVRJCNAEOIAUJMCYYMKOSZFRXDSOKMXVXGGXSHZHJSPFKMNPBOBQRXFRGFTOGGCVNSWUTUQFVWQLRXMBRLIHOESBLCCINFAHGSAYJVDZOTGVAFDTCLAKQVQHRXPPWZHVOMUNCYWGCLBHXRCYOIZDHFWVRSGQHQQSXDHIGOTKUZAOZXPCROWQIZWBEDUJTYADZOMLAADXDOHDZJLNGXVENVCPWUPBZCWAWTIJKYTRUMITKRNHTFEUVDHVAGLXJCIQJCBTBTZGVHONTTMAMWRWFETPAKQJMDZPZLQPMEJICRKYBFPBHWLIOCXVDXQDEBHVLXOSYKYBSUAFHLOJQFNAUZGVVLMVINWBEJAGFJOTWFIDEIOXPZESJBZBXVCOYWTDDAILFJOKHJPVCALFABUJAHUEDMBPESMINFHTSSLJTQDTYREYQPNJYYDQBLBCUCEOYVGGZDMGGAVBOVPZHWCFMALQTQWMOEWUOJDYCCUFULJBHDDULAWCDXZTMAICBKRVXWXYTNUVNPLVZVIBKJEMKIXCDSVGDLNDXFMUNEOWSUWQCIYNPAEJCHGEUAKBXGCFGLJRFSTZHERMKGIQQYDYXMKUFESCWIQWKRBNYRASHYXUGYMRUZOGKZKNGQDWRXJUWAZBISCSLYBASSZUEHUVZIGHONWDSDVFTKUOECNMKYMWLLPNKYDINPKSLFYIBNERJIZUSELWCICKHZJVDBUIIXHGKRQGOGCSUIYVHWROOAPWMWMGBYWGBDSTFKXFUIWPJOBFEICKIJHHHLNVUDCEMLWEPKCANBMFGODZQRRXRIPKXBGYWAKOWIHFWWXCTZXJVCSHEXDZAXVQYDHMKTFTWVSZHPDPQFCBHYDT",
			"WZFPTGLCXKNHAKSFEIXYOHTDCCSDAKYASOBHHDTBEZXRKGKOJOLMEGTPSWQNRODAHQOYMOJDVGVSUSMDMLJLVQZYREJGKMKLXGBQEPYXVCTABIBGSWXFWXIIIKDAOXUWZLNYOWRWAUCQTACZPIZREJGTIBCACYDCGSAUULVWRCDLSMIAPCDPHJEOUPOLYPUDSEBYVIUMPLMGTFWCZKDZOHHSGRZDGGPXCUBINGTMOTEOHGLTPOUGCLLJJYFCJJATPZKJSVVKDBJOCJIUNGJJITGMBSYMJLMOCYEGTFECRMBFSBFLXCPSIOJZONIMPPKXETQEEYHUODXZCZJSGMRPFEIRUBUZMHSHRXNXKHMVEMFCSPQQGVEWOMJXSHDRBSPZTFFBAYKXSVULEHDYTDCWATCQCDDEUZINAXFLMWJLOPVWMAMSTMXKKOXSGLAWMGAIWTYWTPBWYEAKSWKSSDQTWVYDFHNTMLAGJBFUSVXUZCMSMDAUBPQOUYLIGVSWGVTASPISVRYGKCJTSKRXCMQOUVNMIIOMQZGZTGLUSRZXFVRLPKDSTTOTHEBQNJNCAICMFQRISSKYKTQWVLLGLLUCHDCKNXSWKGXHEVJGTLQFZJBAPYPUMYLQJPRJZCSRRPSTJQZGVFINHPRKVMHQUPUAAWVAGZIUMNXXGMWQAZXOFHPIHIMJCJMXVUMFSVYGXIIQOSFCRDXYIBHDAZFCJLQTYPVZUEWHCHABWOBSFZMSDHZUDPLRKWIHKJTJLUDZEGHRERGCEGXAWHXHPNEBANUYQVJDVGTFHOUPIVEHVLYZHELIWZPXLKMRLFIHTBBXPSCZZKFDYNHAHFBACRXGGLWLKZDCKPPHPYHLKVJQNXSZVXXBCREROFCITROSOSILHTCXPTNUETFSWJBNHMIIESTBUPFPRCODHDNHSWUEDBUQLYAPSJNHOVYMHXEOCUAHAMTXAFWRTPGMLIOVXDUKCFIVBRVXWTSJYHYTNGQGVPUEQPZAOLVGHDWBHBRHIFVZHAWZQDMNOUXTNJRPTIASYVQIJLBMMCRNYLHMBPIKGVDPPTGKZZNXXTYGNFKYNQKRCDBKHXYXMGJJQERHDUIGMBVNNDCECRCTAQEJRFKWEAZDJJBVSSGFLQYALBPKCPXUMIELWGRWWHLQLTKZFPKXWTMMIYHLOSZOGELNGJEHTIQKJXLKIUOCVWOZEXMWQMSVEAHQQMYHGIBMSQUPEZNDRJGPGGEYVSHEMAYEEDAMLGUUCOKTMKFJVQMKOWXOLEUCYMSTLWROBVFSIFFELMDPXRFRZMVXTSMBONXRSYSLKDLMTCCATCPNLQUTMIODOVTEBFFJVRXLRHJZLWFOJMQAPTTBOQPBQIUWRBHIVHKCELSUUVYKOHVIPHUGBLFZFEECTYORMHBFQMNDNYJANEAHZEGAGNHZOUYEVLAXJXLKYQKSSMITMPYUHTYXRSXTXYLMFYVXDMJFJCAXUFKJLNRDLERJAPAERTYJJFAUFPKYCILLMPDWKZYPWGHYIUCEDRYAAUAQXSPZMYVPGBJYBABGELVTBRFCMCBRQZIWTFOGMRTWZLTMFKPKJEIXTCQUHBIRIRUZBIRBLQDWDEVQBMNFFHJYVIFATHAZBWKVDPDCAVBQOJUUSBHKKNUVNMXSDDXOAWZWUDFLPBJEDAHLYCHTYXHEFQIJABDKILJYPWLHRXUZYFAJGSREEKIXHKTBLZBYFSJSHJVTNWZZIGJHZHSFWZHLEMBHOYFLQOUPJKLZCOUTAVAJWOTTWNFRZJRTZWUTOGNFTWQXJKRFIEZBEQLVCCHVKMOUOFWNPKAIISJHCZVMBQFBOYVWKAMMSZJTYLUXTOVDUNXKADHSTMFBDRGAIULKUOXIFWHJDXLCSFBIMJGELSAGXEALCLIUVCHESOHYLLJPDKYJICTLFTUAYLJUXZQRVHYGBFEDKEJSXLGBRZFEWDZUWTYRYRRVCZJBWEATAGHSHMABGFBQVXODNHJCGMSYKJETETHKZPJHCTMIGUDRMOKCSBKCQQYOTYHGOYJVAIFYCBKZDVXHSHRSFNIMZNBJLLBNKFXSRXPCOXJBYYEWYCAVCQSJYWGBEXWZIXNLGMNEZTVWSMAPWXHKVCXUCIQOBFDNSJKTMLGRGMHMRSYRZEWSUKKVAGTXWDSFHXTHBCBCQLVTQOPGHXKOZCPCQOGFADJTNNHTMCPPWKRCFFPAPPVQFZSUUHPIZCKOYFZNKSJKNLHFOBWPAZYTOJLRYHTNDLXGWDLMYHUOVBQQZARLKZGCJXHSQMFHJWIUYHNBSHHCOPCMXZLLDCFGLKCGXTYTWLOWNGGBEOZECGZCTMVRMKMWYSHAAACPEILLEPOKDJZKGHCUZKWFBPDGRGBXYUCAEMESCABKAETFXHLFGSDFCDEPRIZVJNWPOCNBALVTJXZQHJFRWLVAVPMRFLZVJAFRZJKUKYJTJRSFXCLMINWSSUFWGZNUNZWLVLZJRIXVWYDDNXCVVZHCTPFTPLEOGAWXLLAAEQUWVJLZKYSXNDOTDJUMSDCABLVCKYWYLGBZGLXEWHRFBOJAIGCDWITRGFDUYFCNEXOKFCRJWRFUUKABRXDXIKBFWQVKIXFEEWUVFTNLKJSXNGNMFQDBSFGPYLBMUVKSMORJWGKGQFGASYIZENJITLEEJGNTTLVABVGSBYOQZJQBOOEEVRHVJHYEXIDYJKWUCNSDDSNTTOIBKDAECIXNYKPJNNGDTXJUBDCNYUUMJPSACOYPZXJCTHULSHQLOAJSQNVNOYGSHRFXYSWXIIJMKPTTLCKUVDWTLITJEJPTQQXAFZHMPCKCNKADEGFBSMMCYSQDFEDQLPUAPPAOAIAHVPIRFWIMTBUZVFNCOBQHIOUJADJQJGHTKLADBXUNTPJEVELRPZXEKXPCHTRKVHBNRJNXAXFHQNRPQZJKGBBJENGTAXGBOUCNEKGOAIKLXEAABOWLXRKPAJBSNVFGOOWOWPNYMKCWQNYRRBYVDXTGRELHZHRWXNKCOYXFCGAMXVZWGWEHQZXNLSENXXURQHXGJFVJMGNPUAJZQEWZWXPYROYCVPLRFBCMDRKXZVEFYWGSVLQAFABEPSOOFOLUBQSYXFDRMWNWWYKMGPWHCYNGAEFBDTKPZXZBMOCTSSAUVLIOVYXVUBPKWOYMEDFDSGGJYVSKLQUGCYSIUCGTBYNPMLSMTMEEIZUDHFQVXSLPEQDLWQFJUNPFAKVDRHJSYDJZFLHIRECPPZZFUGVGAJJXPXVECCLHJGTYITISTNHQGCMFWMDJILIIRHJOOUCMFXVWXNSOJTJHDCRDUNDHQWRTRCECBUALWWMWIQDQBCPFAUSRNNWLBLXSYXAQMEVEOEILOHVEZHNUXVATZVXIBXLTVCVTDYFPZGAJGSMVDHHDWJZAGSRCQHGSZAZDQRVARLPYCTAIEKGHDHXPQTTWGQSCVGOLAUQPLTCVNKANFOSTJIYFPRZXHFZSWYCMDADPYGMTZXICKQGQBNMMHETBDBCCUGYILYVHXVDJVZGLXJNCFUDUDWKKCBBWWFYAAJXAUSVUDPHMPTKHORGLRSBXFPDNMQMJSYCMZMXGQWNPXRWZJEAXRXMJVKVMLYWKMHWPPROVLNBTFXMGLSRNPPWZKHMUORZURRSCFNLUOFAJJVOEZWXOZZBMCKGWWCCXPIUNUPSMYSDSKKYJXFMDVAHODUXEHDNAVZALDXZMKEBBTFLRANQHDCUSXBCNTSLVOHOOUIPJFASOCGOGCMTEVOLRCLOSASVQPTDMQQGAXNSHOTPJMETQCJWXLKQATYGZMHRMROBLDWEOAQPVPNPHUJYSXLPFBRCFJKWZLFYAUQIRVRRBIBXMEHMRDFUKXUTYWMKKVAYFTKGGXOMLNDLEVBMECEKOYSISHGRANUDISOXFELKXCSZARTZAVLJXUZISWMSLKDRBDHBEKGYYYZZJKBSQLBRUXPSDMJFCZSLLYGXTBHYCCKAUUKLUASWFYYPYXYQMHEHIXJVQIHBLABNHINYLVUWZGUXVNQWNVPHGZEHVQCCFNTKEPVABVBXHKEIHHYAJTVBCAUGPWJBQIDEIHFKVLTHEBOMPAYSQMMAVCFCXTRUMIMRZLNAJNRWKRLDTGNLDEACNVOJYUZTCXREGPLPKXFEWPOCESGDPYRJYYZSPWTPALTNZMYAVZDBIAVWKWVBWRPLRSWDHOAWAEIFLXVCTWQXAJHTNYBOBXQGWAICLFLMEICSYMDSKSTJQLTQHELYJTUDUQLNCCVHTSXBWSGGGJXQPPZVUYMYOFSQGSFLMQNWERJORSTYZYJVVKTSDBILTGYMBGAECEOWSAZSHLHWDRAHNEGYEKBR",
		),
	)
}