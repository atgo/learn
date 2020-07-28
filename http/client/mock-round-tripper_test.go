type mockRoundTrip struct {
	callback func(*http.Request) (*http.Response, error)
}

func (m mockRoundTrip) RoundTrip(request *http.Request) (*http.Response, error) {
	return m.callback(request)
}

func Test_Request(t *testing.T) {
		client, err := New("localhost:9000", "minioadmin", "minioadmin", false)
		assert.NoError(t, err)

		policy := NewPostPolicy()
		_ = policy.SetBucket("myBucket")
		_ = policy.SetKey("myObject")
		_ = policy.SetExpires(time.Now().Add(5 * time.Minute))

		url, formData, err := client.PresignedPostPolicy(context.Background(), policy)

		if nil != err {
			t.Errorf("failed executing client.PresignedPostPolicy: %s", err)
		}

		if url.String() != "http://localhost:9000/myBucket/" {
			t.Errorf("unexpected URL: %s", url.String())
		}

		if formData["bucket"] != "myBucket" {
			t.Errorf("unexpected bucket: %s", formData["bucket"])
		}

		if formData["key"] != "myObject" {
			t.Errorf("unexpected key: %s", formData["key"])
		}

		if _, ok := formData["x-amz-signature"]; !ok {
			t.Errorf("missing signagure")
		}
}
