# gocra

**gocra** is a port to Cipres Restful API ([CRA](https://www.phylo.org/restusers/login.action))
using [golang](https://golang.org).

When using **gocra** alway consut the CRA [User Guide](https://www.phylo.org/restusers/docs/guide.html)
along with [gocra documentation](https://godoc.org/github.com/Godrigos/gocra).

**gocra** is capable of submmiting jobs, monitoring jobs status, download results and delete jobs.

You need a _cra_auth.json_ file (the source code contais an example that can be edited)
in your user home folder containing you credentials used to login to Cipres API.

```json
{
    "url": "https://cipresrest.sdsc.edu/cipresrest/v1",
    "appid": "Your application ID",
    "username": "Your Username",
    "password": "Your Password"
}
```

Here is an example of gocran package usage:

```golang
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/Godrigos/gocra"
)

func main() {

    var jr gocra.JobResults
    var total float64
    jl := gocra.ListJobs()
    count := len(jl.Jobs.Jobstatus)

    // Ends the execution if there is no job on server.
    if count == 0 {
        fmt.Println("There are no jobs on Cipres servers.")
        os.Exit(0)
    }

    for i := range jl.Jobs.Jobstatus {
        job := gocra.JobStat(jl.Jobs.Jobstatus[i].SelfURI.URL)
        wdir := gocra.WorkDir(job)
        jr = gocra.JobResult(job)
        // Destination folder for the downloaded files
        destDir := "/home/user/Downloads"

        /* Based on job state passed by jobStage field from a JobStatus structure
           prints to stdout the actual job stage and download its results if the
           job is completed. */
        switch {
        case job.JobStage == "QUEUE":
            fmt.Printf("The job \033[41;1;37m%s\033[0m has been validated"+
                " and placed in CIPRES's queue.\n", job.Metadata.Entry.Value)
            fmt.Printf("  \u25CF Job Handle: %s.\n", job.JobHandle)

        case job.JobStage == "COMMANDRENDERING":
            fmt.Printf("The job \033[41;1;37m%s\033[0m has reached the head"+
                " of the queue and CIPRES has created the command line that"+
                " will be run.\n", job.Metadata.Entry.Value)
            fmt.Printf("  \u25CF Job Handle: %s.\n", job.JobHandle)

        case job.JobStage == "INPUTSTAGING":
            fmt.Printf("CIPRES has created a temporary working directory"+
                " for the job \033[41;1;37m%s\033[0m on the execution host"+
                " and copied the input files over.\n",
                job.Metadata.Entry.Value)
            fmt.Printf("  \u25CF Job Handle: %s.\n", job.JobHandle)

        case job.JobStage == "SUBMITTED":
            fmt.Printf("The job \033[41;1;37m%s\033[0m has been submited"+
                " to the scheduler on the execution host.\n",
                job.Metadata.Entry.Value)
            fmt.Printf("  \u25CF Job Handle: %s.\n", job.JobHandle)
            fmt.Println("  \u25cf Working Directory Files List:")
            for j := range wdir.Jobfiles.Jobfile {
                fmt.Printf("      %-25s %12.2f kB\n",
                    wdir.Jobfiles.Jobfile[j].Filename,
                    float64(wdir.Jobfiles.Jobfile[j].Length)/1000)
                total += float64(wdir.Jobfiles.Jobfile[j].Length) / 1000
            }
            fmt.Printf("  \u25CF Total size %29.2f kB\n\n", total)

        case job.JobStage == "LOAD_RESULTS":
            fmt.Printf("The job \033[41;1;37m%s\033[0m has finished running"+
                " on the execution host and CIPRES has begun to transfer the"+
                " results.\n", job.Metadata.Entry.Value)
            fmt.Printf("  \u25CF Job Handle: %s.\n", job.JobHandle)

        case job.JobStage == "COMPLETED":
            fmt.Printf("Job \033[41;1;37m%s\033[0m results successfully"+
                " transferred and available.\n", job.Metadata.Entry.Value)
            fmt.Printf("  \u25CF Job Handle: %s.\n", job.JobHandle)
            fmt.Println("  \u25cf Results Files List:")
            for j := range jr.Jobfiles.Jobfile {
                fmt.Printf("      %-25s %12.2f kB\n",
                    jr.Jobfiles.Jobfile[j].Filename,
                    float64(jr.Jobfiles.Jobfile[j].Length)/1000)
                total += float64(jr.Jobfiles.Jobfile[j].Length) / 1000
            }
            fmt.Printf("\nDownloading files of job %s (%.2f MB)...\n",
                job.Metadata.Entry.Value, total/1000)
            // Parallel downloading
            gocra.DownloadAll(jr, job, destDir)
            gocra.Delete(job.SelfURI.URL)

        default:
            fmt.Println("Can not define job state. Try again later!")
        }

        /* When there is more than one job on CIPRES server, applies a delay
           between jobs access, acconding to the MinPollIntervalSeconds
           field from a JobStatus structure */
        if count > 1 {
            time.Sleep(time.Duration(job.MinPollIntervalSeconds) * time.Second)
            count--
        }
    }
}
```
