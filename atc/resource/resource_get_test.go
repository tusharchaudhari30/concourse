package resource_test

import (
	"context"
	"errors"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/resource"
	"github.com/concourse/concourse/atc/runtime"
	"github.com/concourse/concourse/atc/runtime/runtimefakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Resource Get", func() {
	var (
		ctx    context.Context
		someProcessSpec runtime.ProcessSpec
		fakeRunnable runtimefakes.FakeRunner

		getVersionResult runtime.VersionResult

		source  atc.Source
		params  atc.Params
		version atc.Version

		resource resource.Resource

		getErr error

		// OLD STUFF

		//inScriptStdout     string
		//inScriptStderr     string
		//inScriptExitStatus int
		//runInError         error
		//attachInError      error
		//
		//inScriptProcess *gardenfakes.FakeProcess
		//
		//stdoutBuf *gbytes.Buffer
		//stderrBuf *gbytes.Buffer
		//
		//fakeVolume *workerfakes.FakeVolume
		//
		//ctx    context.Context
		//cancel func()
		//
		//getErr error
	)

	BeforeEach(func() {
		ctx = context.Background()

		source = atc.Source{"some": "source"}
		version = atc.Version{"some": "version"}
		params = atc.Params{"some": "params"}

		someProcessSpec.Path = "some/fake/path"
		someProcessSpec.Args = []string{"first-arg", "some-other-arg"}
		someProcessSpec.StderrWriter = gbytes.NewBuffer()

		resource = resourceFactory.NewResource(source, params, version)

	})

	JustBeforeEach(func() {
		getVersionResult, getErr = resource.Get(ctx, someProcessSpec, &fakeRunnable)
	})

	Context("when Runnable -> RunScript succeeds", func() {
		BeforeEach(func(){
			fakeRunnable.RunScriptReturns(nil)
		})

		It("Invokes Runnable -> RunScript with the correct arguments", func(){
			actualCtx, actualSpecPath, actualSpecArgs,
				actualInputArgs, actualVersionResultRef, actualSpecStdErrWriter,
				actualRecoverableBool :=  fakeRunnable.RunScriptArgsForCall(0)

				signature, err := resource.Signature()
				Expect(err).ToNot(HaveOccurred())

				Expect(actualCtx).To(Equal(ctx))
				Expect(actualSpecPath).To(Equal(someProcessSpec.Path))
				Expect(actualSpecArgs).To(Equal(someProcessSpec.Args))
				Expect(actualInputArgs).To(Equal(signature))
				Expect(actualVersionResultRef).To(Equal(&getVersionResult))
				Expect(actualSpecStdErrWriter).To(Equal(someProcessSpec.StderrWriter))
				Expect(actualRecoverableBool).To(BeTrue())
		})

		It("doesnt return an error", func(){
			Expect(getErr).To(BeNil())
		})
	})

	Context("when Runnable -> RunScript returns an error", func() {
		var disasterErr = errors.New("there was an issue")
		BeforeEach(func(){
			fakeRunnable.RunScriptReturns(disasterErr)
		})
		It("returns the error", func() {
			Expect(getErr).To(Equal(disasterErr))
		})
	})
	//itCanStreamOut := func() {
	//	Describe("streaming bits out", func() {
	//		Context("when streaming out succeeds", func() {
	//			BeforeEach(func() {
	//				fakeVolume.StreamOutStub = func(ctx context.Context, path string) (io.ReadCloser, error) {
	//					streamOut := new(bytes.Buffer)
	//
	//					if path == "some/subdir" {
	//						streamOut.WriteString("sup")
	//					}
	//
	//					return ioutil.NopCloser(streamOut), nil
	//				}
	//			})
	//
	//			It("returns the output stream of the resource directory", func() {
	//				inStream, err := versionedSource.StreamOut(context.TODO(), "some/subdir")
	//				Expect(err).NotTo(HaveOccurred())
	//
	//				contents, err := ioutil.ReadAll(inStream)
	//				Expect(err).NotTo(HaveOccurred())
	//				Expect(string(contents)).To(Equal("sup"))
	//			})
	//		})
	//
	//		Context("when streaming out fails", func() {
	//			disaster := errors.New("oh no!")
	//
	//			BeforeEach(func() {
	//				fakeVolume.StreamOutReturns(nil, disaster)
	//			})
	//
	//			It("returns the error", func() {
	//				_, err := versionedSource.StreamOut(context.TODO(), "some/subdir")
	//				Expect(err.Error()).To(Equal("oh no!"))
	//			})
	//		})
	//	})
	//}
	//
	//Describe("running", func() {
	//
	//
	//	Context("when a result is already present on the container", func() {
	//		BeforeEach(func() {
	//			fakeContainer.PropertiesReturns(garden.Properties{"concourse:resource-result": `{
	//				"version": {"some": "new-version"},
	//				"metadata": [
	//					{"name": "a", "value":"a-value"},
	//					{"name": "b","value": "b-value"}
	//				]
	//			}`}, nil)
	//		})
	//
	//		It("exits successfully", func() {
	//			Expect(getErr).NotTo(HaveOccurred())
	//		})
	//
	//		It("does not run or attach to anything", func() {
	//			Expect(fakeContainer.RunCallCount()).To(BeZero())
	//			Expect(fakeContainer.AttachCallCount()).To(BeZero())
	//		})
	//
	//		It("can be accessed on the versioned source", func() {
	//			Expect(versionedSource.Version()).To(Equal(atc.Version{"some": "new-version"}))
	//			Expect(versionedSource.Metadata()).To(Equal([]atc.MetadataField{
	//				{Name: "a", Value: "a-value"},
	//				{Name: "b", Value: "b-value"},
	//			}))
	//		})
	//	})
	//
	//	Context("when /in has already been spawned", func() {
	//		BeforeEach(func() {
	//			fakeContainer.PropertiesReturns(nil, nil)
	//		})
	//
	//		It("reattaches to it", func() {
	//			Expect(fakeContainer.AttachCallCount()).To(Equal(1))
	//
	//			_, pid, io := fakeContainer.AttachArgsForCall(0)
	//			Expect(pid).To(Equal(resource.ResourceProcessID))
	//
	//			// send request on stdin in case process hasn't read it yet
	//			request, err := ioutil.ReadAll(io.Stdin)
	//			Expect(err).NotTo(HaveOccurred())
	//
	//			Expect(request).To(MatchJSON(`{
	//				"source": {"some":"source"},
	//				"params": {"some":"params"},
	//				"version": {"some":"version"}
	//			}`))
	//		})
	//
	//		It("does not run an additional process", func() {
	//			Expect(fakeContainer.RunCallCount()).To(BeZero())
	//		})
	//
	//		Context("when /opt/resource/in prints the response", func() {
	//			BeforeEach(func() {
	//				inScriptStdout = `{
	//				"version": {"some": "new-version"},
	//				"metadata": [
	//					{"name": "a", "value":"a-value"},
	//					{"name": "b","value": "b-value"}
	//				]
	//			}`
	//			})
	//
	//			It("can be accessed on the versioned source", func() {
	//				Expect(versionedSource.Version()).To(Equal(atc.Version{"some": "new-version"}))
	//				Expect(versionedSource.Metadata()).To(Equal([]atc.MetadataField{
	//					{Name: "a", Value: "a-value"},
	//					{Name: "b", Value: "b-value"},
	//				}))
	//
	//			})
	//
	//			It("saves it as a property on the container", func() {
	//				Expect(fakeContainer.SetPropertyCallCount()).To(Equal(1))
	//
	//				name, value := fakeContainer.SetPropertyArgsForCall(0)
	//				Expect(name).To(Equal("concourse:resource-result"))
	//				Expect(value).To(Equal(inScriptStdout))
	//			})
	//		})
	//
	//		Context("when /in outputs to stderr", func() {
	//			BeforeEach(func() {
	//				inScriptStderr = "some stderr data"
	//			})
	//
	//			It("emits it to the log sink", func() {
	//				Expect(stderrBuf).To(gbytes.Say("some stderr data"))
	//			})
	//		})
	//
	//		Context("when attaching to the process fails", func() {
	//			disaster := errors.New("oh no!")
	//
	//			BeforeEach(func() {
	//				attachInError = disaster
	//			})
	//
	//			Context("and run succeeds", func() {
	//				It("succeeds", func() {
	//					Expect(getErr).ToNot(HaveOccurred())
	//				})
	//			})
	//
	//			Context("and run subsequently fails", func() {
	//				BeforeEach(func() {
	//					runInError = disaster
	//				})
	//
	//				It("errors", func() {
	//					Expect(getErr).To(HaveOccurred())
	//					Expect(getErr).To(Equal(disaster))
	//				})
	//			})
	//		})
	//
	//		Context("when the process exits nonzero", func() {
	//			BeforeEach(func() {
	//				inScriptExitStatus = 9
	//			})
	//
	//			It("returns an err containing stdout/stderr of the process", func() {
	//				Expect(getErr).To(HaveOccurred())
	//				Expect(getErr.Error()).To(ContainSubstring("exit status 9"))
	//			})
	//		})
	//
	//		itCanStreamOut()
	//	})
	//
	//	Context("when /in has not yet been spawned", func() {
	//		BeforeEach(func() {
	//			fakeContainer.PropertiesReturns(nil, nil)
	//			attachInError = errors.New("not found")
	//		})
	//
	//		It("specifies the process id in the process spec", func() {
	//			Expect(fakeContainer.RunCallCount()).To(Equal(1))
	//
	//			_, spec, _ := fakeContainer.RunArgsForCall(0)
	//			Expect(spec.ID).To(Equal(resource.ResourceProcessID))
	//		})
	//
	//		It("uses the same working directory for all actions", func() {
	//			err := versionedSource.StreamIn(context.TODO(), "a/path", &bytes.Buffer{})
	//			Expect(err).NotTo(HaveOccurred())
	//
	//			Expect(fakeVolume.StreamInCallCount()).To(Equal(1))
	//			_, destPath, _ := fakeVolume.StreamInArgsForCall(0)
	//
	//			_, err = versionedSource.StreamOut(context.TODO(), "a/path")
	//			Expect(err).NotTo(HaveOccurred())
	//
	//			Expect(fakeVolume.StreamOutCallCount()).To(Equal(1))
	//			_, path := fakeVolume.StreamOutArgsForCall(0)
	//			Expect(path).To(Equal("a/path"))
	//
	//			Expect(fakeContainer.RunCallCount()).To(Equal(1))
	//
	//			Expect(destPath).To(Equal("/tmp/build/get/a/path"))
	//		})
	//
	//		It("runs /opt/resource/in <destination> with the request on stdin", func() {
	//			Expect(fakeContainer.RunCallCount()).To(Equal(1))
	//
	//			_, spec, io := fakeContainer.RunArgsForCall(0)
	//			Expect(spec.Path).To(Equal("/opt/resource/in"))
	//			Expect(spec.Args).To(ConsistOf("/tmp/build/get"))
	//
	//			request, err := ioutil.ReadAll(io.Stdin)
	//			Expect(err).NotTo(HaveOccurred())
	//
	//			Expect(request).To(MatchJSON(`{
	//			"source": {"some":"source"},
	//			"params": {"some":"params"},
	//			"version": {"some":"version"}
	//		}`))
	//		})
	//
	//		Context("when /opt/resource/in prints the response", func() {
	//			BeforeEach(func() {
	//				inScriptStdout = `{
	//				"version": {"some": "new-version"},
	//				"metadata": [
	//					{"name": "a", "value":"a-value"},
	//					{"name": "b","value": "b-value"}
	//				]
	//			}`
	//			})
	//
	//			It("can be accessed on the versioned source", func() {
	//				Expect(versionedSource.Version()).To(Equal(atc.Version{"some": "new-version"}))
	//				Expect(versionedSource.Metadata()).To(Equal([]atc.MetadataField{
	//					{Name: "a", Value: "a-value"},
	//					{Name: "b", Value: "b-value"},
	//				}))
	//
	//			})
	//
	//			It("saves it as a property on the container", func() {
	//				Expect(fakeContainer.SetPropertyCallCount()).To(Equal(1))
	//
	//				name, value := fakeContainer.SetPropertyArgsForCall(0)
	//				Expect(name).To(Equal("concourse:resource-result"))
	//				Expect(value).To(Equal(inScriptStdout))
	//			})
	//		})
	//
	//		Context("when /in outputs to stderr", func() {
	//			BeforeEach(func() {
	//				inScriptStderr = "some stderr data"
	//			})
	//
	//			It("emits it to the log sink", func() {
	//				Expect(stderrBuf).To(gbytes.Say("some stderr data"))
	//			})
	//		})
	//
	//		Context("when running /opt/resource/in fails", func() {
	//			disaster := errors.New("oh no!")
	//
	//			BeforeEach(func() {
	//				runInError = disaster
	//			})
	//
	//			It("returns an err", func() {
	//				Expect(getErr).To(HaveOccurred())
	//				Expect(getErr).To(Equal(disaster))
	//			})
	//		})
	//
	//		Context("when /opt/resource/in exits nonzero", func() {
	//			BeforeEach(func() {
	//				inScriptExitStatus = 9
	//			})
	//
	//			It("returns an err containing stdout/stderr of the process", func() {
	//				Expect(getErr).To(HaveOccurred())
	//				Expect(getErr.Error()).To(ContainSubstring("exit status 9"))
	//			})
	//		})
	//
	//		Context("when the output of /opt/resource/in is malformed", func() {
	//			BeforeEach(func() {
	//				inScriptStdout = "ß"
	//			})
	//
	//			It("returns an error", func() {
	//				Expect(getErr).To(HaveOccurred())
	//			})
	//
	//			It("returns original payload in error", func() {
	//				Expect(getErr.Error()).Should(ContainSubstring(inScriptStdout))
	//			})
	//		})
	//
	//		itCanStreamOut()
	//	})
	//})
	//
	//Context("when canceling the context", func() {
	//	var waited chan<- struct{}
	//	var done chan struct{}
	//
	//	BeforeEach(func() {
	//		fakeContainer.AttachReturns(nil, errors.New("not-found"))
	//		fakeContainer.RunReturns(inScriptProcess, nil)
	//		fakeContainer.PropertyReturns("", errors.New("nope"))
	//
	//		waiting := make(chan struct{})
	//		done = make(chan struct{})
	//		waited = waiting
	//
	//		inScriptProcess.WaitStub = func() (int, error) {
	//			// cause waiting to block so that it can be aborted
	//			<-waiting
	//			return 0, nil
	//		}
	//
	//		fakeContainer.StopStub = func(bool) error {
	//			close(waited)
	//			return nil
	//		}
	//
	//		go func() {
	//			versionedSource, getErr = resourceForContainer.Get(ctx, fakeVolume, ioConfig, source, params, version)
	//			close(done)
	//		}()
	//	})
	//
	//	It("stops the container", func() {
	//		cancel()
	//		<-done
	//		Expect(fakeContainer.StopCallCount()).To(Equal(1))
	//		isStopped := fakeContainer.StopArgsForCall(0)
	//		Expect(isStopped).To(BeFalse())
	//	})
	//
	//	It("doesn't send garden terminate signal to process", func() {
	//		cancel()
	//		<-done
	//		Expect(getErr).To(Equal(context.Canceled))
	//		Expect(inScriptProcess.SignalCallCount()).To(BeZero())
	//	})
	//
	//	Context("when container.stop returns an error", func() {
	//		var disaster error
	//
	//		BeforeEach(func() {
	//			disaster = errors.New("gotta get away")
	//
	//			fakeContainer.StopStub = func(bool) error {
	//				close(waited)
	//				return disaster
	//			}
	//		})
	//
	//		It("masks the error", func() {
	//			cancel()
	//			<-done
	//			Expect(getErr).To(Equal(context.Canceled))
	//		})
	//	})
	//})
})
