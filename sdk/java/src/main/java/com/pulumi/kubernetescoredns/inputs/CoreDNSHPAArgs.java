// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetescoredns.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetes.autoscaling.v2beta2.inputs.MetricSpecArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CoreDNSHPAArgs extends com.pulumi.resources.ResourceArgs {

    public static final CoreDNSHPAArgs Empty = new CoreDNSHPAArgs();

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="maxReplicas")
    private @Nullable Output<Integer> maxReplicas;

    public Optional<Output<Integer>> maxReplicas() {
        return Optional.ofNullable(this.maxReplicas);
    }

    @Import(name="metrics")
    private @Nullable Output<MetricSpecArgs> metrics;

    public Optional<Output<MetricSpecArgs>> metrics() {
        return Optional.ofNullable(this.metrics);
    }

    @Import(name="minReplicas")
    private @Nullable Output<Integer> minReplicas;

    public Optional<Output<Integer>> minReplicas() {
        return Optional.ofNullable(this.minReplicas);
    }

    private CoreDNSHPAArgs() {}

    private CoreDNSHPAArgs(CoreDNSHPAArgs $) {
        this.enabled = $.enabled;
        this.maxReplicas = $.maxReplicas;
        this.metrics = $.metrics;
        this.minReplicas = $.minReplicas;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CoreDNSHPAArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CoreDNSHPAArgs $;

        public Builder() {
            $ = new CoreDNSHPAArgs();
        }

        public Builder(CoreDNSHPAArgs defaults) {
            $ = new CoreDNSHPAArgs(Objects.requireNonNull(defaults));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public Builder maxReplicas(@Nullable Output<Integer> maxReplicas) {
            $.maxReplicas = maxReplicas;
            return this;
        }

        public Builder maxReplicas(Integer maxReplicas) {
            return maxReplicas(Output.of(maxReplicas));
        }

        public Builder metrics(@Nullable Output<MetricSpecArgs> metrics) {
            $.metrics = metrics;
            return this;
        }

        public Builder metrics(MetricSpecArgs metrics) {
            return metrics(Output.of(metrics));
        }

        public Builder minReplicas(@Nullable Output<Integer> minReplicas) {
            $.minReplicas = minReplicas;
            return this;
        }

        public Builder minReplicas(Integer minReplicas) {
            return minReplicas(Output.of(minReplicas));
        }

        public CoreDNSHPAArgs build() {
            return $;
        }
    }

}
