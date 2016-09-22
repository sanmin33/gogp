//    CopyRight @Ally Dale 2016
//    Author  : Ally Dale(vipally@gmail.com)
//    Blog    : http://blog.csdn.net/vipally
//    Site    : https://github.com/vipally

//package gogp implement a way to generate go-gp code from *.gp+*.gpg file
package gogp

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/vipally/cmdline"
	"github.com/vipally/cpright"
	"github.com/vipally/gogp/ini"
	xdebug "github.com/vipally/gx/debug"
)

const (
	g_gp_ext         = ".gp"
	g_gpg_ext        = ".gpg"
	g_go_code_ext    = ".go"
	g_gp_file_suffix = "gpg"
	g_gp_fmt         = "<%s>"

	//generic-programming flag <XXX>
	g_gp_regexp = `\<[[:alpha:]][[:word:]]{0,}\>`

	gpFileDir       = "GpFilePath" //read gp file from another path
	thisPackagePath = "github.com/vipally/gogp"

	version = "1.0.2"
)

func init() {
	copyRightCode = cmdline.ReplaceTags(copyRightCode)
	CopyRight(copyRightCode)
}

var (
	g_gp_sign      = regexp.MustCompile(g_gp_regexp)
	g_map_rep      = make(map[string]string)
	g_match_no_rep = false
	g_proc_line    = 0

	copyRightCode = "//    " + strings.Replace(cpright.CopyRight(), "\n", "\n//", strings.Count(cpright.CopyRight(), "\n")-1)
)

func CopyRight(s string) {
	copyRightCode = s
}

func Version() string {
	return version
}

//main func of gogp
func Work(dir string) (nGpg, nGp int, err error) {

	fmt.Printf("Processing path:[%s]\n", dir)
	files, e := collect_sub_files(dir, g_gpg_ext)
	if e != nil {
		err = e
	}
	nGpg = len(files)
	for _, v := range files {
		name := file_base(v)
		path_with_name := path.Join(dir, name)
		n, e := gen_gp_code_by_gpg(path_with_name)
		if e != nil {
			err = e
		}
		nGp += n
	}
	return
}

func gen_gp_code_by_gpg(path_with_name string) (nGen int, err error) {
	fmt.Printf("Processing:%s\n", path_with_name)
	gpg_file := path_with_name + g_gpg_ext
	if ini, err := ini.New(gpg_file); err == nil {
		gpg_imps := ini.Sections()
		for _, gpg_imp := range gpg_imps {
			gp_reg_srcs := ini.Keys(gpg_imp)
			g_map_rep = make(map[string]string) //clear map
			for _, gp_reg_src := range gp_reg_srcs {
				replace := ini.GetString(gpg_imp, gp_reg_src, "")
				if replace == "" {
					fmt.Println("[Warn:]", gpg_file, gpg_imp, gp_reg_src, "has no replace string")
				}
				match := fmt.Sprintf(g_gp_fmt, gp_reg_src)
				g_map_rep[match] = replace
			}
			if err = gen_gp_code_by_gp(path_with_name, gpg_imp); err == nil {
				nGen++
			}
		}
		fmt.Printf("%s finish gen %d file(s)\n", path_with_name, nGen)
	}
	return
}

func gen_gp_code_by_gp(path_with_name string, imp_name string) (err error) {
	var fin, fout *os.File
	fmt.Println("gen_gp_code_by_gp", path_with_name, imp_name, xdebug.BtFileLine())
	if gp, ok := g_map_rep[gpFileDir]; ok { //read gp file from another path
		path_with_name = gp
	}
	gp_file := path_with_name + g_gp_ext

	if fin, err = os.Open(gp_file); err != nil {
		return
	}
	defer fin.Close()

	code_file := get_code_file(path_with_name, imp_name)
	if fout, err = os.OpenFile(code_file,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm); err != nil {
		return
	}
	defer fout.Close()

	rd := bufio.NewReader(fin)
	wt := bufio.NewWriter(fout)
	if err = write_header(wt); err != nil {
		return
	}
	g_proc_line = 0
	g_match_no_rep = false
	for {
		line, e := rd.ReadString('\n')
		if line != "" {
			g_proc_line++
			reped_line, _ := gen_gp_code(line)
			wt.WriteString(reped_line)
		}
		if e != nil {
			break
		}
	}
	if err = wt.Flush(); err != nil {
		return
	}
	if g_match_no_rep {
		s := fmt.Sprintf("error:[%s].[%s] not every gp have been replaced\n", path_with_name, imp_name)
		fmt.Println(s)
		err = fmt.Errorf(s)
	}
	return
}

func gen_gp_code(src string) (r string, err error) {
	if strings.HasPrefix(src, "//") { //never replace comment line
		return src, nil
	}
	r = g_gp_sign.ReplaceAllStringFunc(src, match_replace)
	return
}

func write_header(wt *bufio.Writer) (err error) {
	s := fmt.Sprintf(`// This file was auto-generated by [gogp] tool
// Last modify at: [%s]
// !!!!!!!!!NEVER MODIFY IT MANUALLY!!!!!!!!!

`, time.Now().Format("Mon Jan 02 2006 15:04:05"))
	wt.WriteString(s)
	wt.WriteString(copyRightCode)
	wt.WriteString("\n\n")
	return
}

func get_code_file(path_with_name, imp_name string) (r string) {
	r = fmt.Sprintf("%s_%s_%s%s",
		path_with_name, g_gp_file_suffix, imp_name, g_go_code_ext)
	return
}

func match_replace(src string) (rep string) {
	if v, ok := g_map_rep[src]; ok {
		rep = v
	} else {
		fmt.Printf("error: at line %d, %s has no replacing\n", g_proc_line, src)
		rep = src
		g_match_no_rep = true
	}
	return
}

func collect_sub_files(_dir string,
	ext string) (subfiles []string, err error) {
	f, err := os.Open(_dir)
	if err != nil {
		return
	}
	defer f.Close()

	dirs, err := f.Readdir(0)
	if err != nil {
		return
	}
	//subfiles = make([]string, 0)
	for _, v := range dirs {
		if !v.IsDir() {
			filename := v.Name()
			if ext == "" || path.Ext(filename) == ext {
				subfiles = append(subfiles, filename)
			}
		}
	}
	return
}

func file_base(file_path string) (file string) {
	_, full := path.Split(file_path)
	ext := path.Ext(file_path)
	file = strings.TrimSuffix(full, ext)
	return
}
