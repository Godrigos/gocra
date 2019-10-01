package gocra

import "encoding/xml"

/*
IQTree parameters structure
<!-- Usage: /projects/ps-ngbt/opt/comet/iqtree-omp-1.5.4-Linux/bin/iqtree-omp -s <alignment> [OPTIONS]

GENERAL OPTIONS:
  -? or -h             Print this help dialog
  -s <alignment>       Input alignment in PHYLIP/FASTA/NEXUS/CLUSTAL/MSF format
  -st <data_type>      BIN, DNA, AA, NT2AA, CODON, MORPH (default: auto-detect)
  -q <partition_file>  Edge-linked partition model (file in NEXUS/RAxML format)
 -spp <partition_file> Like -q option but allowing partition-specific rates
  -sp <partition_file> Edge-unlinked partition model (like -M option of RAxML)
  -t <start_tree_file> or -t BIONJ or -t RANDOM
                       Starting tree (default: 99 parsimony tree and BIONJ)
  -te <user_tree_file> Like -t but fixing user tree (no tree search performed)
  -o <outgroup_taxon>  Outgroup taxon name for writing .treefile
  -pre <PREFIX>        Prefix for all output files (default: aln/partition)
  -nt <#cpu_cores>     Number of cores/threads to use (REQUIRED)
  -seed <number>       Random seed number, normally used for debugging purpose
  -v, -vv, -vvv        Verbose mode, printing more messages to screen
  -quiet               Quiet mode, suppress printing to screen (stdout)
  -keep-ident          Keep identical sequences (default: remove & finally add)
  -safe                Safe likelihood kernel to avoid numerical underflow
  -mem RAM             Maximal RAM usage for memory saving mode

CHECKPOINTING TO RESUME STOPPED RUN:
  -redo                Redo analysis even for successful runs (default: resume)
  -cptime <seconds>    Minimum checkpoint time interval (default: 20)

LIKELIHOOD MAPPING ANALYSIS:
  -lmap <#quartets>    Number of quartets for likelihood mapping analysis
  -lmclust <clustfile> NEXUS file containing clusters for likelihood mapping
  -wql                 Print quartet log-likelihoods to .quartetlh file

NEW STOCHASTIC TREE SEARCH ALGORITHM:
  -ninit <number>      Number of initial parsimony trees (default: 100)
  -ntop <number>       Number of top initial trees (default: 20)
  -nbest <number>      Number of best trees retained during search (defaut: 5)
  -n <#iterations>     Fix number of iterations to stop (default: auto)
  -nstop <number>      Number of unsuccessful iterations to stop (default: 100)
  -pers <proportion>   Perturbation strength for randomized NNI (default: 0.5)
  -sprrad <number>     Radius for parsimony SPR search (default: 6)
  -allnni              Perform more thorough NNI search (default: off)
  -g <constraint_tree> (Multifurcating) topological constraint tree file

ULTRAFAST BOOTSTRAP:
  -bb <#replicates>    Ultrafast bootstrap (>=1000)
  -bnni Starting with IQ-TREE version 1.6 we provide a new option -bnni to reduce the risk
of overestimating branch supports with UFBoot due to severe model violations. With
this option UFBoot will further optimize each bootstrap tree using a hill-climbing nearest
neighbor interchange (NNI) search based directly on the corresponding bootstrap alignment.
  -wbt                 Write bootstrap trees to .ufboot file (default: none)
  -wbtl                Like -wbt but also writing branch lengths
  -nm <#iterations>    Maximum number of iterations (default: 1000)
  -nstep <#iterations> #Iterations for UFBoot stopping rule (default: 100)
  -bcor <min_corr>     Minimum correlation coefficient (default: 0.99)
  -beps <epsilon>      RELL epsilon to break tie (default: 0.5)

STANDARD NON-PARAMETRIC BOOTSTRAP:
  -b <#replicates>     Bootstrap + ML tree + consensus tree (>=100)
  -bc <#replicates>    Bootstrap + consensus tree
  -bo <#replicates>    Bootstrap only

SINGLE BRANCH TEST:
  -alrt <#replicates>  SH-like approximate likelihood ratio test (SH-aLRT)
  -alrt 0              Parametric aLRT test (Anisimova and Gascuel 2006)
  -abayes              approximate Bayes test (Anisimova et al. 2011)
  -lbp <#replicates>   Fast local bootstrap probabilities

MODEL-FINDER:
  -m TESTONLY          Standard model selection (like jModelTest, ProtTest)
  -m TEST              Standard model selection followed by tree inference
  -m MF                Extended model selection with FreeRate heterogeneity
  -m MFP               Extended model selection followed by tree inference
  -m TESTMERGEONLY     Find best partition scheme (like PartitionFinder)
  -m TESTMERGE         Find best partition scheme followed by tree inference
  -m MF+MERGE          Find best partition scheme incl. FreeRate heterogeneity
  -m MFP+MERGE         Like -m MF+MERGE followed by tree inference
  -rcluster <percent>  Percentage of partition pairs (relaxed clustering alg.)
  -mset program        Restrict search to models supported by other programs
                       (raxml, phyml or mrbayes)
  -mset m1,...,mk      Restrict search to models in a comma-separated list
                       (e.g. -mset WAG,LG,JTT)
  -msub source         Restrict search to AA models for specific sources
                       (nuclear, mitochondrial, chloroplast or viral)
  -mfreq f1,...,fk     Restrict search to using a list of state frequencies
                       (default AA: -mfreq FU,F; codon: -mfreq ,F1x4,F3x4,F)
  -mrate r1,...,rk     Restrict search to a list of rate-across-sites models
                       (e.g. -mrate E,I,G,I+G,R is used for -m MF)
  -cmin <kmin>         Min #categories for FreeRate model [+R] (default: 2)
  -cmax <kmax>         Max #categories for FreeRate model [+R] (default: 10)
  -merit AIC|AICc|BIC  Optimality criterion to use (default: all)
  -mtree               Perform full tree search for each model considered
  -mredo               Ignore model results computed earlier (default: reuse)
  -madd mx1,...,mxk    List of mixture models to also consider
  -mdef <nexus_file>   A model definition NEXUS file (see Manual)

SUBSTITUTION MODEL:
  -m <model_name>
                  DNA: HKY (default), JC, F81, K2P, K3P, K81uf, TN/TrN, TNef,
                       TIM, TIMef, TVM, TVMef, SYM, GTR, or 6-digit model
                       specification (e.g., 010010 = HKY)
              Protein: LG (default), Poisson, cpREV, mtREV, Dayhoff, mtMAM,
                       JTT, WAG, mtART, mtZOA, VT, rtREV, DCMut, PMB, HIVb,
                       HIVw, JTTDCMut, FLU, Blosum62, GTR20
      Protein mixture: C10,...,C60, EX2, EX3, EHO, UL2, UL3, EX_EHO, LG4M, LG4X
               Binary: JC2 (default), GTR2
      Empirical codon: KOSI07, SCHN05
    Mechanistic codon: GY (default), MG, MGK, GY0K, GY1KTS, GY1KTV, GY2K,
                       MG1KTS, MG1KTV, MG2K
 Semi-empirical codon: XX_YY where XX is empirical and YY is mechanistic model
       Morphology/SNP: MK (default), ORDERED
            Otherwise: Name of file containing user-model parameters
                       (rate parameters and state frequencies)
  -m <model_name>+F or +FO or +FU or +FQ (default: auto)
                       counted, optimized, user-defined, equal state frequency
  -m <model_name>+F1x4 or +F3x4
                       Codon frequencies
  -m <model_name>+ASC  Ascertainment bias correction for morphological/SNP data
  -m "MIX{m1,...mK}"   Mixture model with K components
  -m "FMIX{f1,...fK}"  Frequency mixture model with K components
  -mwopt               Turn on optimizing mixture weights (default: auto)

RATE HETEROGENEITY AMONG SITES:
  -m <model_name>+I or +G[n] or +I+G[n] or +R[n]
                       Invar, Gamma, Invar+Gamma, or FreeRate model where n is
                       number of categories (default: n=4)
  -a <Gamma_shape>     Gamma shape parameter for site rates (default: estimate)
  -amin <min_shape>    Min Gamma shape parameter for site rates (default: 0.02)
  -gmedian             Median approximation for +G site rates (default: mean)
  -*opt-gamma-inv      More thorough estimation for +I+G model parameters
  -i <p_invar>         Proportion of invariable sites (default: estimate)
  -wsr                 Write site rates to .rate file
  -mh                  Computing site-specific rates to .mhrate file using
                       Meyer & von Haeseler (2003) method

SITE-SPECIFIC FREQUENCY MODEL:
  -ft <tree_file>      Input tree to infer site frequency model
  -fs <in_freq_file>   Input site frequency model file
  -fmax                Posterior maximum instead of mean approximation

CONSENSUS RECONSTRUCTION:
  -t <tree_file>       Set of input trees for consensus reconstruction
  -minsup <threshold>  Min split support in range [0,1]; 0.5 for majority-rule
                       consensus (default: 0, i.e. extended consensus)
  -bi <burnin>         Discarding <burnin> trees at beginning of <treefile>
  -con                 Computing consensus tree to .contree file
  -net                 Computing consensus network to .nex file
  -sup <target_tree>   Assigning support values for <target_tree> to .suptree
  -suptag <name>       Node name (or ALL) to assign tree IDs where node occurs

ROBINSON-FOULDS DISTANCE:
  -rf_all              Computing all-to-all RF distances of trees in <treefile>
  -rf <treefile2>      Computing all RF distances between two sets of trees
                       stored in <treefile> and <treefile2>
  -rf_adj              Computing RF distances of adjacent trees in <treefile>

TREE TOPOLOGY TEST:
  -z <trees_file>      Evaluating a set of user trees
  -zb <#replicates>    Performing BP,KH,SH,ELW tests for trees passed via -z
  -zw                  Also performing weighted-KH and weighted-SH tests
  -au                  Also performing approximately unbiased (AU) test

GENERATING RANDOM TREES:
  -r <num_taxa>        Create a random tree under Yule-Harding model
  -ru <num_taxa>       Create a random tree under Uniform model
  -rcat <num_taxa>     Create a random caterpillar tree
  -rbal <num_taxa>     Create a random balanced tree
  -rcsg <num_taxa>     Create a random circular split network
  -rlen <min_len> <mean_len> <max_len>
                       min, mean, and max branch lengths of random trees

MISCELLANEOUS:
  -wt                  Write locally optimal trees into .treels file
  -blfix               Fix branch lengths of user tree passed via -te
  -blscale             Scale branch lengths of user tree passed via -t
  -blmin               Min branch length for optimization (default 0.000001)
  -blmax               Max branch length for optimization (default 100)
  -wsr                 Write site rates and categories to .rate file
  -wsl                 Write site log-likelihoods to .sitelh file
  -wslr                Write site log-likelihoods per rate category
  -wslm                Write site log-likelihoods per mixture class
  -wslmr               Write site log-likelihoods per mixture+rate class
  -wspr                Write site probabilities per rate category
  -wspm                Write site probabilities per mixture class
  -wspmr               Write site probabilities per mixture+rate class
  -wpl                 Write partition log-likelihoods to .partlh file
  -fconst f1,...,fN    Add constant patterns into alignment (N=#nstates)
  -me <epsilon>        LogL epsilon for parameter estimation (default 0.01)
  -*no-outfiles        Suppress printing output files
 -->
<!-- To maintain IQ-TREE, support users and secure fundings, it is important for us that you cite the following papers, whenever the corresponding features were applied for your analysis.
Example 1: "...We obtained branch supports with the ultrafast bootstrap (Minh et al. 2013) implemented in the IQ-TREE software (Nguyen et al. 2015)..."

 Example 2: "...We inferred the maximum-likelihood tree using the edge-linked partition model in IQ-TREE (Chernomor et al. 2016; Nguyen et al. 2015)..."


General Options:

Usage and meaning

-h or -? Print help usage.
-s Specify input alignment file in PHYLIP, FASTA, NEXUS, CLUSTAL or MSF format.
-st Specify sequence type as either of DNA, AA, BIN, MORPH, CODON or NT2AA for DNA, amino-acid, binary, morphological, codon or DNA-to-AA-translated sequences. This is only necessary if IQ-TREE did not detect the sequence type correctly. Note that -st CODON is always necessary when using codon models (otherwise, IQ-TREE applies DNA models) and you also need to specify a genetic code like this if differed from the standard genetic code. -st NT2AA tells IQ-TREE to translate protein-coding DNA into AA sequences and then subsequent analysis will work on the AA sequences. You can also use a genetic code like -st NT2AA5 for the Invertebrate Mitochondrial Code (see genetic code table).
-t Specify a file containing starting tree for tree search. The special option -t BIONJ starts tree search from BIONJ tree and -t RANDOM starts tree search from completely random tree. DEFAULT: 100 parsimony trees + BIONJ tree
-te Like -t but fixing user tree. That means, no tree search is performed and IQ-TREE computes the log-likelihood of the fixed user tree.
-o Specify an outgroup taxon name to root the tree. The output tree in .treefile will be rooted accordingly. DEFAULT: first taxon in alignment
-pre Specify a prefix for all output files. DEFAULT: either alignment file name (-s) or partition file name (-q, -spp or -sp)
-nt Specify the number of CPU cores for iqtree-omp version. A special option -nt AUTO will tell IQ-TREE to automatically determine the best number of cores given the current data and computer.
-seed Specify a random number seed to reproduce a previous run. This is normally used for debugging purposes. DEFAULT: based on current machine clock
-v Turn on verbose mode for printing more messages to screen. This is normally used for debugging purposes. DFAULT: OFF
-quiet Silent mode, suppress printing to the screen. Note that .log file is still written.
-keep-ident Keep identical sequences in the alignment. Bu default: IQ-TREE will remove them during the analysis and add them in the end.
-safe Turn on safe numerical mode to avoid numerical underflow for large data sets with many sequences (typically in the order of thousands). This mode is automatically turned on when having more than 2000 sequences.
-mem Specify maximal RAM usage, for example, -mem 64G to use at most 64 GB of RAM. By default, IQ-TREE will try to not to exceed the computer RAM size.

Restarting:

Option

Usage and meaning

-redo Redo the entire analysis no matter if it was stopped or successful. WARNING: This option will overwrite all existing output files.
-cptime Specify the minimum checkpoint time interval in seconds (default: 20s)


Likelihood mapping analysis

Starting with version 1.4.0, IQ-TREE implements the likelihood mapping approach (Strimmer and von Haeseler, 1997) to assess the phylogenetic information of an input alignment. The detailed results will be printed to .iqtree report file. The likelihood mapping plots will be printed to .lmap.svg and .lmap.eps files.

Compared with the original implementation in TREE-PUZZLE, IQ-TREE is much faster and supports many more substitution models (including partition and mixture models).

Usage and meaning

-lmap Specify the number of quartets to be randomly drawn. If you specify -lmap ALL, all unique quartets will be drawn, instead.
-lmclust Specify a NEXUS file containing taxon clusters (see below for example) for quartet mapping analysis.
-wql Write quartet log-likelihoods into .lmap.quartetlh file (typically not needed).
-n 0 Skip subsequent tree search, useful when you only want to assess the phylogenetic information of the alignment.

Automatic model selection
The default model (e.g., HKY+F for DNA, LG for protein data) may not fit well to the data. Therefore, IQ-TREE allows to automatically determine the
best-fit model via a series of -m TEST... option:

Usage and meaning

-m TESTONLY Perform standard model selection like jModelTest (for DNA) and ProtTest (for protein). Moreover, IQ-TREE also works for codon, binary and morphogical data.
-m TEST Like -m TESTONLY but immediately followed by tree reconstruction using the best-fit model found. So this performs both model selection and tree inference within a single run.
-m TESTNEWONLY or -m MF Perform an extended model selection that additionally includes FreeRate model compared with -m TESTONLY. Recommended as replacement for -m TESTONLY. Note that LG4X is a FreeRate model, but by default is not included because it is also a protein mixture model. To include it, use -madd option (see table below).
-m TESTNEW or -m MFP Like -m MF but immediately followed by tree reconstruction using the best-fit model found.
-m TESTMERGEONLY Select best-fit partitioning scheme like PartitionFinder.
-m TESTMERGE Like -m TESTMERGEONLY but immediately followed by tree reconstruction using the best partitioning scheme found.
-m TESTNEWMERGEONLY or -m MF+MERGE Like -m TESTMERGEONLY but additionally includes FreeRate model.
-m TESTNEWMERGE or -m MFP+MERGE Like -m MF+MERGE but immediately followed by tree reconstruction using the best partitioning scheme found.

“TIP: During model section run, IQ-TREE will write a file with suffix .model that stores information of all models tested so far. Thus, if IQ-TREE is interrupted for whatever reason, restarting the run will load this file to reuse the computation. Thus, this file acts like a checkpoint to resume the model selection.

Several parameters can be set to e.g. reduce computations:

Option

Usage and meaning

-rcluster Specify the percentage for the relaxed clustering algorithm (Lanfear et al., 2014). This is similar to -*rcluster-percent option of PartitionFinder. For example, with -rcluster 10 only the top 10% partition schemes are considered to save computations.
-mset Specify the name of a program (raxml, phyml or mrbayes) to restrict to only those models supported by the specified program. Alternatively, one can specify a comma-separated list of base models. For example, -mset WAG,LG,JTT will restrict model selection to WAG, LG, and JTT instead of all 18 AA models to save computations.
-msub Specify either nuclear, mitochondrial, chloroplast or viral to restrict to those AA models designed for specified source.
-mfreq Specify a comma-separated list of frequency types for model selection. DEFAULT: -mfreq FU,F for protein models (FU = AA frequencies given by the protein matrix, F = empirical AA frequencies from the data), -mfreq ,F1x4,F3x4,F for codon models
-mrate Specify a comma-separated list of rate heterogeneity types for model selection. DEFAULT: -mrate E,I,G,I+G for standard procedure, -mrate E,I,G,I+G,R for new selection procedure
-cmin Specify minimum number of categories for FreeRate model. DEFAULT: 2
-cmax Specify maximum number of categories for FreeRate model. It is recommended to increase if alignment is long enough. DEFAULT: 10
-merit Specify either AIC, AICc or BIC for the optimality criterion to apply for new procedure. DEFAULT: all three criteria are considered
-mtree Turn on full tree search for each model considered, to obtain more accurate result. Only recommended if enough computational resources are available. DEFAULT: fixed starting tree
-mredo Ignore .model file computed earlier. DEFAULT: .model file (if exists) is loaded to reuse previous computations
-madd Specify a comma-separated list of mixture models to additionally consider for model selection. For example, -madd LG4M,LG4X to additionally include these two protein mixture models.
-mdef Specify a NEXUS model file to define new models.

“NOTE: Some of the above options require a comma-separated list, which should not contain any empty space!

Example usages:

•Select best-fit model for alignment data.phy based on Bayesian information criterion (BIC):
  iqtree -s data.phy -m TESTONLY

•Select best-fit model for a protein alignment prot.phy using the new testing procedure and only consider WAG, LG and JTT matrix to save time:
  iqtree -s prot.phy -m MF -mset WAG,LG,JTT

Specifying substitution models
-m is a powerful option to specify substitution models, state frequency and rate heterogeneity type. The general syntax is:
-m MODEL+FreqType+RateType

where MODEL is a model name, +FreqType (optional) is the frequency type and +RateType (optional) is the rate heterogeneity type.

The following MODELs are available:

DataType

Model names
DNA JC/JC69, F81, K2P/K80, HKY/HKY85, TN/TrN/TN93, TNe, K3P/K81, K81u, TPM2, TPM2u, TPM3, TPM3u, TIM, TIMe, TIM2, TIM2e, TIM3, TIM3e, TVM, TVMe, SYM, GTR and 6-digit specification. See DNA models for more details.
Protein BLOSUM62, cpREV, Dayhoff, DCMut, FLU, HIVb, HIVw, JTT, JTTDCMut, LG, mtART, mtMAM, mtREV, mtZOA, Poisson, PMB, rtREV, VT, WAG.
Protein Mixture models: C10, …, C60 (CAT model) (Lartillot and Philippe, 2004), EX2, EX3, EHO, UL2, UL3, EX_EHO, LG4M, LG4X, CF4. See Protein models for more details.
Codon MG, MGK, MG1KTS, MG1KTV, MG2K, GY, GY1KTS, GY1KTV, GY2K, ECMK07/KOSI07, ECMrest, ECMS05/SCHN05 and combined empirical-mechanistic models. See Codon models for more details.
Binary JC2, GTR2. See Binary and morphological models for more details.
Morphology MK, ORDERED. See Binary and morphological models for more details.

The following FreqTypes are supported: FreqType

Meaning
+F Empirical state frequency observed from the data.
+FO State frequency optimized by maximum-likelihood from the data. Note that this is with letter-O and not digit-0.
+FQ Equal state frequency.
+F1x4 See Codon frequencies.
+F3x4 See Codon frequencies.

Usage and meaning
-mwopt Turn on optimizing weights of mixture models. Note that for models like LG+C20+F+G this mode is automatically turned on, but not for LG+C20+G.

Example usages:

•Infer an ML tree for a DNA alignment dna.phy under GTR+I+G model:
  iqtree -s dna.phy -m GTR+I+G


Rate heterogeneity - The following RateTypes are supported:
RateType

+I Allowing for a proportion of invariable sites.
+G Discrete Gamma model (Yang, 1994) with default 4 rate categories. The number of categories can be changed with e.g. +G8.
+I+G Invariable site plus discrete Gamma model (Gu et al., 1995).
+R FreeRate model (Yang, 1995; Soubrier et al., 2012) that generalizes +G by relaxing the assumption of Gamma-distributed rates. The number of categories can be specified with e.g. +R6. DEFAULT: 4 categories
+I+R invariable site plus FreeRate model.

See Rate heterogeneity across sites for more details.

Further options:
Usage and meaning

-a Specify the Gamma shape parameter (default: estimate)
-gmedian Perform the median approximation for Gamma rate heterogeneity instead of the default mean approximation (Yang, 1994)
-i Specify the proportion of invariable sites (default: estimate)
-*opt-gamma-inv Perform more thorough estimation for +I+G model parameters
-wsr Write per-site rates to .rate file

Optionally, one can specify Ascertainment bias correction by appending +ASC to the model string. Advanced mixture models can also be specified via MIX{...} and FMIX{...} syntax. Option -mwopt can be used to turn on optimizing weights of mixture models.

Partition model options - Partition models are used for phylogenomic data with multiple genes. You first have to prepare a partition file in NEXUS or RAxML-style format. Then use the following options to input the partition file:
Usage and meaning

-q Specify partition file for edge-equal partition model. That means, all partitions share the same set of branch lengths (like -q option of RAxML).
-spp Like -q but allowing partitions to have different evolutionary speeds (edge-proportional partition model).
-sp Specify partition file for edge-unlinked partition model. That means, each partition has its own set of branch lengths (like -M option of RAxML). This is the most parameter-rich partition model to accomodate heterotachy.

Site-specific frequency model options - The site-specific frequency model is used to substantially reduce the time and memory requirement compared with full profile mixture models C10 to C60. For full details see site-specific frequency model. To use this model you have to specify a profile mixture model with e.g. -m LG+C20+F+G together with a guide tree or a site frequency file:

Option

Usage and meaning

-ft Specify a guide tree (in Newick format) to infer site frequency profiles.
-fs Specify a site frequency file, e.g. the .sitefreq file obtained from -ft run. This will save memory used for the first phase of the analysis.
-fmax Switch to posterior maximum mode for obtaining site-specific profiles. Default: posterior mean.

With -fs option you can input a file containing your own site frequency profiles. The format of this file is that each line contains the site ID (starting from 1) and the state frequencies (20 for amino-acid) separated by white space. So it has as many lines as the number of sites in the alignment. The order of amino-acids is:
 A   R   N   D   C   Q   E   G   H   I   L   K   M   F   P   S   T   W   Y   V


Tree search parameters - The new IQ-TREE search algorithm (Nguyen et al., 2015) has several parameters that can be changed with:
Usage and meaning

-ninit Specify number of initial parsimony trees. DEFAULT: 100. Here the PLL library (Flouri et al., 2015) is used, which implements the randomized stepwise addition and parsimony subtree pruning and regafting (SPR).
-ntop Specify number of top initial parsimony trees to optimize with ML nearest neighbor interchange (NNI) search to initialize the candidate set. DEFAULT: 20
-nbest Specify number of trees in the candidate set to maintain during ML tree search. DEFAULT: 5
-nstop Specify number of unsuccessful iterations to stop. DEFAULT: 100
-n Specify number of iterations to stop. This option overrides -nstop criterion.
-sprrad Specify SPR radius for the initial parsimony tree search. DEFAULT: 6
-pers Specify perturbation strength (between 0 and 1) for randomized NNI. DEFAULT: 0.5
-allnni Turn on more thorough and slower NNI search. It means that IQ-TREE will consider all possible NNIs instead of only those in the vicinity of previously applied NNIs. DEFAULT: OFF
-djc Avoid computing ML pairwise distances and BIONJ tree.
-g Specify a topological constraint tree file in NEWICK format. The constraint tree can be a multifurcating tree and need not to include all taxa.

“NOTE: While the default parameters were empirically determined to work well under our extensive benchmark (Nguyen et al., 2015), it might not hold true for all data sets. If in doubt that tree search is still stuck in local optima, one should repeat analysis with at least 10 IQ-TREE runs. Moreover, our experience showed that -pers and -nstop are the most relevant options to change in such case. For example, data sets with many short sequences should be analyzed with smaller perturbation strength (e.g. -pers 0.2) and larger number of stop iterations (e.g. -nstop 500).

Example usages:

•Infer an ML tree for an alignment data.phy with increased stopping iteration of 500 and reduced perturbation strength of 0.2:
  iqtree -s data.phy -m TEST -nstop 500 -pers 0.2

•Infer an ML tree for an alignment data.phy obeying a topological constraint tree constraint.tree:
  iqtree -s data.phy -m TEST -g constraint.tree



Ultrafast bootstrap parameters - The ultrafast bootstrap (UFBoot) approximation (Minh et al., 2013) has several parameters that can be changed with:

Usage and meaning

-bb Specify number of bootstrap replicates (>=1000).
-wbt Turn on writing bootstrap trees to .ufboot file. DEFAULT: OFF
-wbtl Like -wbt but bootstrap trees written with branch lengths. DEFAULT: OFF
-nm Specify maximum number of iterations to stop. DEFAULT: 1000
-bcor Specify minimum correlation coefficient for UFBoot convergence criterion. DEFAULT: 0.99
-nstep Specify iteration interval checking for UFBoot convergence. DEFAULT: every 100 iterations
-beps Specify a small epsilon to break tie in RELL evaluation for bootstrap trees. DEFAULT: 0.5
-bspec Specify the resampling strategies for partitioned analysis. By default, IQ-TREE resamples alignment sites within partitions. With -bspec GENE IQ-TREE will resample partitions. With -bspec GENESITE IQ-TREE will resample partitions and then resample sites within resampled partitions (Gadagkar et al., 2005).

Example usages:

•Select best-fit model, infer an ML tree and perform ultrafast bootstrap with 1000 replicates:
  iqtree -s data.phy -m TEST -bb 1000

•Reconstruct ML and perform ultrafast bootstrap (5000 replicates) under a partition model file partition.nex:
  iqtree -s data.phy -spp partition.nex -m TEST -bb 5000

Nonparametric bootstrap - The slow standard nonparametric bootstrap (Felsenstein, 1985) can be run with:

Usage and meaning

-b Specify number of bootstrap replicates (recommended >=100). This will perform both bootstrap and analysis on original alignment and provide a consensus tree.
-bc Like -b but omit analysis on original alignment.
-bo Like -b but only perform bootstrap analysis (no analysis on original alignment and no consensus tree).

Single branch tests - The following single branch tests are faster than all bootstrap analysis and recommended for extremely large data sets (e.g., >10,000 taxa):

-alrt Specify number of replicates (>=1000) to perform SH-like approximate likelihood ratio test (SH-aLRT) (Guindon et al., 2010). If number of replicates is set to 0 (-alrt 0), then the parametric aLRT test (Anisimova and Gascuel 2006) is performed, instead of SH-aLRT.
-abayes Perform approximate Bayes test (Anisimova et al., 2011).
-lbp Specify number of replicates (>=1000) to perform fast local bootstrap probability method (Adachi and Hasegawa, 1996).

“TIP: One can combine all these tests (also including UFBoot -bb option) within a single IQ-TREE run. Each branch in the resulting tree will be assigned several support values separated by slash (/), where the order of support values is stated in the .iqtree report file.


Example usages:

•Infer an ML tree and perform SH-aLRT test as well as ultrafast bootstrap with 1000 replicates:
  iqtree -s data.phy -m TEST -alrt 1000 -bb 1000

Tree topology tests - IQ-TREE provides a number of tests for significant topological differences between trees:

Usage and meaning
-z Specify a file containing a set of trees. IQ-TREE will compute the log-likelihoods of all trees.
-zb Specify the number of RELL (Kishino et al., 1990) replicates (>=1000) to perform several tree topology tests for all trees passed via -z. The tests include bootstrap proportion (BP), KH test (Kishino and Hasegawa, 1989), SH test (Shimodaira and Hasegawa, 1999) and expected likelihood weights (ELW) (Strimmer and Rambaut, 2002).
-zw Used together with -zb to additionally perform the weighted-KH and weighted-SH tests.
-au Used together with -zb to additionally perform the approximately unbiased (AU) test (Shimodaira, 2002). Note that you have to specify the number of replicates for the AU test via -zb.
-n 0 Only estimate model parameters on an initial parsimony tree and ignore a full tree search to save time.
-te Specify a fixed user tree to estimate model parameters. Thus it behaves like -n 0 but uses a user-defined tree instead of parsimony tree.

“NOTE: The AU test implementation in IQ-TREE is much more efficient than the original CONSEL by supporting SSE, AVX and multicore parallelization. Moreover, it is more appropriate than CONSEL for partition analysis by bootstrap resampling sites within partitions, whereas CONSEL is not partition-aware.

Example usages:

•Given alignment data.phy, test a set of trees in data.trees using AU test with 10,000 replicates:
  iqtree -s data.phy -m GTR+G -n 0 -z data.trees -zb 10000 -au

•Same above but for a partitioned data partition.nex and additionally performing weighted test:
  iqtree -s data.phy -spp partition.nex -n 0 -z data.trees -zb 10000 -au -zw


Constructing consensus tree - IQ-TREE provides a fast implementation of consensus tree construction for post analysis:
Usage and meaning

-t Specify a file containing a set of trees.
-con Compute consensus tree of the trees passed via -t. Resulting consensus tree is written to .contree file.
-net Compute consensus network of the trees passed via -t. Resulting consensus network is written to .nex file.
-minsup Specify a minimum threshold (between 0 and 1) to keep branches in the consensus tree. -minsup 0.5 means to compute majority-rule consensus tree. DEFAULT: 0 to compute extended majority-rule consensus.
-bi Specify a burn-in, which is the number of beginning trees passed via -t to discard before consensus construction. This is useful e.g. when summarizing trees from MrBayes analysis.
-sup Specify an input “target” tree file. That means, support values are first extracted from the trees passed via -t, and then mapped onto the target tree. Resulting tree with assigned support values is written to .suptree file. This option is useful to map and compare support values from different approaches onto a single tree.
-suptag Specify name of a node in -sup target tree. The corresponding node of .suptree will then be assigned with IDs of trees where this node appears. Special option -suptag ALL will assign such IDs for all nodes of the target tree.

Computing Robinson-Foulds distance - IQ-TREE provides a fast implementation of Robinson-Foulds (RF) distance computation between trees:
Usage and meaning

-t Specify a file containing a set of trees.
-rf Specify a second set of trees. IQ-TREE computes all pairwise RF distances between two tree sets passed via -t and -rf
-rf_all Compute all-to-all RF distances between all trees passed via -t
-rf_adj Compute RF distances between adjacent trees passed via -t

Example usages:

•Compute the pairwise RF distances between 2 sets of trees:
  iqtree -rf tree_set1 tree_set2

•Compute the all-to-all RF distances within a set of trees:
  iqtree -rf_all tree_set

Generating random trees - IQ-TREE provides several random tree generation models:
Usage and meaning

-r Specify number of taxa. IQ-TREE will create a random tree under Yule-Harding model with specified number of taxa
-ru Like -r, but a random tree is created under uniform model.
-rcat Like -r, but a random caterpillar tree is created.
-rbal Like -r, but a random balanced tree is created.
-rcsg Like -r, bur a random circular split network is created.
-rlen Specify three numbers: minimum, mean and maximum branch lengths of the random tree. DEFAULT: -rlen 0.001 0.1 0.999

Example usages:

•Generate a 100-taxon random tree into the file 100.tree under the Yule Harding model:
  iqtree -r 100 100.tree

•Generate 100-taxon random tree with mean branch lengths of 0.2 and branch lengths are between 0.05 and 0.3:
  iqtree -r 100 -rlen 0.05 0.2 0.3 100.tree

•Generate a random tree under uniform model:
  iqtree -ru 100 100.tree

•Generate a random tree where taxon labels follow an alignment:
  iqtree -s example.phy -r 17 example.random.tree

Note that, you still need to specify the -r option being equal to the number of taxa in the alignment.

Miscellaneous options

Usage and meaning

-wt Turn on writing all locally optimal trees into .treels file. DEFAULT: OFF
-fixbr Turn on fixing branch lengths of tree passed via -t or -te. This is useful to evaluate the log-likelihood of an input tree with fixed tolopogy and branch lengths. DEFAULT: OFF
-wsl Turn on writing site log-likelihoods to .sitelh file in TREE-PUZZLE format. Such file can then be passed on to CONSEL for further tree tests. DEFAULT: OFF
-wslg Turn on writing site log-likelihoods per rate category. DEFAULT: OFF
-fconst Specify a list of comma-separated integer numbers. The number of entries should be equal to the number of states in the model (e.g. 4 for DNA and 20 for protein). IQ-TREE will then add a number of constant sites accordingly. For example, -fconst 10,20,15,40 will add 10 constant sites of all A, 20 constant sites of all C, 15 constant sites of all G and 40 constant sites of all T into the alignment.





	-->
<!-- new rules from Wayne
Rules for running IQ-TREE on Comet via the CIPRES gateway

The runs use varying numbers of cores within a single node of Comet depending upon the
specified options and the data set.

- Check whether the -mtree option is specified.

- Ask the user for the number of partitions in his or her data set.

- Specify the Slurm partition, CIPRES threads, and value for -nt according to the
folowing table.

            Data       Slurm     CIPRES    -nt
 -mtree  partitions  partition  threads   value

   no       <=12      shared       12      AUTO
   no        >12      compute      24      AUTO
   yes      <=12      shared       12       12
   yes       >12      compute      24       24

- For example, for a run without -mtree and with no more than 12 partitions,
include the following in the run script.

  #SBATCH -p shared
  #SBATCH -*qos=cipres
  #SBATCH -N 1
  #SBATCH -*ntasks-per-node=12
  ...
  export MODULEPATH=/projects/ps-ngbt/opt/comet/modules:$MODULEPATH
  module load iqtree/1.6.2
  iqtree -nt AUTO ...
 -->
 <!-- #####################################################################################################  -->
 <!-- revised on 3/9/2018
 [pfeiffer@comet-ln2 rules]$ cat IQ-TREE-1.6.2.Comet.rules
Rules for running IQ-TREE 1.6.2 on Comet via the CIPRES gateway

The runs use varying numbers of cores within a single node of Comet depending upon
the data set.

- Ask the user whether his or her data set has more than 12 partitions.
- Ask the user whether his or her analysis needs more than 60 GB of memory.
- Specify the Slurm partition, CIPRES threads, and value for -nt according to the
following table.

 Data partitions      Slurm     CIPRES    -nt    Additional
  & memory size     partition  threads   value     option

  <=12 & <=60 GB      shared      12      AUTO
   >12 or >60 GB     compute      24      AUTO    -mem 116G

- Thus, for a run with no more than 12 partitions and that requires no more than
60 GB of memory, include the following in the run script.

  #SBATCH -p shared
  #SBATCH -*qos=cipres
  #SBATCH -N 1
  #SBATCH -*ntasks-per-node=12
  ...
  export MODULEPATH=/projects/ps-ngbt/opt/comet/modules:$MODULEPATH
  module load iqtree/1.6.2
  iqtree -nt AUTO ...

- On the other hand, for a run with more than 12 partitions or that requires more
than 60 GB of memory, include the following in the run script.

  #SBATCH -p compute
  #SBATCH -*qos=cipres
  #SBATCH -N 1
  #SBATCH -*ntasks-per-node=24
  ...
  export MODULEPATH=/projects/ps-ngbt/opt/comet/modules:$MODULEPATH
  module load iqtree/1.6.2
  iqtree -nt AUTO -mem 116G ...
  -->
 <!-- ####################################################################################################### -->
 <!-- new rules 4_16_2018
Rules for running IQ-TREE 1.6.2 on Comet via the CIPRES gateway

The runs use varying numbers of cores within a single node of Comet depending upon
the data set.

- Ask the user whether his or her data set has more than 12 partitions.

- Check whether the -mem option is specified and, if so, what its value is in GB.
  If the value is >120, the job will not run.

- Specify the Slurm partition, CIPRES threads, and value for -nt according to
  the following table.

    Data       -mem       Slurm     CIPRES    -nt
 partitions   option    partition  threads   value

(note here that currently 1.6.2 does not support -mem for more than 1 partition)

    1           Blank     shared       12      AUTO
               or  <=60

    2<=12       blank     shared       12      AUTO

     >12        blank     compute      24      AUTO

     1          >60 &     compute      24      AUTO
               <=120

- For a run in the first row of the table, include the following in the run script.

  #SBATCH -p shared
  #SBATCH -N 1
  #SBATCH -*ntasks-per-node=12
  ...
  export MODULEPATH=/projects/ps-ngbt/opt/comet/modules:$MODULEPATH
  module load iqtree/1.6.2
  iqtree -nt AUTO ...

- For a run in the second or third row of the table, include the following in the
  run script.

  #SBATCH -p compute
  #SBATCH -N 1
  #SBATCH -*ntasks-per-node=24
  ...
  export MODULEPATH=/projects/ps-ngbt/opt/comet/modules:$MODULEPATH
  module load iqtree/1.6.2
  iqtree -nt AUTO ...
*/
type IQTree struct {
	XMLName xml.Name `xml:"pise"`
	Text    string   `xml:",chardata"`
	Head    struct {
		Text        string `xml:",chardata"`
		Title       string `xml:"title"`
		Version     string `xml:"version"`
		Description string `xml:"description"`
		Authors     string `xml:"authors"`
		Reference   string `xml:"reference"`
		Category    string `xml:"category"`
	} `xml:"head"`
	Command    string `xml:"command"`
	Parameters struct {
		Text      string `xml:",chardata"`
		Parameter []struct {
			Text        string `xml:",chardata"`
			Ismandatory string `xml:"ismandatory,attr"`
			Ishidden    string `xml:"ishidden,attr"`
			Type        string `xml:"type,attr"`
			Issimple    string `xml:"issimple,attr"`
			Isinput     string `xml:"isinput,attr"`
			Name        string `xml:"name"`
			Attributes  struct {
				Text    string `xml:",chardata"`
				Precond struct {
					Text     string `xml:",chardata"`
					Language string `xml:"language"`
					Code     string `xml:"code"`
				} `xml:"precond"`
				Format struct {
					Text     string `xml:",chardata"`
					Language string `xml:"language"`
					Code     string `xml:"code"`
				} `xml:"format"`
				Group     string `xml:"group"`
				Paramfile string `xml:"paramfile"`
				Prompt    string `xml:"prompt"`
				Filenames string `xml:"filenames"`
				Vdef      struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value"`
				} `xml:"vdef"`
				Ctrls struct {
					Text string `xml:",chardata"`
					Ctrl []struct {
						Text     string `xml:",chardata"`
						Message  string `xml:"message"`
						Language string `xml:"language"`
						Code     string `xml:"code"`
					} `xml:"ctrl"`
				} `xml:"ctrls"`
				Warns struct {
					Text string `xml:",chardata"`
					Warn []struct {
						Text     string `xml:",chardata"`
						Message  string `xml:"message"`
						Language string `xml:"language"`
						Code     string `xml:"code"`
					} `xml:"warn"`
				} `xml:"warns"`
				Comment struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value"`
				} `xml:"comment"`
				Vlist struct {
					Text  string   `xml:",chardata"`
					Value []string `xml:"value"`
					Label []string `xml:"label"`
				} `xml:"vlist"`
			} `xml:"attributes"`
			Paragraph struct {
				Text       string `xml:",chardata"`
				Name       string `xml:"name"`
				Prompt     string `xml:"prompt"`
				Parameters struct {
					Text      string `xml:",chardata"`
					Parameter []struct {
						Text       string `xml:",chardata"`
						Issimple   string `xml:"issimple,attr"`
						Type       string `xml:"type,attr"`
						Ishidden   string `xml:"ishidden,attr"`
						Name       string `xml:"name"`
						Attributes struct {
							Text   string `xml:",chardata"`
							Prompt string `xml:"prompt"`
							Vlist  struct {
								Text  string   `xml:",chardata"`
								Value []string `xml:"value"`
								Label []string `xml:"label"`
							} `xml:"vlist"`
							Format struct {
								Text     string `xml:",chardata"`
								Language string `xml:"language"`
								Code     string `xml:"code"`
							} `xml:"format"`
							Group   string `xml:"group"`
							Precond struct {
								Text     string `xml:",chardata"`
								Language string `xml:"language"`
								Code     string `xml:"code"`
							} `xml:"precond"`
							Comment struct {
								Text  string `xml:",chardata"`
								Value string `xml:"value"`
							} `xml:"comment"`
							Ctrls struct {
								Text string `xml:",chardata"`
								Ctrl []struct {
									Text     string `xml:",chardata"`
									Message  string `xml:"message"`
									Language string `xml:"language"`
									Code     string `xml:"code"`
								} `xml:"ctrl"`
							} `xml:"ctrls"`
							Filenames string `xml:"filenames"`
							Vdef      struct {
								Text  string `xml:",chardata"`
								Value string `xml:"value"`
							} `xml:"vdef"`
						} `xml:"attributes"`
					} `xml:"parameter"`
				} `xml:"parameters"`
			} `xml:"paragraph"`
		} `xml:"parameter"`
	} `xml:"parameters"`
}
