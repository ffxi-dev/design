package design

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// TestTokensMatchCSS checks that tokens.json holds exactly the tokens family.css declares, with the same values.
func TestTokensMatchCSS(t *testing.T) {
	var tokens map[string]string
	if err := json.Unmarshal(Tokens, &tokens); err != nil {
		t.Fatal(err)
	}
	root := regexp.MustCompile(`(?s):root \{(.*?)\n\}`).FindStringSubmatch(CSS)
	if root == nil {
		t.Fatal("family.css has no :root block")
	}
	css := map[string]string{}
	for _, m := range regexp.MustCompile(`(?m)^\s*--([\w-]+):\s*(.+?);$`).FindAllStringSubmatch(root[1], -1) {
		css[m[1]] = m[2]
	}
	if len(css) != len(tokens) {
		t.Errorf("family.css declares %d tokens, tokens.json %d", len(css), len(tokens))
	}
	for name, v := range css {
		if tokens[name] != v {
			t.Errorf("--%s: family.css %q, tokens.json %q", name, v, tokens[name])
		}
	}
}

// TestWordmarksAgree checks that every wrapper draws the glyph from glyph.svg and links sruon to sruon.dev.
func TestWordmarksAgree(t *testing.T) {
	paths := regexp.MustCompile(`d="([^"]+)"`).FindAllStringSubmatch(Glyph, -1)
	if len(paths) != 2 {
		t.Fatalf("glyph.svg has %d paths, want 2", len(paths))
	}
	wrappers := map[string]string{"go": string(Wordmark("site"))}
	for _, f := range []string{"react/Wordmark.js", "astro/Wordmark.astro", "python/ffxi_design/templates/ffxi_design/wordmark.html"} {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		wrappers[f] = string(b)
	}
	for name, src := range wrappers {
		for _, p := range paths {
			if !strings.Contains(src, p[1]) {
				t.Errorf("%s lacks glyph path %q", name, p[1])
			}
		}
		if !strings.Contains(src, "https://sruon.dev") {
			t.Errorf("%s does not link to https://sruon.dev", name)
		}
	}
}

// TestPackagesCarrySameCSS builds the npm and Python packages and checks both carry the family.css Go embeds.
func TestPackagesCarrySameCSS(t *testing.T) {
	for _, tool := range []string{"npm", "uv"} {
		if _, err := exec.LookPath(tool); err != nil {
			if os.Getenv("CI") != "" {
				t.Fatalf("%s not found", tool)
			}
			t.Skipf("%s not found", tool)
		}
	}
	dir := t.TempDir()
	run(t, "npm", "pack", "--pack-destination", dir)
	run(t, "uv", "build", "--wheel", "--out-dir", dir)

	tgz, _ := filepath.Glob(filepath.Join(dir, "*.tgz"))
	whl, _ := filepath.Glob(filepath.Join(dir, "*.whl"))
	if len(tgz) != 1 || len(whl) != 1 {
		t.Fatalf("want one npm and one Python package, got %v %v", tgz, whl)
	}
	if got := fromTarball(t, tgz[0], "package/family.css"); got != CSS {
		t.Error("the npm package's family.css differs from the Go module's")
	}
	if got := fromWheel(t, whl[0], "ffxi_design/static/family.css"); got != CSS {
		t.Error("the Python package's family.css differs from the Go module's")
	}

	var npm struct{ Version string }
	b, _ := os.ReadFile("package.json")
	json.Unmarshal(b, &npm)
	if !strings.Contains(filepath.Base(whl[0]), "-"+npm.Version+"-") {
		t.Errorf("package.json is %s but the wheel is %s; bump both", npm.Version, filepath.Base(whl[0]))
	}
}

func run(t *testing.T, name string, args ...string) {
	t.Helper()
	if out, err := exec.Command(name, args...).CombinedOutput(); err != nil {
		t.Fatalf("%s %v: %v\n%s", name, args, err, out)
	}
}

func fromTarball(t *testing.T, path, name string) string {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	tr := tar.NewReader(gz)
	for {
		h, err := tr.Next()
		if err != nil {
			t.Fatalf("%s not in %s: %v", name, path, err)
		}
		if h.Name == name {
			b, _ := io.ReadAll(tr)
			return string(b)
		}
	}
}

func fromWheel(t *testing.T, path, name string) string {
	t.Helper()
	z, err := zip.OpenReader(path)
	if err != nil {
		t.Fatal(err)
	}
	defer z.Close()
	f, err := z.Open(name)
	if err != nil {
		t.Fatalf("%s not in %s: %v", name, path, err)
	}
	defer f.Close()
	b, _ := io.ReadAll(f)
	return string(b)
}
