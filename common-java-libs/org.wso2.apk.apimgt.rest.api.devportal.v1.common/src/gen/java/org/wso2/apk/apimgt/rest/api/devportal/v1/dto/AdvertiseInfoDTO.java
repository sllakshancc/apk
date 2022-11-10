package org.wso2.apk.apimgt.rest.api.devportal.v1.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import javax.validation.constraints.*;


import io.swagger.annotations.*;
import java.util.Objects;



public class AdvertiseInfoDTO   {
  
  private Boolean advertised;

  private String apiExternalProductionEndpoint;

  private String apiExternalSandboxEndpoint;

  private String originalDevPortalUrl;

  private String apiOwner;


public enum VendorEnum {

    WSO2(String.valueOf("WSO2")), AWS(String.valueOf("AWS"));


    private String value;

    VendorEnum(String v) {
        value = v;
    }

    public String value() {
        return value;
    }

    @Override
    public String toString() {
        return String.valueOf(value);
    }

    public static VendorEnum fromValue(String value) {
        for (VendorEnum b : VendorEnum.values()) {
            if (b.value.equals(value)) {
                return b;
            }
        }
        throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
}

  private VendorEnum vendor = VendorEnum.WSO2;


  /**
   **/
  public AdvertiseInfoDTO advertised(Boolean advertised) {
    this.advertised = advertised;
    return this;
  }

  
  @ApiModelProperty(example = "true", value = "")
  @JsonProperty("advertised")
  public Boolean getAdvertised() {
    return advertised;
  }
  public void setAdvertised(Boolean advertised) {
    this.advertised = advertised;
  }


  /**
   **/
  public AdvertiseInfoDTO apiExternalProductionEndpoint(String apiExternalProductionEndpoint) {
    this.apiExternalProductionEndpoint = apiExternalProductionEndpoint;
    return this;
  }

  
  @ApiModelProperty(example = "https://localhost:9443/devportal", value = "")
  @JsonProperty("apiExternalProductionEndpoint")
  public String getApiExternalProductionEndpoint() {
    return apiExternalProductionEndpoint;
  }
  public void setApiExternalProductionEndpoint(String apiExternalProductionEndpoint) {
    this.apiExternalProductionEndpoint = apiExternalProductionEndpoint;
  }


  /**
   **/
  public AdvertiseInfoDTO apiExternalSandboxEndpoint(String apiExternalSandboxEndpoint) {
    this.apiExternalSandboxEndpoint = apiExternalSandboxEndpoint;
    return this;
  }

  
  @ApiModelProperty(example = "https://localhost:9443/devportal", value = "")
  @JsonProperty("apiExternalSandboxEndpoint")
  public String getApiExternalSandboxEndpoint() {
    return apiExternalSandboxEndpoint;
  }
  public void setApiExternalSandboxEndpoint(String apiExternalSandboxEndpoint) {
    this.apiExternalSandboxEndpoint = apiExternalSandboxEndpoint;
  }


  /**
   **/
  public AdvertiseInfoDTO originalDevPortalUrl(String originalDevPortalUrl) {
    this.originalDevPortalUrl = originalDevPortalUrl;
    return this;
  }

  
  @ApiModelProperty(example = "https://localhost:9443/devportal", value = "")
  @JsonProperty("originalDevPortalUrl")
  public String getOriginalDevPortalUrl() {
    return originalDevPortalUrl;
  }
  public void setOriginalDevPortalUrl(String originalDevPortalUrl) {
    this.originalDevPortalUrl = originalDevPortalUrl;
  }


  /**
   **/
  public AdvertiseInfoDTO apiOwner(String apiOwner) {
    this.apiOwner = apiOwner;
    return this;
  }

  
  @ApiModelProperty(example = "admin", value = "")
  @JsonProperty("apiOwner")
  public String getApiOwner() {
    return apiOwner;
  }
  public void setApiOwner(String apiOwner) {
    this.apiOwner = apiOwner;
  }


  /**
   **/
  public AdvertiseInfoDTO vendor(VendorEnum vendor) {
    this.vendor = vendor;
    return this;
  }

  
  @ApiModelProperty(value = "")
  @JsonProperty("vendor")
  public VendorEnum getVendor() {
    return vendor;
  }
  public void setVendor(VendorEnum vendor) {
    this.vendor = vendor;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AdvertiseInfoDTO advertiseInfo = (AdvertiseInfoDTO) o;
    return Objects.equals(advertised, advertiseInfo.advertised) &&
        Objects.equals(apiExternalProductionEndpoint, advertiseInfo.apiExternalProductionEndpoint) &&
        Objects.equals(apiExternalSandboxEndpoint, advertiseInfo.apiExternalSandboxEndpoint) &&
        Objects.equals(originalDevPortalUrl, advertiseInfo.originalDevPortalUrl) &&
        Objects.equals(apiOwner, advertiseInfo.apiOwner) &&
        Objects.equals(vendor, advertiseInfo.vendor);
  }

  @Override
  public int hashCode() {
    return Objects.hash(advertised, apiExternalProductionEndpoint, apiExternalSandboxEndpoint, originalDevPortalUrl, apiOwner, vendor);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AdvertiseInfoDTO {\n");
    
    sb.append("    advertised: ").append(toIndentedString(advertised)).append("\n");
    sb.append("    apiExternalProductionEndpoint: ").append(toIndentedString(apiExternalProductionEndpoint)).append("\n");
    sb.append("    apiExternalSandboxEndpoint: ").append(toIndentedString(apiExternalSandboxEndpoint)).append("\n");
    sb.append("    originalDevPortalUrl: ").append(toIndentedString(originalDevPortalUrl)).append("\n");
    sb.append("    apiOwner: ").append(toIndentedString(apiOwner)).append("\n");
    sb.append("    vendor: ").append(toIndentedString(vendor)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}

