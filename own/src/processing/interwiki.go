package processing

import "regexp"

const (
	interwikiRegexp = "\\[\\[(af|als|am|an|ar|arc|arz|as|ast|az|ba|bar|bat-smg|be|be-x-old|bg|bh|bjn|bn|bpy|br|bs|bxr|ca|ceb|co|cs|cv|cy|da|de|diq|dsb|el|eo|es|et|eu|ext|fa|fi|fr|frp|frr|fur|fy|ga|gan|gl|gn|gv|haw|he|hi|hr|hsb|ht|hu|hy|ia|id|ig|ilo|io|is|it|ja|jv|ka|kk|kn|ko|ksh|ku|la|lb|li|lij|lmo|lt|lv|mg|mk|ml|mn|mr|ms|my|nah|nap|nds|ne|new|nl|nn|no|nrm|oc|or|os|pa|pcd|pfl|pi|pl|pms|pnb|pt|qu|ro|roa-tara|ru|rue|sa|sah|scn|sco|sh|si|simple|sk|sl|so|sq|sr|srn|sv|sw|szl|ta|te|tg|th|tk|tl|tr|tt|uk|ur|uz|vec|vi|vls|vo|war|wuu|xal|yi|yo|zh|zh-classical|zh-min-nan|zh-yue):(.*?)\\]\\]\\n?"
)

func stripInterwiki(in string) string {
	re := regexp.MustCompile(interwikiRegexp)
	return re.ReplaceAllString(in, "")
}
