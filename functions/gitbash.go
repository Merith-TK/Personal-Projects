package merith

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// please only use this package for msys2 environtments on windows

// MountPoints returns a list of moutpoints in a msys2 system
func MountPoints() [][]string {
	mountPoints := [][]string{{filepath.ToSlash(os.Getenv("HOME")), "~"}}

	// Load mount points
	out, err := exec.Command("mount").Output()
	if err != nil {
		//log.Error(err)
		log.Println(err)
	}
	lines := strings.Split(string(out), "\n")
	var mountRx = regexp.MustCompile(`^(.*) on (.*) type`)
	for _, line := range lines {
		extract := mountRx.FindStringSubmatch(line)
		if len(extract) > 0 {
			mountPoints = append(mountPoints, []string{extract[1], extract[2]})
		}
	}

	// Sort by size to get more restrictive mount points first
	sort.Slice(mountPoints, func(i, j int) bool {
		return len(mountPoints[i][0]) > len(mountPoints[j][0])
	})
	return mountPoints
}

// ConvertDir converts windows path to unix path based off msys2 mount points
func ConvertDir(dir string) string {

	if dir == "~" {
		return "~/"
	}

	//If msys feeds a `C:` path, clense
	if strings.HasPrefix(dir, "C:") {
		mountPoints := MountPoints()

		// Apply mount points
		absDir, _ := filepath.Abs(dir)
		absDir = filepath.ToSlash(absDir)
		for _, mp := range mountPoints {
			if strings.HasPrefix(absDir, mp[0]) {
				resolved := strings.Replace(absDir, mp[0], mp[1], 1)

				// Sometimes the resolved path can go from user input `/Workspace`
				// to `//Workspace`, this if statement checks that and removes it
				if strings.HasPrefix(resolved, "//") {
					resolved = strings.TrimPrefix(resolved, "/")
					return resolved
				}
				return resolved
			}
		}

	}
	return dir
}
