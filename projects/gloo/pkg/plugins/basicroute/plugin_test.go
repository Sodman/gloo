package basicroute_test

import (
	"time"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/protocol_upgrade"

	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/ptypes/wrappers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/pkg/utils/gogoutils"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/retries"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	. "github.com/solo-io/gloo/projects/gloo/pkg/plugins/basicroute"
)

var _ = Describe("prefix rewrite", func() {
	It("works", func() {
		p := NewPlugin()
		routeAction := &envoy_config_route_v3.RouteAction{
			PrefixRewrite: "/",
		}
		out := &envoy_config_route_v3.Route{
			Action: &envoy_config_route_v3.Route_Route{
				Route: routeAction,
			},
		}
		err := p.ProcessRoute(plugins.RouteParams{}, &v1.Route{
			Options: &v1.RouteOptions{
				PrefixRewrite: &types.StringValue{Value: "/foo"},
			},
		}, out)
		Expect(err).NotTo(HaveOccurred())
		Expect(routeAction.PrefixRewrite).To(Equal("/foo"))
	})

	It("distinguishes between empty string and nil", func() {
		p := NewPlugin()
		routeAction := &envoy_config_route_v3.RouteAction{
			PrefixRewrite: "/",
		}
		out := &envoy_config_route_v3.Route{
			Action: &envoy_config_route_v3.Route_Route{
				Route: routeAction,
			},
		}

		// should be no-op
		err := p.ProcessRoute(plugins.RouteParams{}, &v1.Route{
			Options: &v1.RouteOptions{},
		}, out)
		Expect(err).NotTo(HaveOccurred())
		Expect(routeAction.PrefixRewrite).To(Equal("/"))

		// should rewrite prefix rewrite
		err = p.ProcessRoute(plugins.RouteParams{}, &v1.Route{
			Options: &v1.RouteOptions{
				PrefixRewrite: &types.StringValue{Value: ""},
			},
		}, out)
		Expect(err).NotTo(HaveOccurred())
		Expect(routeAction.PrefixRewrite).To(BeEmpty())
	})
})

var _ = Describe("timeout", func() {
	It("works", func() {
		t := time.Minute
		p := NewPlugin()
		routeAction := &envoy_config_route_v3.RouteAction{}
		out := &envoy_config_route_v3.Route{
			Action: &envoy_config_route_v3.Route_Route{
				Route: routeAction,
			},
		}
		err := p.ProcessRoute(plugins.RouteParams{}, &v1.Route{
			Options: &v1.RouteOptions{
				Timeout: &t,
			},
		}, out)
		Expect(err).NotTo(HaveOccurred())
		Expect(routeAction.Timeout).NotTo(BeNil())
		Expect(routeAction.Timeout).To(Equal(gogoutils.DurationStdToProto(&t)))
	})
})

var _ = Describe("retries", func() {

	var (
		plugin              *Plugin
		retryPolicy         *retries.RetryPolicy
		expectedRetryPolicy *envoy_config_route_v3.RetryPolicy
	)
	BeforeEach(func() {
		t := time.Minute
		retryPolicy = &retries.RetryPolicy{
			RetryOn:       "if at first you don't succeed",
			NumRetries:    5,
			PerTryTimeout: &t,
		}
		expectedRetryPolicy = &envoy_config_route_v3.RetryPolicy{
			RetryOn: "if at first you don't succeed",
			NumRetries: &wrappers.UInt32Value{
				Value: 5,
			},
			PerTryTimeout: gogoutils.DurationStdToProto(&t),
		}

		plugin = NewPlugin()
	})

	It("works", func() {
		routeAction := &envoy_config_route_v3.RouteAction{}
		out := &envoy_config_route_v3.Route{
			Action: &envoy_config_route_v3.Route_Route{
				Route: routeAction,
			},
		}
		err := plugin.ProcessRoute(plugins.RouteParams{}, &v1.Route{
			Options: &v1.RouteOptions{
				Retries: retryPolicy,
			},
		}, out)
		Expect(err).NotTo(HaveOccurred())
		Expect(routeAction.RetryPolicy).To(Equal(expectedRetryPolicy))
	})
	It("works on vhost", func() {
		out := &envoy_config_route_v3.VirtualHost{}
		err := plugin.ProcessVirtualHost(plugins.VirtualHostParams{}, &v1.VirtualHost{
			Options: &v1.VirtualHostOptions{
				Retries: retryPolicy,
			},
		}, out)
		Expect(err).NotTo(HaveOccurred())
		Expect(out.RetryPolicy).To(Equal(expectedRetryPolicy))
	})
})

var _ = Describe("host rewrite", func() {
	It("rewrites using provided string", func() {

		p := NewPlugin()
		routeAction := &envoy_config_route_v3.RouteAction{
			HostRewriteSpecifier: &envoy_config_route_v3.RouteAction_HostRewriteLiteral{HostRewriteLiteral: "/"},
		}
		out := &envoy_config_route_v3.Route{
			Action: &envoy_config_route_v3.Route_Route{
				Route: routeAction,
			},
		}
		err := p.ProcessRoute(plugins.RouteParams{}, &v1.Route{
			Options: &v1.RouteOptions{
				HostRewriteType: &v1.RouteOptions_HostRewrite{HostRewrite: "/foo"},
			},
		}, out)
		Expect(err).NotTo(HaveOccurred())
		Expect(routeAction.GetHostRewriteLiteral()).To(Equal("/foo"))
	})

	It("distinguishes between empty string and nil", func() {
		p := NewPlugin()
		routeAction := &envoy_config_route_v3.RouteAction{
			HostRewriteSpecifier: &envoy_config_route_v3.RouteAction_HostRewriteLiteral{HostRewriteLiteral: "/"},
		}
		out := &envoy_config_route_v3.Route{
			Action: &envoy_config_route_v3.Route_Route{
				Route: routeAction,
			},
		}

		// should be no-op
		err := p.ProcessRoute(plugins.RouteParams{}, &v1.Route{
			Options: &v1.RouteOptions{},
		}, out)
		Expect(err).NotTo(HaveOccurred())
		Expect(routeAction.GetHostRewriteLiteral()).To(Equal("/"))

		// should rewrite host rewrite
		err = p.ProcessRoute(plugins.RouteParams{}, &v1.Route{
			Options: &v1.RouteOptions{
				HostRewriteType: &v1.RouteOptions_HostRewrite{HostRewrite: ""},
			},
		}, out)
		Expect(err).NotTo(HaveOccurred())
		Expect(routeAction.GetHostRewriteLiteral()).To(BeEmpty())
	})

	It("sets auto_host_rewrite", func() {

		p := NewPlugin()
		routeAction := &envoy_config_route_v3.RouteAction{
			HostRewriteSpecifier: &envoy_config_route_v3.RouteAction_AutoHostRewrite{
				AutoHostRewrite: &wrappers.BoolValue{
					Value: false,
				},
			},
		}
		out := &envoy_config_route_v3.Route{
			Action: &envoy_config_route_v3.Route_Route{
				Route: routeAction,
			},
		}
		err := p.ProcessRoute(plugins.RouteParams{}, &v1.Route{
			Options: &v1.RouteOptions{
				HostRewriteType: &v1.RouteOptions_AutoHostRewrite{
					AutoHostRewrite: &types.BoolValue{
						Value: true,
					},
				},
			},
		}, out)
		Expect(err).NotTo(HaveOccurred())
		Expect(routeAction.GetAutoHostRewrite().GetValue()).To(Equal(true))
	})
})

var _ = Describe("upgrades", func() {
	It("works", func() {
		p := NewPlugin()

		routeAction := &envoy_config_route_v3.RouteAction{}

		out := &envoy_config_route_v3.Route{
			Action: &envoy_config_route_v3.Route_Route{
				Route: routeAction,
			},
		}

		err := p.ProcessRoute(plugins.RouteParams{}, &v1.Route{
			Options: &v1.RouteOptions{
				Upgrades: []*protocol_upgrade.ProtocolUpgradeConfig{
					{
						UpgradeType: &protocol_upgrade.ProtocolUpgradeConfig_Websocket{
							Websocket: &protocol_upgrade.ProtocolUpgradeConfig_ProtocolUpgradeSpec{
								Enabled: &types.BoolValue{Value: true},
							},
						},
					},
				},
			},
		}, out)

		Expect(err).NotTo(HaveOccurred())
		Expect(len(routeAction.GetUpgradeConfigs())).To(Equal(1))
		Expect(routeAction.GetUpgradeConfigs()[0].UpgradeType).To(Equal("websocket"))
		Expect(routeAction.GetUpgradeConfigs()[0].Enabled.Value).To(Equal(true))
	})
	It("fails on double config", func() {
		p := NewPlugin()

		routeAction := &envoy_config_route_v3.RouteAction{}

		out := &envoy_config_route_v3.Route{
			Action: &envoy_config_route_v3.Route_Route{
				Route: routeAction,
			},
		}

		err := p.ProcessRoute(plugins.RouteParams{}, &v1.Route{
			Options: &v1.RouteOptions{
				Upgrades: []*protocol_upgrade.ProtocolUpgradeConfig{
					{
						UpgradeType: &protocol_upgrade.ProtocolUpgradeConfig_Websocket{
							Websocket: &protocol_upgrade.ProtocolUpgradeConfig_ProtocolUpgradeSpec{
								Enabled: &types.BoolValue{Value: true},
							},
						},
					},
					{
						UpgradeType: &protocol_upgrade.ProtocolUpgradeConfig_Websocket{
							Websocket: &protocol_upgrade.ProtocolUpgradeConfig_ProtocolUpgradeSpec{
								Enabled: &types.BoolValue{Value: true},
							},
						},
					},
				},
			},
		}, out)

		Expect(err).To(MatchError(ContainSubstring("upgrade config websocket is not unique")))
	})
})
