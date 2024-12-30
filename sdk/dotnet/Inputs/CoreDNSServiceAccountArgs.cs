// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesCoreDNS.Inputs
{

    public sealed class CoreDNSServiceAccountArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// If true, create &amp; use serviceAccount.
        /// </summary>
        [Input("create")]
        public Input<bool>? Create { get; set; }

        /// <summary>
        /// The name of the ServiceAccount to use. If not set and create is true, a name is generated using the fullname template
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public CoreDNSServiceAccountArgs()
        {
        }
        public static new CoreDNSServiceAccountArgs Empty => new CoreDNSServiceAccountArgs();
    }
}
