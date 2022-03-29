# goLang-injectors
This project will guide yout to awareness of injection in almost every window API and process.
<h6>Credit: <a href="https://github.com/zaneGittins/go-inject">zaneGittins</a> & <a href="https://github.com/Ne0nd0g/go-shellcode"> Ne0nd0g </a>  </h6>
After combining these two repositries and adding some custom code, current tool/repository is created.
This repository will help you to inject 
<ol>
  <li>CreateFiber</li>
   <li>CreateProcess</li>
   <li>CreateRemoteThread</li>
   <li>CreateRemoteThreadNative</li>
   <li>CreateThread</li>
   <li>CreateThreadNative</li>
   <li>EarlyBird</li>
   <li>EnumerateLoadedModules</li>
   <li>EtwpCreateEtwThread</li>
   <li>NtQueueApcThreadEx-Local</li>
   <li>CreateFiber</li>
   <li>RtlCreateUserThread</li>
   <li>Syscall</li>
   <li>UuidFromString</li>
   <li>CreateProcessWithPipe</li>
  </ol>
  <h4>Things needed</h4>
   <ol>
  <li>Go Lang Installed</li>
  <li>Packeges of this repo installed by using go get ...</li>
  <li>An independent hex code generator (https://github.com/Jhangju/portable-executable-2-hex-code) </li>
  <li>Some basic go lang knowledge</ol>
  </lo>
  <h3>Usage</h3>
  <ol>
<li>At first place your payload, donut.exe(<a href="https://github.com/Jhangju/portable-executable-2-hex-code/blob/main/donut.exe">click here to download donut.exe</a>) and pe2hex.exe(<a href="https://github.com/Jhangju/portable-executable-2-hex-code/blob/main/pe2hex/pe2hex/bin/Debug/pe2hex.exe">click here to download pe2hex.exe</a>) in same folder.
<li>Start donut.exe -f {YouPayload.exe} //you will get payload.bin which is independent binary.</li>
<li>Start pe2hex.exe -h {payload.bin}  //You will get hex.txt which is an independent hex code of you exploit.</li>
<li>Now just update hex code in any of top go file</li>
<li>Run Go-Lang code with go run {anyfile.go}</li>
  </lo>
