package org.wso2.apk.apimgt.rest.api.devportal.v1.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import javax.validation.constraints.*;


import io.swagger.annotations.*;
import java.util.Objects;



public class PaginationDTO   {
  
  private Integer offset;

  private Integer limit;

  private Integer total;

  private String next;

  private String previous;


  /**
   **/
  public PaginationDTO offset(Integer offset) {
    this.offset = offset;
    return this;
  }

  
  @ApiModelProperty(example = "0", value = "")
  @JsonProperty("offset")
  public Integer getOffset() {
    return offset;
  }
  public void setOffset(Integer offset) {
    this.offset = offset;
  }


  /**
   **/
  public PaginationDTO limit(Integer limit) {
    this.limit = limit;
    return this;
  }

  
  @ApiModelProperty(example = "10", value = "")
  @JsonProperty("limit")
  public Integer getLimit() {
    return limit;
  }
  public void setLimit(Integer limit) {
    this.limit = limit;
  }


  /**
   **/
  public PaginationDTO total(Integer total) {
    this.total = total;
    return this;
  }

  
  @ApiModelProperty(example = "1", value = "")
  @JsonProperty("total")
  public Integer getTotal() {
    return total;
  }
  public void setTotal(Integer total) {
    this.total = total;
  }


  /**
   * Link to the next subset of resources qualified. Empty if no more resources are to be returned. 
   **/
  public PaginationDTO next(String next) {
    this.next = next;
    return this;
  }

  
  @ApiModelProperty(example = "", value = "Link to the next subset of resources qualified. Empty if no more resources are to be returned. ")
  @JsonProperty("next")
  public String getNext() {
    return next;
  }
  public void setNext(String next) {
    this.next = next;
  }


  /**
   * Link to the previous subset of resources qualified. Empty if current subset is the first subset returned. 
   **/
  public PaginationDTO previous(String previous) {
    this.previous = previous;
    return this;
  }

  
  @ApiModelProperty(example = "", value = "Link to the previous subset of resources qualified. Empty if current subset is the first subset returned. ")
  @JsonProperty("previous")
  public String getPrevious() {
    return previous;
  }
  public void setPrevious(String previous) {
    this.previous = previous;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    PaginationDTO pagination = (PaginationDTO) o;
    return Objects.equals(offset, pagination.offset) &&
        Objects.equals(limit, pagination.limit) &&
        Objects.equals(total, pagination.total) &&
        Objects.equals(next, pagination.next) &&
        Objects.equals(previous, pagination.previous);
  }

  @Override
  public int hashCode() {
    return Objects.hash(offset, limit, total, next, previous);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PaginationDTO {\n");
    
    sb.append("    offset: ").append(toIndentedString(offset)).append("\n");
    sb.append("    limit: ").append(toIndentedString(limit)).append("\n");
    sb.append("    total: ").append(toIndentedString(total)).append("\n");
    sb.append("    next: ").append(toIndentedString(next)).append("\n");
    sb.append("    previous: ").append(toIndentedString(previous)).append("\n");
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

