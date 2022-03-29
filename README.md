<h5>Credit <a href="https://github.com/zaneGittins/go-inject">zaneGittins</a> & <a href="https://github.com/Ne0nd0g/go-shellcode/tree/master/cmd"> Ne0nd0g </a></h5>
<h5>Never scan these exploits on public scanners.</h5>
This tool is created after combining online go injectors and some custom code.<hr>

<h4>Injection can be performed in following</h4>
<ol>
  <li>CreateFiber </li>
  <li>CreateProcess </li>
  <li>CreateProcessWithPipe </li>
  <li>CreateRemoteThread </li>
  <li>CreateRemoteThreadNative </li>
  <li>CreateThread </li>
  <li>CreateThreadNative </li>
  <li>EarlyBird </li>
  <li>EnumerateLoadedModules </li>
  <li>EtwpCreateEtwThread </li>
  <li>NtQueueApcThreadEx-Local </li>
  <li>RtlCreateUserThread </li>
  <li>Syscall </li>
  <li>UuidFromString </li>
  </ol>
 <hr>
<h4>Stuff Needed</h4>
<ol>
  <li>An independent HEX Code. </li>
  <li>GO-LANG Installed in System. </li>
  <li>Some basic Go-Lang knowlodge</li>
  <li>Packeges of this repo installed // go get ....</li>
  </ol>
  <hr>
  <h3>Usage</h3>
<ol>
  <li>At first put donut.exe(<a href="https://github.com/Jhangju/portable-executable-2-hex-code/blob/main/donut.exe">donut.exe</a>), pe2hex.exe(<a href="https://github.com/Jhangju/portable-executable-2-hex-code/blob/main/pe2hex/pe2hex/bin/Debug/pe2hex.exe">pe2hex.exe</a>) and your payload in same folder. </li>
  <li>Start donut.exe -f {{your_payload}} // It will create independet payload.bin binary.</li>
  <li>Start pe2hex.exe -h {{payload.bin}} //It will create an independent hex.txt which is hex code.</li>
  <li>Use this HEX code in go lang file update already existing hex code.</li>
  <li>Run the file. Using Go run {file.go}</li>
  </ol>


## References
* https://blog.sunggwanchoi.com/eng-uuid-shellcode-execution/
* https://github.com/Adepts-Of-0xCC/VBA-macro-experiments/blob/main/EDRHookDetector.vba
* https://github.com/brimstone/go-shellcode
* https://github.com/sysdream/hershell
* https://github.com/yoda66/MalwareDevTalk
* https://labs.jumpsec.com/2019/06/20/bypassing-antivirus-with-golang-gopher-it/
* https://medium.com/@justen.walker/breaking-all-the-rules-using-go-to-call-windows-api-2cbfd8c79724
* https://posts.specterops.io/adventures-in-dynamic-evasion-1fe0bac57aa
* https://research.nccgroup.com/2021/01/23/rift-analysing-a-lazarus-shellcode-execution-method/
* https://www.ired.team/offensive-security/code-injection-process-injection/apc-queue-code-injection
* https://github.com/abdullah2993/go-runpe
