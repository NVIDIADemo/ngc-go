// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NVIDIADemo/ngc-go/internal/apijson"
	"github.com/NVIDIADemo/ngc-go/internal/requestconfig"
)

type PageNumberOrganizations[T any] struct {
	Organizations []T                         `json:"organizations"`
	JSON          pageNumberOrganizationsJSON `json:"-"`
	cfg           *requestconfig.RequestConfig
	res           *http.Response
}

// pageNumberOrganizationsJSON contains the JSON metadata for the struct
// [PageNumberOrganizations[T]]
type pageNumberOrganizationsJSON struct {
	Organizations apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PageNumberOrganizations[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageNumberOrganizationsJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PageNumberOrganizations[T]) GetNextPage() (res *PageNumberOrganizations[T], err error) {
	u := r.cfg.Request.URL
	currentPage, err := strconv.Atoi(u.Query().Get("page-number"))
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page-number", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PageNumberOrganizations[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PageNumberOrganizations[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PageNumberOrganizationsAutoPager[T any] struct {
	page *PageNumberOrganizations[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPageNumberOrganizationsAutoPager[T any](page *PageNumberOrganizations[T], err error) *PageNumberOrganizationsAutoPager[T] {
	return &PageNumberOrganizationsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageNumberOrganizationsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Organizations) == 0 {
		return false
	}
	if r.idx >= len(r.page.Organizations) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Organizations) == 0 {
			return false
		}
	}
	r.cur = r.page.Organizations[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageNumberOrganizationsAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageNumberOrganizationsAutoPager[T]) Err() error {
	return r.err
}

func (r *PageNumberOrganizationsAutoPager[T]) Index() int {
	return r.run
}

type PageNumberUsers[T any] struct {
	Users []T                 `json:"users"`
	JSON  pageNumberUsersJSON `json:"-"`
	cfg   *requestconfig.RequestConfig
	res   *http.Response
}

// pageNumberUsersJSON contains the JSON metadata for the struct
// [PageNumberUsers[T]]
type pageNumberUsersJSON struct {
	Users       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageNumberUsers[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageNumberUsersJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PageNumberUsers[T]) GetNextPage() (res *PageNumberUsers[T], err error) {
	u := r.cfg.Request.URL
	currentPage, err := strconv.Atoi(u.Query().Get("page-number"))
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page-number", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PageNumberUsers[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PageNumberUsers[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PageNumberUsersAutoPager[T any] struct {
	page *PageNumberUsers[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPageNumberUsersAutoPager[T any](page *PageNumberUsers[T], err error) *PageNumberUsersAutoPager[T] {
	return &PageNumberUsersAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageNumberUsersAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Users) == 0 {
		return false
	}
	if r.idx >= len(r.page.Users) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Users) == 0 {
			return false
		}
	}
	r.cur = r.page.Users[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageNumberUsersAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageNumberUsersAutoPager[T]) Err() error {
	return r.err
}

func (r *PageNumberUsersAutoPager[T]) Index() int {
	return r.run
}

type PageNumberTeams[T any] struct {
	Teams []T                 `json:"teams"`
	JSON  pageNumberTeamsJSON `json:"-"`
	cfg   *requestconfig.RequestConfig
	res   *http.Response
}

// pageNumberTeamsJSON contains the JSON metadata for the struct
// [PageNumberTeams[T]]
type pageNumberTeamsJSON struct {
	Teams       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageNumberTeams[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageNumberTeamsJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PageNumberTeams[T]) GetNextPage() (res *PageNumberTeams[T], err error) {
	u := r.cfg.Request.URL
	currentPage, err := strconv.Atoi(u.Query().Get("page-number"))
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page-number", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PageNumberTeams[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PageNumberTeams[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PageNumberTeamsAutoPager[T any] struct {
	page *PageNumberTeams[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPageNumberTeamsAutoPager[T any](page *PageNumberTeams[T], err error) *PageNumberTeamsAutoPager[T] {
	return &PageNumberTeamsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageNumberTeamsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Teams) == 0 {
		return false
	}
	if r.idx >= len(r.page.Teams) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Teams) == 0 {
			return false
		}
	}
	r.cur = r.page.Teams[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageNumberTeamsAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageNumberTeamsAutoPager[T]) Err() error {
	return r.err
}

func (r *PageNumberTeamsAutoPager[T]) Index() int {
	return r.run
}

type PageNumberInvitations[T any] struct {
	Invitations []T                       `json:"invitations"`
	JSON        pageNumberInvitationsJSON `json:"-"`
	cfg         *requestconfig.RequestConfig
	res         *http.Response
}

// pageNumberInvitationsJSON contains the JSON metadata for the struct
// [PageNumberInvitations[T]]
type pageNumberInvitationsJSON struct {
	Invitations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageNumberInvitations[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageNumberInvitationsJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PageNumberInvitations[T]) GetNextPage() (res *PageNumberInvitations[T], err error) {
	u := r.cfg.Request.URL
	currentPage, err := strconv.Atoi(u.Query().Get("page-number"))
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page-number", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PageNumberInvitations[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PageNumberInvitations[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PageNumberInvitationsAutoPager[T any] struct {
	page *PageNumberInvitations[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPageNumberInvitationsAutoPager[T any](page *PageNumberInvitations[T], err error) *PageNumberInvitationsAutoPager[T] {
	return &PageNumberInvitationsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageNumberInvitationsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Invitations) == 0 {
		return false
	}
	if r.idx >= len(r.page.Invitations) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Invitations) == 0 {
			return false
		}
	}
	r.cur = r.page.Invitations[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageNumberInvitationsAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageNumberInvitationsAutoPager[T]) Err() error {
	return r.err
}

func (r *PageNumberInvitationsAutoPager[T]) Index() int {
	return r.run
}
